<template>
	<view class="container">
		<view class="top_gradient"></view>
		
		<view class="content">
			<view class="logo-container">
				<view class="log">
					<image :src="APP_CONFIG.logo" mode="aspectFit"/>
				</view>
			</view>
			
			<view class="title">欢迎回来</view>
			<view class="title_decoration"></view>
			
			<view class="welcome_text">
				登录您的{{APP_CONFIG.name}}账号，开启社交新体验
			</view>
			
			<view class="form-container">
				<view class="form-group">
					<input 
						type="email" 
						class="form-input" 
						v-model="userInfo.email"
						placeholder="邮箱地址"
						@input="inputEmail"
					>
					<view v-if="emailTouched && !isEmailValid" class="error_message">请输入有效邮箱地址</view>
				</view>
				<view class="form-group">
					<input 
						:type="passwordType" 
						class="form-input" 
						placeholder="登录密码"
						v-model="userInfo.password"
						@input="inputPassword"
					>
					<image 
						class="icon_password"
						mode="aspectFit"
						:src="passwordType === 'password' ? '/static/login/eye.svg': '/static/login/eye-slash.svg'"
						@click="togglePasswordVisibility"
					/>
					<view v-if="passwordTouched && !isPasswordValid" class="error_message">密码长度不少于8位，且必须包含大小写和数字</view>
				</view>
				
				<view class="forgot-password">
					<text class="jump-text" @click="navigateToPage('/pages/forget/forget')">忘记密码?</text>
				</view>

				<button 
					class="btn btn-primary" 
					:class="{'btn-disabled': !isFormValid}"
					@click="goHome"
				>登录
				</button>
				
				<view class="register-link">
					还没有账号？<text class="jump-text" @click="navigateToPage('/pages/register/register')">立即注册</text>
				</view>
			</view>
		</view>
	</view>
</template>

<script setup lang="ts">
import { resolve } from "dns";
import { APP_CONFIG } from "@/config/data";
import { reactive, ref, watch } from "vue";
import { loginUserApi } from "@/src/api/auth";
import { encodePassword } from "@/src/utils/encode/password";
import { setLocal } from "@/src/utils/local/local";

const emailTouched = ref(false);
const isEmailValid = ref(false);

const passwordType = ref("password");
const passwordTouched = ref(false);
const isPasswordValid = ref(false);

const isFormValid = ref(false);

interface UserInfo {
	email: string;
	password: string;
}

const userInfo = reactive<UserInfo>({
	email: "",
	password: ""
})

function inputEmail(): void {
	emailTouched.value = true;
}

function inputPassword(): void {
	passwordTouched.value = true;
}

watch([() => userInfo.email, () => userInfo.password], () => {
	isEmailValid.value = validateEmail(userInfo.email);
	isPasswordValid.value = validatePassword(userInfo.password);
	isFormValid.value = isEmailValid.value && isPasswordValid.value;
})

function validateEmail(email: string): boolean {
	const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
	return emailRegex.test(email);
}

function validatePassword(password: string): boolean {
	return password.length >= 8 && /[a-z]/.test(password) && /[A-Z]/.test(password) && /[0-9]/.test(password);
}

function togglePasswordVisibility(): void {
	passwordType.value = passwordType.value === 'password' ? 'text': 'password';
}

function navigateToPage(url: string): void {
	uni.navigateTo({
		url: url,
		animationType: 'slide-in-right',
		animationDuration: 200
	})
}

function goHome(): void {
	loginUserApi({
		email: userInfo.email,
		password: encodePassword(userInfo.password, APP_CONFIG.salt)
	}).then((res) => {
		if (res.code === 200) {
			console.log(res);
			uni.reLaunch({
				url: "/pages/home/home",
				animationType: 'pop-in',
				animationDuration: 200
			});
			
			setLocal('token', res.token);
			console.log("token", res.token);
		} else {
			console.log("err", res)
			if (res == '' || res.message == 'undefined') {
				uni.showToast({
					title: '登录失败',
					duration: 3000,
					icon: 'error',
				});
			} else {
				uni.showToast({
					title: res.message,
					duration: 3000,
					icon: 'error',
				});
			}
		}
	}).catch(() => {
		uni.showToast({
			title: '登录失败，请重试',
			duration: 3000,
			icon: 'error',
		});
	});
}

</script>

<style lang="scss" scoped>

.container {
	min-height: 100vh;
	background-color: #FFFFFF;
	position: relative;
}

.top_gradient {
	height: 240rpx;
	background: linear-gradient(180deg, rgba(255,125,69,0.1) 0%, rgba(255,255,255,0) 100%);
}

.content {
	padding: 0 100rpx;
	max-width: 750rpx;
	margin: 0 auto;
}

.logo-container {
	display: flex;
	justify-content: center;
	margin-bottom: 48rpx;
}

.log {
	width: 112rpx;
	height: 112rpx;
	display: flex;
	position: relative;
	justify-content: center;
	// TODO: 为什么加上这个属性就在 content 盒子里？
	align-items: center;
	box-shadow: 0 8rpx 24rpx rgba(255, 125, 69, 0.2);
	border-radius: 32rpx;
	background: linear-gradient(135deg, #FF7D45 0%, #E86835 100%);
	
	&:after {
	content: '';
	position: absolute;
	top: 0;
	left: 0;
	right: 0;
	height: 50%;
	background: linear-gradient(180deg, rgba(255,255,255,0.2) 0%, rgba(255,255,255,0) 100%);
	border-radius: 32rpx 32rpx 0 0;
	}
}

.title {
	text-align: center;
	font-size: 48rpx;
	font-weight: 700;
	color: #2d3436;
	line-height: 1.3;
	margin-bottom: 16rpx;
}

.title_decoration {
	width: 40rpx;
	height: 4rpx;
	background-color: #FF7D45;
	margin: 0 auto 48rpx;
}

.welcome_text {
	font-size: 28rpx;
	text-align: center;
	color: #636E72;
	margin-bottom: 64rpx;
}

.form-container {
	margin-top: 20rpx;
}

.form-group {
	position: relative;
	display: flex;
	border-radius: 28rpx;
	height: 96rpx;
	box-shadow: inset 0 2rpx 6rpx rgba(0,0,0,0.05);
	align-items: center;
	margin-bottom: 34rpx;
	background-color: #F9FAFB;
	padding: 0 32rpx;
}

.form-input {
	border: none;
	color: #2D3436;
	padding: 0;
	font-size: 30rpx;
	background: transparent;
	width: 100%;
	height: 100%;
}

.error_message {
	color: #FF7D45;
	font-size: 24rpx;
	position: absolute;
	padding-left: 32rpx;
	margin-top: 8rpx;
	bottom: -32rpx;
	left: 0;
}

.icon_password {
	width: 40rpx;
	height: 40rpx;
	padding: 0 16rpx;
	filter: brightness(0) saturate(100%) invert(67%) sepia(8%) saturate(123%) hue-rotate(169deg) brightness(89%) contrast(86%);
}

.forgot-password {
	text-align: right;
	margin-bottom: 48rpx;
	
	.jump-text {
		color: #FF7D45;
		font-size: 24rpx;
		font-weight: 500;
	}
}

.btn {
	width: 100%;
	height: 96rpx;
	border-radius: 28rpx;
	border: none;
	font-weight: 500;
	cursor: pointer;
	transition: all 0.2s cubic-bezier(0.33, 1, 0.68, 1);
	display: flex;
	justify-content: center;
	align-items: center;
	position: relative;
	font-size: 32rpx;

	&.btn-primary {
		background: linear-gradient(135deg, #FF7D45 0%, #E86835 100%);
		color: white;
		box-shadow: 0 8rpx 24rpx rgba(255, 125, 69, 0.15);

		&.btn-disabled {
			background: linear-gradient(135deg, #FF7D45 0%, #E86835 100%);
			box-shadow: none;
			pointer-events: none;
			cursor: not-allowed;
			opacity: 0.6;
		}

		&:hover {
			transform: translateY(-4rpx);
			box-shadow: 0 12rpx 32rpx rgba(255, 125, 69, 0.2);
		}

		&:active {
			transform: translateY(2rpx);
			box-shadow: 0 4rpx 16rpx rgba(255, 125, 69, 0.1);
		}
	}
}

.register-link {
	text-align: center;
	margin-top: 48rpx;
	color: #636E72;
	font-size: 28rpx;
	
	.jump-text {
		color: #FF7D45;
		font-weight: 500;
	}
}
</style>
