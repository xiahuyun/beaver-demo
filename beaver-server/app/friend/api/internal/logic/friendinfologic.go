// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"beaver-server/app/friend/api/internal/svc"
	"beaver-server/app/friend/api/internal/types"
	"beaver-server/app/friend/rpc/friend"
	"beaver-server/app/usercenter/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取好友详细信息
func NewFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendInfoLogic {
	return &FriendInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendInfoLogic) FriendInfo(req *types.FriendInfoReq) (resp *types.FriendInfoRes, err error) {
	if req.FriendID == "" {
		return &types.FriendInfoRes{
			Code:    1,
			Message: "friendId is required",
		}, nil
	}
	if req.UserID == req.FriendID {
		return &types.FriendInfoRes{
			Code:    1,
			Message: "不能查询自己的信息",
		}, nil
	}

	userRes, err := l.svcCtx.UsercenterRpc.SearchUser(l.ctx, &usercenter.SearchUserReq{
		Email: req.FriendID,
	})
	if err != nil {
		l.Logger.Errorf("获取用户信息失败: friendID=%s, error=%v", req.FriendID, err)
		return &types.FriendInfoRes{
			Code:    1,
			Message: "用户不存在",
		}, nil
	}

	l.Logger.Infof("获取用户信息成功: friendID=%s", userRes.UserInfo.Email)

	// 生成会话Id
	conversationID, err := GenerateConversation([]string{req.UserID, req.FriendID})
	if err != nil {
		l.Logger.Errorf("生成会话Id失败: %v", err)
		return nil, fmt.Errorf("生成会话Id失败: %v", err)
	}

	isFriendReq := &friend.IsFriendReq{
		UserId1: req.UserID,
		UserId2: req.FriendID,
	}
	friendRes, err := l.svcCtx.FriendRpc.IsFriend(l.ctx, isFriendReq)
	if err != nil {
		l.Logger.Errorf("检查好友关系失败: %v", err)
		return nil, fmt.Errorf("检查好友关系失败: %v", err)
	}

	return &types.FriendInfoRes{
		Code:    0,
		Message: "success",
		Result: types.FriendInfo{
			UserID:         userRes.UserInfo.Email,
			NickName:       userRes.UserInfo.Name,
			Avatar:         userRes.UserInfo.Avatar,
			ConversationID: conversationID,
			IsFriend:       friendRes.IsFriend,
		},
	}, nil
}

/*
 * @description: 生成会话Id
 */
func GenerateConversation(userIds []string) (string, error) {
	if len(userIds) == 1 {
		return userIds[0], nil
	} else if len(userIds) == 2 {
		sort.Strings(userIds)
		return strings.Join(userIds, "_"), nil
	} else {
		return "", errors.New("userIds must have a length of 1 or 2")
	}
}
