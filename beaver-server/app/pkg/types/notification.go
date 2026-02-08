package types

const (
	// 好友相关
	EventTypeFriendRequest       = "friend_request"
	EventTypeFriendRequestAccept = "friend_request_accept"
	EventTypeFriendRequestReject = "friend_request_reject"
)

// 分类常量
const (
	CategoryMoment = "moment" // 动态
	CategorySocial = "social" // 社交
	CategoryGroup  = "group"  // 群组
)

// 目标类型常量（事件指向的实体类型）
const (
	TargetTypeMoment        = "moment"
	TargetTypeMomentComment = "moment_comment"
	TargetTypeGroup         = "group"
	TargetTypeUser          = "user"
)

// 版本号作用域（用于 Version 约定）
const (
	// NotificationEvent：全局递增版本
	VersionScopeEventGlobal = "notificationEvent"
	// NotificationInbox：按用户递增版本（per user）
	VersionScopeInboxPerUser = "notificationInbox"
	// NotificationRead：按用户+分类递增版本（per user per category）
	VersionScopeCursorPerUser = "notification_reads"
)
