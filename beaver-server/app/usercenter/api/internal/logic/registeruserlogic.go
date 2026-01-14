// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"beaver-server/app/usercenter/api/internal/svc"
	"beaver-server/app/usercenter/api/internal/types"

	"beaver-server/app/usercenter/rpc/usercenter"

	ptypes "beaver-server/app/pkg/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 注册用户
func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterUserLogic) RegisterUser(req *types.RegisterUserRequest) (resp *types.RegisterUserResponse, err error) {

	switch req.Type {
	case ptypes.RegisterTypeEmail:
		return l.RegisterUserByEmail(req)
	case ptypes.RegisterTypePhone:
		return l.RegisterUserByPhone(req)
	default:
		return nil, fmt.Errorf("register type %s not support", req.Type)
	}

}

func (l *RegisterUserLogic) RegisterUserByEmail(req *types.RegisterUserRequest) (resp *types.RegisterUserResponse, err error) {

	cacheVerificationCode, err := l.svcCtx.Redis.Get(l.ctx, GetEmailVerificationCodeKey(req.Email, "register")).Result()
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

	resp = &types.RegisterUserResponse{}
	rpcResp, err := l.svcCtx.UsercenterRpc.RegisterUser(l.ctx, &usercenter.RegisterUserReq{
		Email:        &req.Email,
		Password:     req.Password,
		RegisterType: req.Type,
	})
	if err != nil {
		l.Logger.Errorf("注册邮箱用户 %s 失败: %v", req.Email, err)
		return nil, err
	}

	logx.Infof("register user %s type %s success, message: %s", req.Email, req.Type, rpcResp.Message)
	resp.Message = rpcResp.Message
	resp.Code = http.StatusOK
	return resp, nil
}

func (l *RegisterUserLogic) RegisterUserByPhone(req *types.RegisterUserRequest) (resp *types.RegisterUserResponse, err error) {

	cacheVerificationCode, err := l.svcCtx.Redis.Get(l.ctx, GetPhoneVerificationCodeKey(req.Phone, "register")).Result()
	if err != nil {
		return nil, fmt.Errorf("验证码过期或不存在")
	}
	if cacheVerificationCode != req.VerificationCode {
		return nil, fmt.Errorf("验证码错误")
	}
	l.svcCtx.Redis.Del(l.ctx, GetPhoneVerificationCodeKey(req.Phone, "register"))

	if !isPhoneValid(req.Phone) {
		return nil, fmt.Errorf("phone %s is invalid, must be in format of 13800000000", req.Phone)
	}

	resp = &types.RegisterUserResponse{}
	rpcResp, err := l.svcCtx.UsercenterRpc.RegisterUser(l.ctx, &usercenter.RegisterUserReq{
		Phone:        &req.Phone,
		Password:     req.Password,
		RegisterType: req.Type,
	})
	if err != nil {
		return nil, fmt.Errorf("register user %s type %s failed, err: %v", req.Phone, req.Type, err)
	}

	logx.Infof("register user %s type %s success, message: %s", req.Phone, req.Type, rpcResp.Message)
	resp.Message = rpcResp.Message
	resp.Code = http.StatusOK
	return resp, nil
}

func isEmailValid(email string) bool {
	var emailRegex = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return emailRegex.MatchString(email)
}

func isPhoneValid(phone string) bool {
	// TODO: add phone regex
	var phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)
	return phoneRegex.MatchString(phone)
}

func isPasswordValid(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)

	typesCount := 0
	if hasLower {
		typesCount++
	}
	if hasUpper {
		typesCount++
	}
	if hasDigit {
		typesCount++
	}

	return typesCount == 3
}
