<template>
	<view class="container">
		<!-- 使用通用Header组件 -->
		<BeaverHeader 
			title="消息"
			:show-back="false"
			:right-icon="plusIcon"
			@right-click="showDropdown = !showDropdown"
		></BeaverHeader>
		
		<!-- 下拉菜单 -->
		<view class="dropdown" :class="{ 'show': showDropdown }" :style="{ top: (statusBarHeight + 88) + 'rpx'}">
			<view 
				v-for="menu in homeMenusList"
				:key="menu.id"
				class="dropdown-item"
				@click="handleMenuClick(menu)"
			>
				<view class="dropdown-icon">
					<image :src="menu.icon" mode="aspectFit" class="icon-img"></image>
				</view>
				<text>{{ menu.title }}</text>
			</view>
		</view>
		
		<!-- 遮罩层 -->
		<view class="mask" :class="{ 'show': showDropdown }" @click="showDropdown = false"></view>
	</view>
</template>

<script setup lang="ts">

import { ref } from 'vue';
import plusIcon from '@/static/common/plus-icon.svg';
import { homeMenusList } from './home';

const showDropdown = ref(false);
const statusBarHeight = ref(uni.getSystemInfoSync().statusBarHeight || 0);

// 处理菜单点击
const handleMenuClick = (menu: any) => {
	showDropdown.value = false;
  
	switch (menu.id) {
		case 1: // 添加好友
			uni.navigateTo({
				url: '/pages/searchFriend/searchFriend',
				animationType: 'slide-in-right',
				animationDuration: 200
			})
			break;
		case 2: // 创建群聊
			uni.navigateTo({
				url: '/pages/createGroup/createGroup',
				animationType: 'slide-in-right',
				animationDuration: 200
			})
			break;
		case 3: // 扫一扫
			scanCode();
			break;
		default:
			console.log('未知菜单项:', menu);
	}
}

const scanCode = () => {
	uni.scanCode({
		scanType: ['qrCode'],
		success: (res) => {
			console.log("二维码数据", res);
			handleScanResult(res.result);
		},
		fail: (error) => {
			console.error('扫码失败', error);
		}
	})
}

</script>

<style lang="scss" scoped>

.container {
  background-color: #FFFFFF;
  color: #2D3436;
  font-size: 28rpx;
  line-height: 1.5;
  overflow-x: hidden;
  position: relative;
  min-height: 100vh;
}

/* 遮罩层 */
.mask {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.2);
  z-index: 150;
  display: none;
  opacity: 0;
  transition: opacity 0.3s;
}

/* 下拉菜单 */
.dropdown {
  position: fixed;
  right: 32rpx;
  background: white;
  box-shadow: 0 4rpx 24rpx rgba(0, 0, 0, 0.1);
  border-radius: 24rpx;
  width: 320rpx;
  overflow: hidden;
  z-index: 200;
  display: none;
  opacity: 1;
  transform: translateY(-20rpx);
  transform-origin: top right;
  transition: opacity 0.2s cubic-bezier(0, 0, 0.2, 1), transform 0.2s cubic-bezier(0, 0, 0.2, 1);
}

.dropdown.show {
  display: block;
  opacity: 1;
  transform: translateY(0);
}

.dropdown-item {
  height: 88rpx;
  padding: 0 32rpx;
  display: flex;
  align-items: center;
  color: #2D3436;
  font-weight: 500;
  font-size: 28rpx;
}

.dropdown-icon {
	width: 40rpx;
	height: 40rpx;
	margin-right: 24rpx;
	display: flex;
	align-items: center;
	justify-content: center;
	color: #FF7D45;
}

/* 图标样式 */
.icon-img {
  width: 100%;
  height: 100%;
}

.mask.show {
  display: block;
  opacity: 1;
}

</style>
