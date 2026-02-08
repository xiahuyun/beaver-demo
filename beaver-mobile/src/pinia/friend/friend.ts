import { defineStore } from 'pinia';
import type { IFriendInfo } from '@/src/types/ajax/friend';
import { friendListApi } from '@/src/api/friend';

/**
 * @description: 好友信息管理
 */
export const useFriendStore = defineStore('friendStore', {
	state: (): {
		friendList: IFriendInfo[],
	} => ({
		friendList: [],
	}),
	
	actions: {
		async initFriendApi() {
			try {
				const res = await friendListApi({
					page: 1,
					limit: 1000,
				});
				console.log("init friend result: ", res);
				if (res.code === 0) {
					this.friendList = res.result.list || [];
				}
			} catch (error) {
				console.error('Failed to initialize friend list:', error);
				throw error;
			}
		}
	}
});