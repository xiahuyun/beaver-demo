<template>
	<view>
		<view class="bottom-bar">
			<view class="actions-container">
				<!-- 好友模式 -->
				<template v-if="isFriend">
				</template>

				<!-- 非好友模式 -->
				<template v-else>
					<view class="action-button primary-button" @click="showAddFriendDialog">
						<image src="@/static/detail/add-friend-icon.svg" mode="aspectFit" class="button-icon" />
						<text>添加好友</text>
					</view>
				</template>
			</view>
		</view>
		
		<BeaverDialog
			v-model="showDialog"
			title="添加好友"
			content="请输入验证消息"
			type="default"
			size="medium"
			confirm-text="发送请求"
			cancel-text="取消"
			@confirm="handleConfirmAddFriend"
			@cancel="handleCancelAddFriend"
		>
			<view class="verify-input-container">
				<input 
					v-model="verifyMessage" 
					class="verify-input" 
					placeholder="请输入验证消息" 
					maxlength="50"
				/>
				<text class="input-counter">{{ verifyMessage.length }}/50</text>
			</view>
		</BeaverDialog>
	</view>
</template>

<script setup lang="ts">

import { ref } from 'vue';
import { addFriendApi } from '@/src/api/friend';

const props = defineProps({
	isFriend: {
		type: Boolean,
		default: false
	},
	source: {
		type: String,
		default: ''
	}
})

const showDialog = ref(false);

const showAddFriendDialog = () => {
	showDialog.value = true;
};

const verifyMessage = ref('你好，我想添加你为好友');

const handleCancelAddFriend = () => {
	showDialog.value = false;
	verifyMessage.value = '你好，我想添加你为好友'; // 重置为默认值
};

const emit = defineEmits(['add-friend-success']);

const handleConfirmAddFriend = async () => {
	try {
		console.log("props", props);
		console.log("verify message: ", verifyMessage);
		
		const res = await addFriendApi({
			friendId: props.userId,
			verify: verifyMessage.value,
			source: props.source
		});
		
		if (res.code === 0) {
			uni.showToast({ title: '好友请求发送成功', icon: 'success' });
			emit('add-friend-success');
			showDialog.value = false;
		} else {
			uni.showToast({ title: res.msg || '添加失败', icon: 'none' });
		}
	} catch (error) {
		uni.showToast({ title: '添加失败，请稍后再试', icon: 'none' });
	}
};

</script>

<style lang="scss" scoped>
/* 底部固定按钮 */
.bottom-bar {
	background-color: white;
	padding: 24rpx 32rpx 48rpx;
}

.actions-container {
	display: flex;
}

.action-button {
	height: 96rpx;
	border-radius: 20rpx;
	display: flex;
	align-items: center;
	justify-content: center;
	transition: all 0.2s ease;
	cursor: pointer;
}

.primary-button {
	background: linear-gradient(135deg, #FF7D45 0%, #E86835 100%);
	color: white;
	box-shadow: 0 4rpx 16rpx rgba(255, 125, 69, 0.2);
	font-size: 30rpx;
	font-weight: 600;
	flex: 3;
	display: flex;
	align-items: center;
	justify-content: center;

	&:active {
		transform: translateY(2rpx);
		box-shadow: 0 2rpx 8rpx rgba(255, 125, 69, 0.15);
	}

	text {
		color: white;
		font-weight: 600;
		margin-left: 12rpx;
	}
}

.button-icon {
	width: 40rpx;
	height: 40rpx;
}

/* 验证消息输入框样式 */
.verify-input-container {
	width: 100%;
	position: relative;
}

.input-counter {
	position: absolute;
	right: 24rpx;
	bottom: 20rpx;
	font-size: 22rpx;
	color: #B2BEC3;
}

.verify-input {
	width: 100%;
	height: 80rpx;
	padding: 0 24rpx;
	border-radius: 16rpx;
	border: 2rpx solid #EBEEF5;
	background-color: #F9FAFB;
	font-size: 28rpx;
	color: #2D3436;
	box-sizing: border-box;
	transition: all 0.2s ease;

	&:focus {
		outline: none;
		border-color: #FF7D45;
		box-shadow: 0 0 0 4rpx rgba(255, 125, 69, 0.1);
		background-color: white;
	}
}

</style>