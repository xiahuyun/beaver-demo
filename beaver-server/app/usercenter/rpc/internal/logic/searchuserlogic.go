package logic

import (
	"context"

	"beaver-server/app/usercenter/rpc/internal/svc"
	"beaver-server/app/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 搜索用户
func (l *SearchUserLogic) SearchUser(in *pb.SearchUserReq) (*pb.SearchUserRes, error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, err
	}
	return &pb.SearchUserRes{
		UserInfo: &pb.UserInfo{
			Name:     user.Name,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Email:    user.Email,
			Gender:   user.Gender,
		},
	}, nil
}
