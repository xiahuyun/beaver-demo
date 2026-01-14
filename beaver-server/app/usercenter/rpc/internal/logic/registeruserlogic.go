package logic

import (
	"context"
	"fmt"

	"beaver-server/app/usercenter/model"
	"beaver-server/app/usercenter/rpc/internal/svc"
	"beaver-server/app/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册用户
func (l *RegisterUserLogic) RegisterUser(in *pb.RegisterUserReq) (*pb.RegisterUserResp, error) {

	switch in.RegisterType {
	case "email":
		return l.RegisterUserByEmail(in)
	case "phone":
		return l.RegisterUserByPhone(in)
	}

	return nil, fmt.Errorf("register type %s not support", in.RegisterType)
}

func (l *RegisterUserLogic) RegisterUserByEmail(in *pb.RegisterUserReq) (*pb.RegisterUserResp, error) {

	l.Logger.Infof("注册邮箱用户: %v", *in.Email)

	l.Logger.Infof("检查邮箱 %s 是否存在", *in.Email)
	exist, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, *in.Email)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Errorf("检查邮箱 %s 是否存在时出错: %v", *in.Email, err)
		return nil, err
	}
	if exist != nil {
		return nil, fmt.Errorf("邮箱 %s 已存在", *in.Email)
	}

	l.Logger.Infof("邮箱 %s 不存在，注册用户", *in.Email)
	result, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Email:        *in.Email,
		Password:     in.Password,
		RegisterType: in.RegisterType,
	})
	if err != nil {
		return nil, err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	l.Logger.Infof("邮箱 %s 注册成功，用户ID: %d", *in.Email, userID)
	return &pb.RegisterUserResp{
		Message: fmt.Sprintf("用户 %s 注册成功", *in.Email),
	}, nil
}

func (l *RegisterUserLogic) RegisterUserByPhone(in *pb.RegisterUserReq) (*pb.RegisterUserResp, error) {

	l.Logger.Infof("register phone user: %v", in.Phone)

	l.Logger.Infof("check whether the user phone %s exists", *in.Phone)
	exist, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, *in.Phone)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Errorf("check user phone %s error: %v", *in.Phone, err)
		return nil, err
	}
	if exist != nil {
		return nil, fmt.Errorf("phone %s already exists", *in.Phone)
	}

	l.Logger.Infof("user phone %s not exists, register the user", *in.Phone)
	result, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Phone:        *in.Phone,
		Password:     in.Password,
		RegisterType: in.RegisterType,
	})
	if err != nil {
		return nil, err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	l.Logger.Infof("register phone user %s success, userID: %d", *in.Phone, userID)
	return &pb.RegisterUserResp{
		Message: fmt.Sprintf("register phone user %s success, userID: %d", *in.Phone, userID),
	}, nil
}
