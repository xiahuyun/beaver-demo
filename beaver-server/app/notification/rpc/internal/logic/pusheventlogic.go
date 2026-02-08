package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"beaver-server/app/notification/rpc/internal/svc"
	"beaver-server/app/notification/rpc/pb"

	"beaver-server/app/notification/model"
	pkgtypes "beaver-server/app/pkg/types"
	"beaver-server/app/pkg/websocket"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type PushEventLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushEventLogic {
	return &PushEventLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushEventLogic) PushEvent(in *pb.PushEventReq) (*pb.PushEventRes, error) {
	if len(in.ToUserIds) == 0 {
		return nil, errors.New("to_user_ids 不能为空")
	}
	if in.EventType == "" || in.Category == "" {
		return nil, errors.New("event_type 和 category 不能为空")
	}

	eventID := uuid.NewString()

	// 生成事件版本（全局）
	eventVersion, err := l.svcCtx.Version.GetNextVersion(pkgtypes.VersionScopeEventGlobal, "", "")
	if err != nil {
		return nil, err
	}

	fmt.Println("PushEvent eventVersion:", eventVersion)

	err = l.svcCtx.NotificationEventModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		event := &model.NotificationEvent{
			EventId:    eventID,
			EventType:  in.EventType,
			Category:   in.Category,
			Version:    eventVersion,
			TargetType: in.TargetType,
			Payload:    sql.NullString{String: in.PayloadJson, Valid: true},
			Priority:   5,
			Status:     1,
			DedupHash:  sql.NullString{String: in.DedupHash, Valid: true},
		}
		if event.FromUserId != "" {
			event.FromUserId = in.FromUserId
		}
		if event.TargetId != "" {
			event.TargetId = in.TargetId
		}

		if _, err := l.svcCtx.NotificationEventModel.Insert(ctx, session, event); err != nil {
			return err
		}

		inboxRows := make([]model.NotificationInbox, 0, len(in.ToUserIds))
		for _, uid := range in.ToUserIds {
			inboxVersion, err := l.svcCtx.Version.GetNextVersion(pkgtypes.VersionScopeInboxPerUser, "user_id", uid)
			if err != nil {
				return err
			}

			inboxRows = append(inboxRows, model.NotificationInbox{
				UserId:    uid,
				EventId:   eventID,
				EventType: in.EventType,
				Category:  in.Category,
				Version:   inboxVersion,
				IsRead:    0,
				ReadAt:    sql.NullTime{},
				Status:    1,
				Silent:    0,
			})
		}
		for _, inboxRow := range inboxRows {
			if _, err := l.svcCtx.NotificationInboxModel.Insert(l.ctx, &inboxRow); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	go func(addr string, toUsers []string, eventID string, eventVersion int64) {
		var tableUpdates []map[string]interface{}
		eventUpdates := map[string]interface{}{
			"table": "notification_event",
			"data": []map[string]interface{}{
				{
					"version": eventVersion,
					"eventId": eventID,
				},
			},
		}
		tableUpdates = append(tableUpdates, eventUpdates)

		for _, uid := range toUsers {
			inboxVersion, err := l.svcCtx.Version.GetNextVersion(pkgtypes.VersionScopeInboxPerUser, "user_id", uid)
			if err != nil {
				l.Logger.Error("获取inbox版本失败", err)
				continue
			}
			inboxUpdates := map[string]interface{}{
				"table":  "notification_inbox",
				"userId": uid,
				"data": []map[string]interface{}{
					{
						"version": inboxVersion,
						"eventId": eventID,
					},
				},
			}
			userTableUpdates := append([]map[string]interface{}{}, tableUpdates...)
			userTableUpdates = append(userTableUpdates, inboxUpdates)
			websocket.SendMessageToWs(addr, websocket.NOTIFICATION, websocket.NotificationReceive, in.FromUserId, uid, map[string]interface{}{
				"tableUpdates": userTableUpdates,
			}, "")
		}
	}("127.0.0.1:20802", in.ToUserIds, eventID, eventVersion)

	return &pb.PushEventRes{
		EventId: eventID,
		Version: eventVersion,
	}, nil
}
