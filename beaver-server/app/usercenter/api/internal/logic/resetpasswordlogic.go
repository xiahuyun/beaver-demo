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

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 重置密码
func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req *types.ResetPasswordRequest) (resp *types.ResetPasswordResponse, err error) {

	cacheVerificationCode, err := l.svcCtx.Redis.Get(l.ctx, GetEmailVerificationCodeKey(req.Email, "reset")).Result()
	if err != nil {
		return nil, fmt.Errorf("验证码过期或不存在")
	}
	if cacheVerificationCode != req.VerificationCode {
		return nil, fmt.Errorf("验证码错误")
	}
	l.svcCtx.Redis.Del(l.ctx, GetEmailVerificationCodeKey(req.Email, "register"))

	if !isEmailValid(req.Email) {
		return nil, fmt.Errorf("无效的邮箱格式")
	}

	resp = &types.ResetPasswordResponse{}
	rpcResp, err := l.svcCtx.UsercenterRpc.ResetPassword(l.ctx, &usercenter.ResetPasswordReq{
		Email:       req.Email,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		l.Logger.Errorf("reset password user %s failed, message: %s", req.Email, err.Error())
		return nil, fmt.Errorf("重置密码失败")
	}

	l.Logger.Infof("reset password user %s success, message: %s", req.Email, rpcResp.Message)
	resp = &types.ResetPasswordResponse{
		Code:    http.StatusOK,
		Message: rpcResp.Message,
	}
	return
}
