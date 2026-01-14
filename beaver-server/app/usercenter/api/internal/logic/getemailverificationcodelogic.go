// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"beaver-server/app/usercenter/api/internal/svc"
	"beaver-server/app/usercenter/api/internal/types"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/gomail.v2"
)

const (
	VerificationCodeLength         = 6
	VerificationCodeCharSet        = "123456789"
	VerificationCodeExpireDuration = time.Minute * 5
)

type GetEmailVerificationCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取邮箱验证码
func NewGetEmailVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailVerificationCodeLogic {
	return &GetEmailVerificationCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmailVerificationCodeLogic) GetEmailVerificationCode(req *types.GetEmailVerificationCodeRequest) (resp *types.GetEmailVerificationCodeResponse, err error) {
	verificationCode := l.GenerateEmailVerificationCode()
	l.Logger.Infof("generate request id %s verification code %s", l.ctx.Value("requestId"), verificationCode)

	if ok := l.SendEmailVerificationCode(req.Email, verificationCode); !ok {
		l.Logger.Errorf("send request id %s verification code %s to email %s failed", l.ctx.Value("requestId"), verificationCode, req.Email)
		return nil, fmt.Errorf("send request id %s verification code %s to email %s failed", l.ctx.Value("requestId"), verificationCode, req.Email)
	}
	l.Logger.Infof("send request id %s verification code %s to email %s success", l.ctx.Value("requestId"), verificationCode, req.Email)

	key := GetEmailVerificationCodeKey(req.Email, req.Type)
	if err := l.svcCtx.Redis.Set(l.ctx, key, verificationCode, VerificationCodeExpireDuration).Err(); err != nil {
		l.Logger.Errorf("set request id %s verification code %s to email %s failed, err: %v", l.ctx.Value("requestId"), verificationCode, req.Email, err)
		return nil, fmt.Errorf("set request id %s verification code %s to email %s failed, err: %v", l.ctx.Value("requestId"), verificationCode, req.Email, err)
	}
	l.Logger.Infof("set email %s type %s verification code %s to cache success, expire duration %s", req.Email, req.Type, verificationCode, VerificationCodeExpireDuration.String())

	resp = &types.GetEmailVerificationCodeResponse{
		Message: "send verification code success",
		Code:    http.StatusOK,
	}
	return resp, nil
}

func (l *GetEmailVerificationCodeLogic) GenerateEmailVerificationCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	var builder strings.Builder
	builder.Grow(VerificationCodeLength)

	for range VerificationCodeLength {
		builder.WriteByte(VerificationCodeCharSet[rnd.Intn(len(VerificationCodeCharSet))])
	}

	return builder.String()
}

func (l *GetEmailVerificationCodeLogic) SendEmailVerificationCode(email, verificationCode string) bool {
	m := gomail.NewMessage()
	m.SetHeader("From", "xiahuyun043@126.com")
	m.SetHeader("To", "xiahuyun043@126.com")
	m.SetHeader("Subject", "友喵验证码")
	m.SetBody("text/plain", fmt.Sprintf("您的验证码是: %s", verificationCode))

	dialer := gomail.NewDialer("smtp.126.com", 465, "xiahuyun043@126.com", "XQurrpewgBJPcXGj")

	if err := dialer.DialAndSend(m); err != nil {
		l.Logger.Errorf("send request id %s verification code %s to email %s failed, err: %v", l.ctx.Value("requestId"), verificationCode, email, err)
		return false
	}

	return true
}

func GetEmailVerificationCodeKey(email string, codeType string) string {
	return fmt.Sprintf("email_verification_code_%s_%s", email, codeType)
}

func GetPhoneVerificationCodeKey(phone string, codeType string) string {
	return fmt.Sprintf("phone_verification_code_%s_%s", phone, codeType)
}
