<script lang="ts">

import { useInitStore } from '@/src/pinia/init/init';
import { getLocal } from '@/src/utils/local/local';

export default {
	onLaunch: async function() {
		console.log('App Launch')
		
		const initStore = useInitStore();
		
		if (getLocal('token')) {
			await initStore.getAuthentication();
		} else {
			uni.reLaunch({url: '/pages/login/login'});
			return
		}

		await initStore.initApp()
		if (initStore.initError) {
			uni.showToast({
				title: '应用初始化失败 请稍后重试',
				icon: 'error',
				duration: 3000
			});
		}
	},
	onShow: function() {
		console.log('App Show')
	},
	onHide: function() {
		console.log('App Hide')
	}
}
</script>

<style lang="scss">
	/*每个页面公共css */
	@import '@/uni_modules/uni-scss/index.scss';
	/* #ifndef APP-NVUE */
	@import '@/static/customicons.css';
	// 设置整个项目的背景色
	page {
		background-color: #f5f5f5;
	}

	/* #endif */
	.example-info {
		font-size: 14px;
		color: #333;
		padding: 10px;
	}
</style>
