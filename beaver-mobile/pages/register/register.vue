<template>
	<view class="container">
		<view class="top_gradient"></view>
		
		<view class="content">
			<view class="logo-container">
				<view class="log">
					<image :src="APP_CONFIG.logo" mode="aspectFit"/>
				</view>
			</view>
			
			<view class="title">创建账号</view>
			<view class="title_decoration"></view>
			
			<view class="form-container">
				<view class="form-group">
					<input 
						type="email" 
						class="form-input" 
						v-model="registerInfo.email"
						placeholder="邮箱地址"
						@input="inputEmail"
					>
					<view v-if="emailTouched && !isEmailValid" class="error_message">请输入有效邮箱地址</view>
				</view>
				<view class="form-group">
					<input 
						:type="passwordType" 
						class="form-input" 
						placeholder="设置密码"
						v-model="registerInfo.password"
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

				<view class="form-group">
					<input 
						type="text" 
						class="form-input" 
						placeholder="验证码"
						v-model="registerInfo.verificationCode"
						@input="inputVerificationCode"
					>
					<button class="code-btn" :disabled="isCodeButtonDisabled" @click="sendVerificationCode">{{ isCodeButtonDisabled ? countdown + 's' : '获取验证码' }}</button>
					<view v-if="codeBtnTouched && !isEmailValid" class="error_message">请输入有效邮箱</view>
					<view v-if="verificationCodeTouched && !isVerificationCodeValid" class="error_message">验证码为6位数字</view>
				</view>

				<view class="agreement">
					<view class="checkbox-wrapper">
						<view class="custom-checkbox" :class="{ checked: isAgreed }" @click="toggleAgreement"></view>
						<!--<checkbox :checked="isAgreed" @tap="toggleAgreement" style="opacity: 0; position: absolute; top: 0; left: 0; width: 100%; height: 100%;" />-->
						<text class="agreement-text">我已阅读并同意</text>
						<text class="link" @click="navigateToPage('/pages/agreement/agreement')">用户协议</text>
						<text class="agreement-text">和</text>
						<text class="link" @click="navigateToPage('/pages/privacy/privacy')">隐私政策</text>
					</view>
				</view>

				<button 
					class="btn btn-primary" 
					:class="{'btn-disabled': !isFormValid}"
					@click="register"
				>注册
				</button>
					
				<view class="register-link">
					已有账号？<text class="jump-text" @click="navigateToPage('/pages/login/login')">登录</text>
				</view>
			</view>
		</view>
	</view>
</template>

<script setup lang="ts">

import { resolve } from "dns";
import { APP_CONFIG } from "@/config/data";
import { reactive, ref, watch } from "vue";
import { getEmailCodeApi,registerUserApi } from "@/src/api/auth";
import { encodePassword } from "@/src/utils/encode/password";

const emailTouched = ref(false);
const isEmailValid = ref(false);

const passwordType = ref("password");
const passwordTouched = ref(false);
const isPasswordValid = ref(false);

const verificationCodeTouched = ref(false);
const isVerificationCodeValid = ref(false);

const isFormValid = ref(false);

const countdown = ref(60);
const isCodeButtonDisabled = ref(false);

const codeBtnTouched = ref(false);

interface RegisterInfo {
	email: string;
	password: string;
	verificationCode: string;
}

const registerInfo = reactive<RegisterInfo>({
	email: "",
	password: "",
	verificationCode: "",
})

function inputEmail(): void {
	emailTouched.value = true;
}

function register(): void {
	registerUserApi({
		email: registerInfo.email,
		password: encodePassword(registerInfo.password, APP_CONFIG.salt),
		verificationCode: registerInfo.verificationCode,
		type: "email"
	}).then((res) => {
		if (res.code === 200) {
			uni.reLaunch({
				url: "/pages/login/login",
				animationType: 'pop-in',
				animationDuration: 200
			});
		} else {
			uni.showToast({
				title: res.message,
				duration: 3000,
				icon: 'error',
			});
		}
	}).catch(() => {
		uni.showToast({
			title: '注册失败，请重试',
			duration: 3000,
			icon: 'error',
		});
	});
}

function inputPassword(): void {
	passwordTouched.value = true;
}

function inputVerificationCode(): void {
	verificationCodeTouched.value = true;
}

function sendVerificationCode(): void {
	codeBtnTouched.value = true;
	if (!isEmailValid.value) {
		return
	}

	getEmailCodeApi({
		email: registerInfo.email,
		type: "register"
	}).then((res) => {
		if (res.code === 200) {
			uni.showToast({
			  title: '验证码已发送',
			  duration: 2000
			});
			
			isCodeButtonDisabled.value = true;
			let counter = countdown.value;
			const interval = setInterval(() => {
				counter -= 1;
				countdown.value = counter;
				if (counter <= 0) {
					clearInterval(interval);
					isCodeButtonDisabled.value = false;
					countdown.value = 60;
				}
			}, 1000);
		}
		console.log(res);
	}).catch((error) => {
		console.log(error);
	})
}

const isAgreed = ref(false)

function toggleAgreement(): void {
	isAgreed.value = !isAgreed.value;
}

watch([()=>registerInfo.email, ()=>registerInfo.password, ()=>registerInfo.verificationCode, ()=>isAgreed.value], () => {
	isEmailValid.value = validateEmail(registerInfo.email);
	isPasswordValid.value = validatePassword(registerInfo.password);
	isVerificationCodeValid.value = validateVerificationCode(registerInfo.verificationCode);
	isFormValid.value = isEmailValid.value && isPasswordValid.value && isVerificationCodeValid.value && isAgreed.value;
})

function validateEmail(email: string): boolean {
	const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
	return emailRegex.test(email);
}

function validatePassword(password: string): boolean {
	return password.length >= 8 && /[a-z]/.test(password) && /[A-Z]/.test(password) && /[0-9]/.test(password);
}

function validateVerificationCode(verificationCode: string): boolean {
	return verificationCode.length == 6 && /^[0-9]*$/.test(verificationCode);
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

.code-btn {
	width: 210rpx;
	height: 70rpx;
	background: linear-gradient(135deg, #FF7D45 0%, #E86835 100%);
	color: white;
	border-radius: 20rpx;
	padding: 0rpx;
	display: flex;
	align-items: center;
	justify-content: center;
	font-size: 25rpx;

	&:active {
		transform: translateY(-5%) translateY(1px);
		box-shadow: 0 1px 4px rgba(255, 125, 69, 0.1);
	}

	&[disabled] {
		opacity: 0.6;
		background: linear-gradient(135deg, #FF7D45 0%, #E86835 100%);
		color: white;
		box-shadow: 0 2px 8px rgba(255, 125, 69, 0.15);
		cursor: not-allowed;
		pointer-events: none;
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
			opacity: 0.6;
			box-shadow: none;
			pointer-events: none;
			cursor: not-allowed;
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

.agreement {
	display: flex;
	align-items: center;
	font-size: 24rpx;
	color: #636E72;
	margin: 24rpx 0;
}

.checkbox-wrapper {
  position: relative;
  margin-right: 20rpx;
  display: flex;
  align-items: center;
}

.custom-checkbox {
	border: 2rpx solid #B2BEC3;
	position: relative;
	height: 36rpx;
	width: 36rpx;
	display: inline-box;
	background: #FFFFFF;
	cursor: pointer;
	transition: all 0.2s;
	border-radius: 10rpx;
	
	&.checked {
		background: #FF7D45;
		border-color: #FF7D45;
	
		&:after {
			content: '';
			position: absolute;
			left: 12rpx;
			top: 6rpx;
			width: 8rpx;
			height: 16rpx;
			border: solid white;
			border-width: 0rpx 4rpx 4rpx 0rpx;
			transform: rotate(45deg);
		}
	}
}

.agreement-text {
	color: #636E72;
	padding-left: 5rpx;
}

.link {
	color: #FF7D45;
	font-weight: 500;
}

</style>
