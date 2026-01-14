package logic

import (
	"context"
	"fmt"

	"beaver-server/app/usercenter/rpc/internal/svc"
	"beaver-server/app/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 重置密码
func (l *ResetPasswordLogic) ResetPassword(in *pb.ResetPasswordReq) (*pb.ResetPasswordResp, error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, fmt.Errorf("用户 %s 不存在: %v", in.Email, err)
	}

	updatedUser := user
	updatedUser.Password = in.NewPassword

	if err = l.svcCtx.UserModel.Update(l.ctx, updatedUser); err != nil {
		return nil, fmt.Errorf("重置用户 %s 密码失败: %v", in.Email, err)
	}

	return &pb.ResetPasswordResp{
		Message: "重置密码成功",
	}, nil
}
