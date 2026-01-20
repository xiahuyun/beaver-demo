import { defineStore } from 'pinia';
import { getLocal } from '@/src/utils/local/local';
import { authenticationApi } from '@/src/api/auth';
import wsManager from '@/src/websocket/websocket'

/**
 * @description: 应用初始化和生命周期管理
 */
export const useInitStore = defineStore('useInitStore', {
	state: () => ({
		/**
		* @description: 是否正在初始化
		*/
		isInitializing: false,

		/**
		* @description: 是否已完成初始化
		*/
		isInitialized: false,

		/**
		* @description: 初始化错误信息
		*/
		initError: null as Error | null,
	}),

	getters: {
		/**
		* @description: 获取初始化状态
		* @return {Object} 初始化状态对象
		*/
		initStatus: (state): object => ({
			hasError: !!state.initError,
		}),
	},

	actions: {
		/**
		* @description: 初始化应用
		* @return {Promise<void>}
		* @throws {Error} 当初始化失败时抛出错误
		*/
		async initApp(): Promise<void> {
			if (this.isInitializing || this.isInitialized) {
				return;
			}
			
			this.isInitializing = true;
			this.initError = null;
			
			try {
				if (!getLocal('token')) {
					throw new Error('No token found');
				}
				
				// 初始化认证
				await this.getAuthentication();
				
				// 初始化 WebSocket 连接
				await wsManager.initSocket();
				
				
			} catch (error) {
				this.initError = error instanceof Error ? error : new Error('App init failed');
			} finally {
				console.log("ok");
				this.isInitializing = false;
				this.isInitialized = true;
			}
		},

		/**
		* @description: 获取认证状态
		* @return {Promise<void>}
		*/
		async getAuthentication(): Promise<void> {
			await authenticationApi();
		}
	}
})