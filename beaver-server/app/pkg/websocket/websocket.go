package websocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForwardRequest struct {
	Addr   string
	Method string
	Token  string
	UserID string
	Body   *bytes.Buffer
}

type Response struct {
	Code   int             `json:"code"`
	Msg    string          `json:"msg"`
	Result json.RawMessage `json:"result"`
}

type WsProxyReq struct {
	UserID         string                 `header:"Beaver-User-Id"`
	Command        WebsocketCommand       `json:"command"`
	TargetID       string                 `json:"targetId"`
	Type           WebsocketType          `json:"type"`
	Body           map[string]interface{} `json:"body"`
	ConversationId string                 `json:"conversationId"`
}

func SendMessageToWs(addr string, command WebsocketCommand, types WebsocketType, senderID string, targetID string, requestBody map[string]interface{}, conversationId string) error {
	apiEndpoint := fmt.Sprintf("http://%s/api/chat/proxySendMsg", addr)

	wsProxyReq := WsProxyReq{
		UserID:         senderID,
		Command:        command,
		TargetID:       targetID,
		Type:           types,
		Body:           requestBody,
		ConversationId: conversationId,
	}
	body, _ := json.Marshal(wsProxyReq)

	_, err := ForwardMessage(ForwardRequest{
		Addr:   apiEndpoint,
		Method: "POST",
		Token:  "",
		UserID: senderID,
		Body:   bytes.NewBuffer(body),
	})
	return err
}

func ForwardMessage(forwardReq ForwardRequest) (json.RawMessage, error) {
	client := &http.Client{}

	var req *http.Request
	var err error

	// 根据请求方法生成对应的HTTP请求
	if forwardReq.Method == "GET" {
		req, err = http.NewRequest("GET", forwardReq.Addr, nil)
	} else if forwardReq.Method == "POST" {
		req, err = http.NewRequest("POST", forwardReq.Addr, forwardReq.Body)
	} else {
		return nil, fmt.Errorf("不支持的请求方法: %s", forwardReq.Method)
	}

	if err != nil {
		return nil, fmt.Errorf("API请求创建错误: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", forwardReq.Token)           // 使用Token进行鉴权
	req.Header.Set("Beaver-User-Id", forwardReq.UserID) // 使用Token进行鉴权

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API请求错误: %v", err)
	}
	defer resp.Body.Close()

	// 检查API响应
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("消息转发未成功: %v", resp.Status)
	}

	// 读取API响应
	byteData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("消息转发错误: %v", err)
	}

	// 解析API响应
	var authResponse Response
	authErr := json.Unmarshal(byteData, &authResponse)
	if authErr != nil {
		return nil, fmt.Errorf("消息转发错误: %v", authErr)
	}

	if authResponse.Code != 0 {
		return nil, fmt.Errorf("消息转发失败: %v", authResponse.Msg)
	}
	sendAjaxJSON, _ := json.Marshal(authResponse.Result)

	logx.Infof("消息转发成功: %s", string(sendAjaxJSON))
	return authResponse.Result, nil
}
