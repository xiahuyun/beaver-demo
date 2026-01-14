// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"net/http"

	"beaver-server/app/usercenter/api/internal/svc"
	"beaver-server/app/usercenter/api/internal/types"
	"beaver-server/app/usercenter/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginUserLogic) LoginUser(req *types.LoginUserRequest) (resp *types.LoginUserResponse, err error) {

	rpcResp, err := l.svcCtx.UsercenterRpc.LoginUser(l.ctx, &usercenter.LoginUserReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("login user %s failed, err: %v", req.Email, err)
	}

	return &types.LoginUserResponse{
		Code:    http.StatusOK,
		Message: rpcResp.Message,
	}, nil
}
