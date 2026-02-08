// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"net/http"
	"time"

	"beaver-server/app/pkg/utils/device"
	"beaver-server/app/websocket/api/internal/svc"
	"beaver-server/app/websocket/api/internal/types"

	wsutils "beaver-server/app/websocket/api/internal/logic/websocket/utils"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatWebsocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// WebSocket连接
func NewChatWebsocketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatWebsocketLogic {
	return &ChatWebsocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatWebsocketLogic) ChatWebsocket(req *types.ChatWebsocketRequest, w http.ResponseWriter, r *http.Request) (resp *types.ChatWebsocketResponse, err error) {

	conn, err := UpgradeToWebsocket(w, r)
	if err != nil {
		l.Logger.Errorf("升级为 WebSocket 失败: %v", err)
		return
	}

	// 处理 WebSocket 连接
	configureWebsocket(conn, l.svcCtx)

	defer func() {
		conn.Close()
	}()

	l.Logger.Info("建立 WebSocket 连接")

	userAgent := r.Header.Get("User-Agent")
	deviceType := device.GetDeviceType(userAgent)
	userKey := wsutils.GetUserKey(req.UserID, deviceType)
	managerUserConnection(userKey, conn, req.UserID, deviceType)

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			l.Logger.Errorf("读取 WebSocket 消息失败: %v", err)
			break
		}

		l.Logger.Infof("收到 WebSocket 消息，类型: %d, 用户: %s, 内容: %s", messageType, req.UserID, p)
	}

	return &types.ChatWebsocketResponse{}, nil
}

func UpgradeToWebsocket(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logx.Errorf("升级为 WebSocket 失败: %v", err)
		return nil, err
	}

	// 关闭连接时，需要在逻辑层处理，而不是在 handler 层
	// defer conn.Close()

	return conn, nil
}

func configureWebsocket(conn *websocket.Conn, svcCtx *svc.ServiceContext) {
	// 配置 WebSocket 连接
	conn.SetReadLimit(svcCtx.Config.WebSocket.DefaultReadLimit)
	conn.SetReadDeadline(time.Now().Add(time.Duration(svcCtx.Config.WebSocket.DefaultReadTimeout) * time.Second))
	conn.SetWriteDeadline(time.Now().Add(time.Duration(svcCtx.Config.WebSocket.DefaultWriteTimeout) * time.Second))

	conn.SetPongHandler(func(appData string) error {
		conn.SetReadDeadline(time.Now().Add(time.Duration(svcCtx.Config.WebSocket.DefaultReadTimeout) * time.Second))
		return nil
	})
}

func managerUserConnection(userKey string, conn *websocket.Conn, userID string, deviceType string) {
	wsutils.WsMutex.Lock()
	defer wsutils.WsMutex.Unlock()

	addr := conn.RemoteAddr().String()

	// 初始化用户 WebSocket 信息
	wsUserInfo, ok := wsutils.WsOnlineUserMap[userKey]
	if !ok {
		wsUserInfo = &wsutils.WsUserInfo{
			WsClientMap: map[string]*websocket.Conn{
				addr: conn,
			},
		}
		wsutils.WsOnlineUserMap[userKey] = wsUserInfo
		logx.Infof("创建用户连接映射, 用户: %s, 设备: %s, userKey: %s", userID, deviceType, userKey)
	} else {
		if deviceType == device.Desktop || deviceType == device.Mobile {
			if len(wsUserInfo.WsClientMap) >= 1 {
				logx.Errorf("用户 %s 已在其他设备上登录, 设备: %s, userKey: %s, addr: %s", userID, deviceType, userKey, addr)
				logx.Errorf("请先在设备上登出, 设备: %s, userKey: %s, addr: %s", deviceType, userKey, addr)
				return
			}
		}
		wsUserInfo.WsClientMap[addr] = conn
		logx.Infof("添加用户连接, 用户: %s, 设备: %s, userKey: %s, addr: %s", userID, deviceType, userKey, addr)
	}

	// 打印当前用户的所有连接状态
	deviceTypes := []string{device.Desktop, device.Mobile, device.Web}
	for _, dt := range deviceTypes {
		uk := wsutils.GetUserKey(userID, dt)
		if info, exists := wsutils.WsOnlineUserMap[uk]; exists {
			logx.Infof("用户: %s, 设备类型: %s, userKey: %s, 连接数: %d", userID, dt, uk, len(info.WsClientMap))
		}
	}
}
