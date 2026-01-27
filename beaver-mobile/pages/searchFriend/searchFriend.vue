<template>
	<view class="search-friend-page">
		<BeaverLayout
			title="添加好友"
			:show-background="false"
			@back="handleClickGoBack"
		>
			<!-- 主内容区域 -->
			<view class="main-content">
				<!-- 搜索区域 -->
				<view class="search-container">
					<view class="search-form">
						<view class="search-icon">
							<image src="@/static/searchFriend/search-icon.svg" mode="aspectFit" class="icon-img" />
						</view>
						<input type="text" class="search-input" v-model="searchQuery" placeholder="输入邮箱搜索好友"
						placeholder-class="search-placeholder" />
						<view class="search-button" @click="performSearch">
							<text>搜索</text>
						</view>
					</view>
				</view>
				
				<!-- 初始状态（未搜索） -->
				<view class="state-empty" :class="{ 'active': !showResult }">
					<!-- 添加方式 -->
					<view class="add-methods">
						<view class="method-item" @click="scanCode">
							<view class="method-icon">
								<image src="@/static/searchFriend/qr-code-icon.svg" mode="aspectFit" class="icon-img" />
							</view>
							<view class="method-info">
								<text class="method-name">扫一扫</text>
								<text class="method-desc">扫描好友的二维码添加</text>
							</view>
							<view class="arrow-right">
								<image src="@/static/searchFriend/arrow-right-icon.svg" mode="aspectFit" class="icon-img" />
							</view>
						</view>
					</view>
				</view>
				
				<!-- 搜索结果状态 -->
				<view class="state-result" :class="{ 'active': showResult }">
					<view class="search-result show fade-in">
						<view class="result-card" @click="goToDetail">
							<view class="user-profile">
								<view class="user-avatar">
									<BeaverImage :fileName="searchResult.avatar" mode="aspectFill" />
								</view>
								<view class="user-info">
									<text class="user-name">{{ searchResult.email }}</text>
									<!--<text class="user-id">ID: {{ searchResult.userId }}</text>-->
								</view>
								
							</view>
							<view class="view-detail-hint">
								<text>点击查看详情</text>
								<image src="@/static/searchFriend/arrow-right-icon.svg" mode="aspectFit" class="arrow-icon" />
							</view>
						</view>
					</view>
				</view>
			</view>
		</BeaverLayout>
	</view>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import type { ISearchRes } from '@/src/types/ajax/friend';
import { searchFriendApi } from '@/src/api/friend';

const handleClickGoBack = () => {
	uni.navigateBack();
};

const searchQuery = ref('');
const showResult = ref(false);
const searchResult = ref<ISearchRes>({
	name: '',
	nickname: '',
	gender: '',
	avatar: '',
	email: ''
});

// 邮箱格式验证
const validateEmail = (email: string): boolean => {
	const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
	return emailRegex.test(email);
};

const performSearch = async () => {
	if (searchQuery.value.length > 0) {
		if (!validateEmail(searchQuery.value)) {
			uni.showToast({title: '请输入正确的邮箱格式', duration: 3000,icon: 'none'});
			return;
		}

		try {
			const res = await searchFriendApi({email: searchQuery.value});
			if (res.code === 0) {
				searchResult.value = res.result;
				showResult.value = true;
				uni.showToast({title: '搜索成功', duration: 3000, icon: 'none'});
			} else {
				console.log("result: ", res);
				uni.showToast({title: '未找到相关用户', duration: 3000, icon: 'none'});
			}
		} catch (error) {
			console.error("搜索失败", error);
			uni.showToast({title: '搜索失败，请稍后再试', duration: 3000, icon: 'none'});
		}
	} else {
		uni.showToast({title: '请输入邮箱地址', duration: 3000, icon: 'none'});
	}
}

const scanCode = () => {
	uni.showToast({title: '打开相机扫描', duration: 3000, icon: 'none'});
	uni.scanCode({
		success: (res) => {
		try {
			const data = JSON.parse(res.result);
			uni.showToast({title: '扫描成功', duration: 3000, icon: 'none'});
		} catch (e) {
			uni.showToast({title: '无效的二维码', duration: 3000, icon: 'none'});
		}
	}});
};

const goToDetail = () => {
		uni.navigateTo({
		url: `/pages/detail/detail?id=${searchResult.value.email}&source=email`
	});
};

</script>

<style lang="scss" scoped>
.search-friend-page {
  min-height: 100vh;
  background-color: #FFFFFF;
}

.main-content {
	box-sizing: border-box;
}

/* 搜索区域 */
.search-container {
  padding: 16rpx 20rpx;
  margin-bottom: 20rpx;
}

.search-form {
  position: relative;
  display: flex;
  align-items: center;
  background: #FFFFFF;
  border-radius: 48rpx;
  box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.08);
  overflow: hidden;
  border: 1rpx solid rgba(255, 125, 69, 0.1);
}

.search-icon {
  position: absolute;
  left: 24rpx;
  top: 50%;
  transform: translateY(-50%);
  width: 36rpx;
  height: 36rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 10;
  color: #FF7D45;
}

.method-icon .icon-img {
  width: 40rpx;
  height: 40rpx;
}

.search-input {
  flex: 1;
  height: 88rpx;
  background-color: transparent;
  padding: 0 32rpx 0 72rpx;
  font-size: 30rpx;
  color: #2D3436;
  box-sizing: border-box;
}

.search-placeholder {
  color: #B2BEC3;
}

.search-button {
  height: 88rpx;
  background: linear-gradient(135deg, #FF7D45 0%, #E86835 100%);
  padding: 0 32rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: 600;
  font-size: 30rpx;
  position: relative;
  overflow: visible;
}

/* 状态类 */
.state-empty,
.state-result {
  display: none;
}

.state-empty.active,
.state-result.active {
  display: block;
}

/* 添加方式卡片 */
.add-methods {
  background-color: #FFFFFF;
  border-radius: 20rpx;
  overflow: hidden;
  margin: 0 20rpx;
  box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.06);
}

.method-item {
  display: flex;
  align-items: center;
  padding: 32rpx 24rpx;
  position: relative;
}

.method-icon {
  width: 80rpx;
  height: 80rpx;
  border-radius: 40rpx;
  background-color: #FFE6D9;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 20rpx;
  flex-shrink: 0;
}

.method-info {
  flex: 1;
}

.method-name {
  font-size: 32rpx;
  font-weight: 600;
  color: #2D3436;
  margin-bottom: 8rpx;
  display: block;
}

.method-desc {
  font-size: 26rpx;
  color: #636E72;
  display: block;
}

.arrow-right {
  width: 28rpx;
  height: 28rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 搜索结果 */
.search-result {
  background-color: #FFFFFF;
  border-radius: 20rpx;
  overflow: hidden;
  box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.08);
  margin: 0 20rpx;
}

.fade-in {
  animation: fadeIn 0.3s ease-out;
}

.result-card {
  padding: 36rpx 24rpx;
}

.user-profile {
  display: flex;
  align-items: flex-start;
}

.user-avatar {
  width: 88rpx;
  height: 88rpx;
  border-radius: 64rpx;
  overflow: hidden;
  background-color: #F9FAFB;
  margin-right: 24rpx;
  box-shadow: 0 6rpx 20rpx rgba(0, 0, 0, 0.1);
  flex-shrink: 0;
}

.user-info {
  flex: 1;
}

.user-name {
  font-size: 36rpx;
  font-weight: 600;
  color: #2D3436;
  margin-bottom: 10rpx;
  display: block;
}

.arrow-icon {
  width: 24rpx;
  height: 24rpx;
  margin-left: 8rpx;
}

.view-detail-hint {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24rpx 0;
  color: #FF7D45;
  font-size: 28rpx;
  font-weight: 500;
}

</style>
