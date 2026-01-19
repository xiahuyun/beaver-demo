// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"beaver-server/app/usercenter/api/internal/svc"
	"beaver-server/app/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户认证
func NewAuthenticationUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationUserLogic {
	return &AuthenticationUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationUserLogic) AuthenticationUser(req *types.AuthenticationUserRequest) (resp *types.AuthenticationUserResponse, err error) {
	payload, err := l.svcCtx.TokenClint.VerifyToken(req.Token, l.svcCtx.Config.JwtAuth.AccessSecret)
	if err != nil {
		return &types.AuthenticationUserResponse{
			Code:    1,
			Message: "token 无效",
		}, nil
	}

	l.Logger.Infof("用户 %s token 有效", payload.Name)

	return &types.AuthenticationUserResponse{
		Code:    0,
		Message: "token 有效",
	}, nil
}
