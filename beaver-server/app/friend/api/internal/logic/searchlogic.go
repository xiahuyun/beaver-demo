// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"beaver-server/app/friend/api/internal/svc"
	"beaver-server/app/friend/api/internal/types"
	"beaver-server/app/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过邮箱搜索用户
func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Search 通过邮箱搜索用户
func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchRes, err error) {
	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	// 调用用户中心服务搜索用户
	userInfo, err := l.svcCtx.UsercenterRpc.SearchUser(l.ctx, &pb.SearchUserReq{
		Email: req.Email,
	})
	if err != nil {
		return &types.SearchRes{
			Code: 1,
		}, err
	}
	return &types.SearchRes{
		Code: 0,
		Result: types.User{
			Name:     userInfo.UserInfo.Name,
			NickName: userInfo.UserInfo.Nickname,
			Gender:   userInfo.UserInfo.Gender,
			Avatar:   userInfo.UserInfo.Avatar,
			Email:    userInfo.UserInfo.Email,
		},
	}, nil
}
