package logic

import (
	"context"
	"fmt"

	"beaver-server/app/friend/model"
	"beaver-server/app/friend/rpc/internal/svc"
	"beaver-server/app/friend/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFriendLogic {
	return &IsFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 是否好友
func (l *IsFriendLogic) IsFriend(in *pb.IsFriendReq) (*pb.IsFriendRes, error) {
	generatedFriendshipId, err := GenerateFriendshipId(in.UserId1, in.UserId2)
	if err != nil {
		return nil, err
	}
	friendship, err := l.svcCtx.FriendModel.FindOneByFriendshipId(l.ctx, generatedFriendshipId)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	l.Logger.Infof("IsFriend: userId1=%s, userId2=%s, isFriend=%v", in.UserId1, in.UserId2, friendship != nil)
	return &pb.IsFriendRes{
		IsFriend: friendship != nil,
	}, nil
}

func GenerateFriendshipId(userId1, userId2 string) (string, error) {
	if userId1 == "" || userId2 == "" {
		return "", fmt.Errorf("userId1 or userId2 cannot be empty")
	}
	if userId1 == userId2 {
		return "", fmt.Errorf("userId1 and userId2 cannot be the same")
	}

	// 确保 userId1 小于 userId2
	if userId1 > userId2 {
		return fmt.Sprintf("%s_%s", userId2, userId1), nil
	}
	return fmt.Sprintf("%s_%s", userId1, userId2), nil
}
