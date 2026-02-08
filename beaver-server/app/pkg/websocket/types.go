package websocket

import "encoding/json"

type WebsocketType string

// send（发起消息给服务端）
// receive（发给对方设备）
// sync（发给自己其他设备进行记录同步）

const (
	PrivateMessageSend WebsocketType = "private_message_send" // 客户端->服务端 私聊消息发送
	GroupMessageSend   WebsocketType = "group_message_send"   // 客户端->服务端 群聊消息发送
	// --------------------------------------------------------
	// --------------------------------------------------------

	// 会话信息同步
	ChatConversationMetaReceive    WebsocketType = "chat_conversation_meta_receive"    //  服务端->客户端 会话信息同步
	ChatUserConversationReceive    WebsocketType = "chat_user_conversation_receive"    //  服务端->客户端 用户会话信息同步
	ChatConversationMessageReceive WebsocketType = "chat_conversation_message_receive" //  服务端->客户端 会话消息同步
)
const (
	// -------------------------------------------------------------------------------------
	FriendReceive       WebsocketType = "friend_receive"        // 服务端->客户端 好友信息同步
	FriendVerifyReceive WebsocketType = "friend_verify_receive" // 服务端->客户端 好友验证信息同步
)

// -------------------------------------------------------------------------------------

const (
	GroupReceive            WebsocketType = "group_receive"              // 服务端->客户端 群组信息同步
	GroupJoinRequestReceive WebsocketType = "group_join_request_receive" // 服务端->客户端 群成员添加请求
	GroupMemberReceive      WebsocketType = "group_member_receive"       // 服务端->客户端 群成员变动（加入，离开、被踢出等）

)

const (
	// --------------------------------------------------------
	UserReceive WebsocketType = "user_receive" // 服务端->客户端 用户信息同步
)

const (
	// 通知中心
	NotificationReceive         WebsocketType = "notification_receive"           // 服务端->客户端 通知推送
	NotificationMarkReadReceive WebsocketType = "notification_mark_read_receive" // 服务端->客户端 标记已读同步
)

const (
	// 表情中心
	EmojiReceive WebsocketType = "emoji_receive" // 服务端->客户端 表情数据同步
)

type WebsocketContent struct {
	Timestamp int64         `json:"timestamp"` //消息发送时间
	MessageID string        `json:"messageId"` //客户端消息ID
	Data      WebsocketData `json:"data"`      //消息内容
}

type WebsocketData struct {
	Type           WebsocketType   `json:"type"`           // 消息类型
	ConversationID string          `json:"conversationId"` // 会话ID
	Body           json.RawMessage `json:"body"`           // 消息内容
}
