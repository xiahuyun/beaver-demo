package logic

import (
	"context"

	"beaver-server/app/friend/model"
	"beaver-server/app/friend/rpc/internal/svc"
	"beaver-server/app/friend/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFriendVerifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFriendVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFriendVerifyLogic {
	return &IsFriendVerifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 是否验证中
func (l *IsFriendVerifyLogic) IsFriendVerify(in *pb.IsFriendVerifyReq) (*pb.IsFriendVerifyRes, error) {

	friendship, err := l.svcCtx.FriendverifyModel.IsFriendVerify(l.ctx, in.UserId1, in.UserId2)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	l.Logger.Infof("IsFriendVerify: userId1=%s, userId2=%s, isFriendVerify=%v", in.UserId1, in.UserId2, friendship != nil)
	return &pb.IsFriendVerifyRes{
		IsFriendVerify: friendship != nil,
	}, nil
}
