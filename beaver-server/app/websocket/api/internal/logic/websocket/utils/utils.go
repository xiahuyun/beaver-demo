package utils

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"beaver-server/app/pkg/rand"
	wstypes "beaver-server/app/pkg/websocket"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

var WsOnlineUserMap = make(map[string]*WsUserInfo)
var WsMutex sync.RWMutex

type WsUserInfo struct {
	WsClientMap map[string]*websocket.Conn // 用户管理的所有 WebSocket 通道
}

// GetUserKey 生成用户连接的唯一key
func GetUserKey(userID string, deviceType string) string {
	return userID + "_" + deviceType
}

// SendMsgToUser 只发送消息给指定用户的所有设备
func SendMsgToUser(targetUserID string, command wstypes.WebsocketCommand, content wstypes.WebsocketContent) {
	WsMutex.RLock()
	defer WsMutex.RUnlock()

	fmt.Println("WsOnlineUserMap:", WsOnlineUserMap)

	// 遍历用户的所有连接
	for userKey, userInfo := range WsOnlineUserMap {
		/*
			if strings.HasPrefix(userKey, targetUserID+"_") {
				deviceType := strings.TrimPrefix(userKey, targetUserID+"_")
				jsonContent, _ := json.Marshal(content)
				logx.Infof("发送消息给用户：%s, 设备类型：%s, 消息内容：%s", targetUserID, deviceType, string(jsonContent))
				sendWsMapMsg(userInfo.WsClientMap, command, content)
			}
		*/
		//deviceType := strings.TrimPrefix(userKey, targetUserID+"_")
		jsonContent, _ := json.Marshal(content)
		logx.Infof("发送消息给用户：%s, 设备类型：%s, 消息内容：%s", targetUserID, userKey, string(jsonContent))
		sendWsMapMsg(userInfo.WsClientMap, command, content)
	}
}

// sendWsMapMsg 给一组的 WebSocket 通道发送消息
func sendWsMapMsg(wsMap map[string]*websocket.Conn, command wstypes.WebsocketCommand, content wstypes.WebsocketContent) {
	for addr, conn := range wsMap {
		if err := WsResponse(conn, command, content); err != nil {
			logx.Errorf("发送WebSocket消息失败, 地址: %s, 错误: %v", addr, err)
			// 如果发送失败，关闭连接并从map中移除
			conn.Close()
			delete(wsMap, addr)
		}
	}
}

type Response struct {
	Code       int                      `json:"code"`
	Command    wstypes.WebsocketCommand `json:"command"`
	Content    wstypes.WebsocketContent `json:"content"`
	MessageID  string                   `json:"messageId"`
	ServerTime int64                    `json:"serverTime"`
}

func WsResponse(conn *websocket.Conn, command wstypes.WebsocketCommand, content wstypes.WebsocketContent) error {
	code := 0

	response := Response{
		Command:    command,
		Code:       code,
		Content:    content,
		MessageID:  rand.GenerateRandomString(8),
		ServerTime: time.Now().Unix(),
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		logx.Errorf("序列化WebSocket响应失败: %v", err)
		return err
	}

	// 设置写入超时
	if err := conn.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
		logx.Errorf("设置WebSocket写入超时失败: %v", err)
		return err
	}

	if err := conn.WriteMessage(websocket.TextMessage, responseJSON); err != nil {
		logx.Errorf("发送WebSocket消息失败: %v", err)
		return err
	}

	return nil
}
