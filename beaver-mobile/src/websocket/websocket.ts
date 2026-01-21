/**
 * @description: WebSocket连接管理器 - 纯连接层，不处理业务逻辑
 */
class WsManager {
	constructor() {}
	public async initSocket(): Promise<void> {
		return new Promise((resolve, reject) => {
			try {
				uni.connectSocket({
					url: `ws://127.0.0.1:20802/api/chat/websocket?token=test&userId=1`,
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
}

const wsManager = new WsManager();
export default wsManager;
