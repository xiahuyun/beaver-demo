/**
 * @description: WebSocket连接管理器 - 纯连接层，不处理业务逻辑
 */
class WsManager {
	constructor() {}
	public async initSocket(): Promise<void> {
		return new Promise((resolve, reject) => {
			try {
				uni.connectSocket({
					url: 'https://127.0.0.1:28088/api/ws/ws?token=reqreq',
					method: 'GET',
					success() {
						console.log('WebSocket 连接请求成功');
						resolve();
					},
					fail(error) {
						console.error('WebSocket 连接请求失败:', error);
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
