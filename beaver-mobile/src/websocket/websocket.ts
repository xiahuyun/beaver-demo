import { IWsMessage } from '@/src/types/websocket/command'

/**
 * @description: WebSocket配置选项
 */
interface WsConfig {
  reconnectInterval?: number; // 重连间隔时间
  maxReconnectAttempts?: number; // 最大重连次数
  heartbeatInterval?: number; // 心跳间隔
}

/**
 * @description: WebSocket事件回调
 */
interface WsEventCallbacks {
  onMessage?: (message: IWsMessage) => void;
  onConnect?: () => void;
  onDisconnect?: () => void;
  onError?: (error: any) => void;
}

/**
 * @description: WebSocket状态
 */
type WsStatus = 'connecting' | 'connected' | 'closed' | 'error';

/**
 * @description: WebSocket连接管理器 - 纯连接层，不处理业务逻辑
 */
class WsManager {
	private reconnectInterval: number;
	private maxReconnectAttempts: number;
	private heartbeatInterval: number;
	public isConnected = false;
	public status: WsStatus = 'closed';
	private reconnectAttempts = 0;
	private eventCallbacks: WsEventCallbacks = {};
	
	constructor(config: WsConfig = {}) {
		this.reconnectInterval = config.reconnectInterval || 5000;
		this.maxReconnectAttempts = config.maxReconnectAttempts || 5;
		this.heartbeatInterval = config.heartbeatInterval || 30000;

		console.log("running 1111");
		this.initEventListeners();
		//this.initAppStateListeners();
		console.log("running 2222");
	}

	/**
	* @description: 设置事件回调
	*/
	public setEventCallbacks(callbacks: WsEventCallbacks) {
		this.eventCallbacks = { ...this.eventCallbacks, ...callbacks };
	}

	public async initSocket(): Promise<void> {
		return new Promise((resolve, reject) => {
			try {
				// mock request connection
				uni.connectSocket({
					url: `ws://127.0.0.1:20802/api/chat/websocket?token=test&userId=xiahuyun043@126.com`,
					method: 'GET',
					success(res) {
						resolve();
					},
					fail(error) {
						reject(error);
					}
				});
				
				uni.connectSocket({
					url: `ws://127.0.0.1:20802/api/chat/websocket?token=test&userId=15968152314@163.com`,
					method: 'GET',
					success(res) {
						resolve();
					},
					fail(error) {
						reject(error);
					}
				});
			} catch (error) {
				console.error('WebSocket 连接异常:', error);
				reject(error);
			}
		})
	}

	/**
	* @description: 初始化WebSocket事件监听
	*/
	private initEventListeners(): void {
		uni.onSocketOpen(this.onOpen.bind(this));
		uni.onSocketClose(this.onClose.bind(this));
		uni.onSocketError(this.onError.bind(this));
		uni.onSocketMessage(this.onMessage.bind(this));
	}
	
	/**
	* @description: WebSocket连接成功回调
	*/
	private onOpen(): void {
		console.log('WebSocket 连接成功');
		this.isConnected = true;
		this.status = 'connected';
		this.reconnectAttempts = 0;
		//this.clearTimers();
		//this.startHeartbeat();
	
		// 触发连接成功回调
		if (this.eventCallbacks.onConnect) {
			this.eventCallbacks.onConnect();
		}
	}
	
	/**
	* @description: WebSocket连接关闭回调
	*/
	private onClose(data: any): void {
		console.log('WebSocket 连接关闭', data);
		this.isConnected = false;
		this.status = 'closed';
		//this.clearTimers();
	
	    // 分析关闭原因
	    if (data.code === 1006) {
	      console.warn('WebSocket 异常关闭 (1006)，可能是网络问题或服务器问题');
	    } else if (data.code === 1000) {
	      console.log('WebSocket 正常关闭');
	    } else if (data.code === 1001) {
	      console.log('WebSocket 端点离线或服务器重启');
	    }
	
	    // 触发断开连接回调
	    if (this.eventCallbacks.onDisconnect) {
	      this.eventCallbacks.onDisconnect();
	    }
	
		/*
			if (!this.isManualClose) {
			  // 如果是 1006 错误，稍微延长重连间隔
			  const delay = data.code === 1006 ? this.reconnectInterval * 2 : this.reconnectInterval;
			  setTimeout(() => {
				this.reconnect();
			  }, Math.min(delay, 30000)); // 最大延迟 30 秒
			}
		*/
	}

	/**
	* @description: WebSocket连接错误回调
	*/
	private onError(error: any): void {
		console.error('WebSocket 连接错误:', error);
		this.isConnected = false;
		this.status = 'error';
		//this.clearTimers();

		// 触发错误回调
		if (this.eventCallbacks.onError) {
			this.eventCallbacks.onError(error);
		}
	}
	
	/**
	* @description: 处理接收到的消息 - 只负责转发给消息处理器
	*/
	private onMessage(event: any): void {
		try {
			const data = JSON.parse(event.data);
			console.log('收到WebSocket消息:', data);
	  
			// 转发给消息处理器
			if (this.eventCallbacks.onMessage) {
				this.eventCallbacks.onMessage(data);
			}
		} catch (error) {
			console.error('解析WebSocket消息失败:', error);
		}
	}
}

const wsManager = new WsManager();
export default wsManager;
