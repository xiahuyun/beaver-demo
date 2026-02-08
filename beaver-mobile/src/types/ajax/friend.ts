// 搜索请求
export interface ISearchReq {
  email: string;
}

// 搜索响应
export interface ISearchRes {
  userId: string;
  nickname: string;
  fileName: string;
  abstract: string;
  notice: string;
  isFriend: boolean;
  conversationId: string;
  email: string;
}

// 好友信息请求
export interface IFriendInfoReq {
  friendId: string;
}

// 好友信息
export interface IFriendInfo {
  userId: string;
  nickname: string;
  fileName: string;
  abstract: string;
  notice: string;
  isFriend: boolean;
  conversationId: string;
  email: string;
  source?: string; // 好友关系来源：email/qrcode
}

// 添加好友请求
export interface IAddFriendReq {
  friendId: string;
  verify?: string;
  source: string; // 添加好友来源：email(邮箱搜索)/qrcode(扫码)
}

// 好友列表请求
export interface IFriendListReq {
  page?: number;
  limit?: number;
}

// 好友列表响应
export interface IFriendListRes {
  list: IFriendInfo[];
}

// 添加好友响应
export interface IAddFriendRes {}

