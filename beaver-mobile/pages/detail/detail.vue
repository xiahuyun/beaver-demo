<template>
	<BeaverLayout
		:title="pageTitle"
		:show-back="true"
		:scrollable="true"
		:show-background="true"
		:show-scrollbar="true"
		background-type="gradient"
		:background-height="180"
		@back="handleClickGoBack"
	>
		<template #right>
			<view class="more-button" @click="toggleMoreMenu" v-if="isFriend">
				<image src="@/static/detail/more-icon.svg" mode="aspectFit" class="more-icon" />
			</view>
		</template>

		<!-- 内容区域 -->
		<view class="content">
			<!-- 统一的用户信息卡片 -->
			<view class="user-info-card">
				<!-- 用户头像和基本信息 -->
				<view class="user-basic-info">
					<view class="user-avatar">
						<BeaverImage 
							:fileName="userInfo.fileName"
							:image-class="'avatar-img'"
						/>
					</view>
					<view class="user-details">
						<view class="user-name-row">
							<view class="user-name">{{ userInfo.userId || userInfo.nickname }}</view>
							<image v-if="userInfo.gender === 'male'" src="@/static/detail/gender-male-icon.svg" mode="aspectFit" class="gender-icon" />
						</view>
						<view class="user-id">ID: {{ userInfo.userId }}</view>
						<view class="user-signature">{{ userInfo.signature || '这个人很懒，什么都没写~' }}</view>
					</view>
				</view>
		
				<!-- 信息列表 -->
				<view class="info-list">
					<!-- 好友信息（仅好友显示） -->
					<template v-if="isFriend && friendInfoItems.length > 0">
						<view class="info-item" v-for="(item, index) in friendInfoItems" :key="'friend-' + index">
							<view class="info-label">{{ item.label }}</view>
							<view class="info-value">{{ item.value }}</view>
						</view>
					</template>

					<!-- 基本资料 -->
					<view class="info-item" v-for="(item, index) in basicInfoItems" :key="'basic-' + index">
						<view class="info-label">{{ item.label }}</view>
						<view class="info-value">{{ item.value }}</view>
					</view>
				</view>
			</view>

			<!-- 相册预览卡片 -->
			<PhotoCard 
				:photos="userInfo.photos"
				@view-all="viewAllPhotos"
			/>
		</view>
		
		<!-- 底部操作栏 -->
		<template #after>
			<ActionBar 
				:is-friend="isFriend"
				:user-id="userInfo.userId"
				:source="source"
				@send-message="sendMessage"
				@audio-call="audioCall"
				@video-call="videoCall"
				@add-friend-success="handleAddFriendSuccess"
			/>
		</template>
	</BeaverLayout>
</template>

<script setup lang="ts">
import { ref,computed } from 'vue';
import { onLoad } from '@dcloudio/uni-app';
import { userInfo } from 'os';
import { friendInfoApi } from '@/src/api/friend';

const userInfo = ref({
	userId: '',
	nickname: '',
	fileName: '',
	remarkName: '',
	signature: '',
	gender: '',
	location: '',
	age: '',
	constellation: '',
	occupation: '',
	education: '',
	hobbies: '',
	photos: [] as string[],
	conversationId: '',
	source: ''
});

const isFriend = ref(false);
const source = ref(''); // 添加好友来源
const showMoreMenu = ref(false);

// 计算页面标题
const pageTitle = computed(() => {
	return isFriend.value ? '好友资料' : '用户资料';
});

onLoad((option: any) => {
	if (option.id) {
		loadUserInfo(option.id);
	}
	// 记录添加好友来源
	if (option.source) {
		source.value = option.source;
	}
});

const loadUserInfo = async (userId: string) => {
	try {
		const res = await friendInfoApi({ friendId: userId, userId: "xiahuyun043@126.com" });
		if (res.code === 0) {
			userInfo.value = { ...userInfo.value, ...res.result };
			// 使用接口返回的isFriend字段判断好友关系
			isFriend.value = res.result.isFriend || false;
		}
	} catch (error) {
		console.log("获取用户信息失败", error);
		uni.showToast({ title: '获取用户信息失败', icon: 'none' });
	}
}

const handleClickGoBack = () => {
	uni.navigateBack();
};

const toggleMoreMenu = () => {
	showMoreMenu.value = true;
};

// 计算好友信息项
const friendInfoItems = computed(() => {
	const items = [];
	
	// 添加备注
	if (userInfo.value.remarkName) {
		items.push({ label: '备注', value: userInfo.value.remarkName });
	}
	
	// 添加好友来源
	if (userInfo.value.source) {
		items.push({ label: '来源', value: getSourceText(userInfo.value.source) });
	}
	
	return items;
});

// 计算基本资料项
const basicInfoItems = computed(() => {
	const items = [];

	// 添加昵称
	if (userInfo.value.nickname) {
		items.push({ label: '昵称', value: userInfo.value.nickname });
	}
	
	// 添加性别
	items.push({ label: '性别', value: userInfo.value.gender === 'male' ? '男' : userInfo.value.gender === 'female' ? '女' : '未设置' });
	
	return items;
});

const viewAllPhotos = () => {
	uni.showToast({ title: '查看更多照片', icon: 'none' });
};

const handleAddFriendSuccess = () => {
	console.log("add friend finished");
	uni.showToast({ title: '添加好友成功', icon: 'success' });
	isFriend.value = true; // 更新好友状态
	uni.navigateBack(); // 返回上一页
};

const audioCall = () => {
	uni.showToast({ title: '发起语音通话', icon: 'none' });
};

const videoCall = () => {
	uni.showToast({ title: '发起视频通话', icon: 'none' });
};

const sendMessage = () => {
	if (isFriend.value && userInfo.value.conversationId) {
		uni.navigateTo({
			url: `/pages/chat/chat?id=${userInfo.value.conversationId}`,
			animationType: 'slide-in-right',
			animationDuration: 200
		});
	} else {
		uni.showToast({ title: '暂不支持与非好友聊天', icon: 'none' });
	}
};

</script>

<style lang="scss" scoped>
/* 更多按钮样式 */
.more-button {
	width: 44rpx;
	height: 44rpx;
	border-radius: 50%;
	display: flex;
	align-items: center;
	justify-content: center;
	background: rgba(255, 255, 255, 0.9);
	box-shadow: 0 2rpx 8rpx rgba(0,0,0,0.1);
	transition: all 0.2s ease;

	&:active {
		transform: scale(0.95);
	}
}

.more-icon {
	width: 32rpx;
	height: 32rpx;
}

/* 基础样式 */
.content {
	padding: 0 24rpx;
	padding-bottom: 120rpx; /* 为底部操作栏留出空间 */
	max-width: 750rpx;
	box-sizing: border-box;
	margin: 0 auto;
}

.gender-icon {
	width: 36rpx;
	height: 36rpx;
	margin-left: 16rpx;
}

/* 统一用户信息卡片 */
.user-info-card {
	background-color: white;
	border-radius: 24rpx;
	padding: 40rpx 32rpx;
	box-shadow: 0 4rpx 24rpx rgba(0,0,0,0.08);
	margin-bottom: 0;
}

.user-basic-info {
	display: flex;
	flex-direction: column;
	align-items: center;
	text-align: center;
	margin-bottom: 40rpx;
	padding-bottom: 32rpx;
	border-bottom: 1rpx solid #F0F3F4;
}

.user-avatar {
	width: 140rpx;
	height: 140rpx;
	border-radius: 50%;
	overflow: hidden;
	margin-bottom: 24rpx;
	border: 4rpx solid white;
	box-shadow: 0 6rpx 20rpx rgba(0,0,0,0.12);
}

.avatar-img {
	width: 100%;
	height: 100%;
	object-fit: cover;
}

.user-details {
	width: 100%;
}

.user-name-row {
	display: flex;
	align-items: center;
	justify-content: center;
	margin-bottom: 12rpx;
}

.user-name {
	font-size: 40rpx;
	font-weight: 700;
	color: #2D3436;
}

.user-id {
	font-size: 24rpx;
	color: #B2BEC3;
	margin-bottom: 16rpx;
	font-family: 'Monaco', 'Menlo', monospace;
	letter-spacing: 0.5rpx;
}

.user-signature {
	font-size: 28rpx;
	color: #636E72;
	line-height: 1.6;
	padding: 16rpx 20rpx;
	background: #F8F9FA;
	border-radius: 16rpx;
	max-width: 480rpx;
	margin: 0 auto;
}

/* 信息列表 */
.info-list {
	display: flex;
	flex-direction: column;
}

.info-item {
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: 20rpx 24rpx;
	background: linear-gradient(135deg, #F8F9FA 0%, #F0F3F4 100%);
	border-radius: 16rpx;
	border: 1rpx solid transparent;
}

.info-label {
	color: #636E72;
	font-size: 28rpx;
	font-weight: 600;
}

.info-value {
	color: #2D3436;
	font-size: 28rpx;
	font-weight: 700;
	text-align: right;
	max-width: 60%;
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;
}

</style>
