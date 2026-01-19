package logic

import (
	"context"
	"fmt"

	"beaver-server/app/usercenter/rpc/internal/svc"
	"beaver-server/app/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录
func (l *LoginUserLogic) LoginUser(in *pb.LoginUserReq) (*pb.LoginUserResp, error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, fmt.Errorf("user %s not found, err: %v", in.Email, err)
	}
	if user.Password != in.Password {
		return nil, fmt.Errorf("password error")
	}
	return &pb.LoginUserResp{
		Id:      user.Id,
		Name:    user.Name,
		Message: "login success",
	}, nil
}
