import { IWsContent } from "./websocket";

/**
 * @description: WebSocket 命令类型枚举 - 与服务端保持一致
 */
export enum WsCommand {
  /**
   * @description: 聊天消息类
   */  
  CHAT_MESSAGE = 'CHAT_MESSAGE',
  /**
   * @description: 好友关系类
   */  
  FRIEND_OPERATION = 'FRIEND_OPERATION',
  /**
   * @description: 群组操作类
   */  
  GROUP_OPERATION = 'GROUP_OPERATION',
  /**
   * @description: 用户信息类
   */  
  USER_PROFILE = 'USER_PROFILE',
  /**
   * @description: 系统通知类
   */  
  SYSTEM_NOTIFICATION = 'SYSTEM_NOTIFICATION',
  /**
   * @description: 在线状态类
   */  
  PRESENCE = 'PRESENCE',
  /**
   * @description: 消息同步类
   */  
  MESSAGE_SYNC = 'MESSAGE_SYNC',
  /**
   * @description: 心跳消息类（应用级心跳）
   */  
  HEARTBEAT = 'HEARTBEAT',
}

/**
 * @description: WebSocket 消息接口 - 与服务端保持一致
 */
export interface IWsMessage {
  command: WsCommand;
  content: IWsContent;
}
