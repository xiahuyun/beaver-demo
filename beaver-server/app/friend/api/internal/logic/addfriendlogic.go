// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"beaver-server/app/friend/api/internal/svc"
	"beaver-server/app/friend/api/internal/types"
	"beaver-server/app/friend/model"
	"beaver-server/app/friend/rpc/friend"
	"beaver-server/app/notification/rpc/notification"
	"beaver-server/app/usercenter/rpc/usercenter"

	pkgtypes "beaver-server/app/pkg/types"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送好友验证请求
func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFriendLogic) AddFriend(req *types.AddFriendReq) (resp *types.AddFriendRes, err error) {

	isFriendReq := &friend.IsFriendReq{
		UserId1: req.UserID,
		UserId2: req.FriendID,
	}
	friendRes, err := l.svcCtx.FriendRpc.IsFriend(l.ctx, isFriendReq)
	if err != nil {
		l.Logger.Errorf("检查好友关系失败: %v", err)
		return nil, fmt.Errorf("检查好友关系失败: %v", err)
	}
	// 检查好友关系是否已存在
	if friendRes.IsFriend {
		l.Logger.Errorf("用户 %s 已添加好友 %s", req.UserID, req.FriendID)
		return nil, fmt.Errorf("用户 %s 已添加好友 %s", req.UserID, req.FriendID)
	}

	if _, err = l.svcCtx.UsercenterRpc.SearchUser(l.ctx, &usercenter.SearchUserReq{
		Email: req.FriendID,
	}); err != nil {
		l.Logger.Errorf("搜索用户 %s 失败: %v", req.FriendID, err)
		return nil, fmt.Errorf("搜索用户 %s 失败: %v", req.FriendID, err)
	}

	// 检查是否有待处理的好友验证请求
	verifyReq := &friend.IsFriendVerifyReq{
		UserId1: req.UserID,
		UserId2: req.FriendID,
	}

	verifyRes, err := l.svcCtx.FriendRpc.IsFriendVerify(l.ctx, verifyReq)
	if err != nil {
		l.Logger.Errorf("检查好友验证请求失败: %v", err)
		return nil, fmt.Errorf("检查好友验证请求失败: %v", err)
	}
	if verifyRes.IsFriendVerify {
		l.Logger.Errorf("用户 %s 已发送好友验证请求给 %s", req.UserID, req.FriendID)
		return nil, fmt.Errorf("用户 %s 已发送好友验证请求给 %s", req.UserID, req.FriendID)
	}

	nextVersion, err := l.svcCtx.Version.GetNextVersion("friendverify", "", "")
	if err != nil {
		l.Logger.Errorf("获取版本号失败: %v", err)
		return nil, err
	}

	verifyModel := model.Friendverify{
		VerifyId:   uuid.New().String(),
		SendUserId: req.UserID,
		RevUserId:  req.FriendID,
		Message:    req.Verify,
		Source:     req.Source,
		Version:    nextVersion,
	}
	err = l.svcCtx.DB.Create(&verifyModel).Error
	if err != nil {
		l.Logger.Errorf("创建好友验证请求失败: %v", err)
		return nil, fmt.Errorf("创建好友验证请求失败: %v", err)
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				l.Logger.Errorf("发送好友验证通知时 panic: %v", r)
			}
		}()

		payload, _ := json.Marshal(map[string]interface{}{
			"verifyId": verifyModel.VerifyId,
			"message":  verifyModel.Message,
			"source":   verifyModel.Source,
		})
		_, err := l.svcCtx.NotifyRpc.PushEvent(context.Background(), &notification.PushEventReq{
			EventType:   pkgtypes.EventTypeFriendRequest,
			Category:    pkgtypes.CategorySocial,
			FromUserId:  req.UserID,
			TargetId:    verifyModel.VerifyId,
			TargetType:  pkgtypes.TargetTypeUser,
			PayloadJson: string(payload),
			ToUserIds:   []string{req.FriendID},
			DedupHash:   verifyModel.VerifyId,
		})
		if err != nil {
			l.Logger.Errorf("投递好友申请通知失败: %v", err)
		}

	}()

	return &types.AddFriendRes{
		Version: nextVersion,
	}, nil
}
