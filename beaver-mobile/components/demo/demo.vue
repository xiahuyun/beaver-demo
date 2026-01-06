<template>
	<view class="demo">
		<image :src="APP_CONFIG.logo" mode="aspectFit" />
		<view class="demo">我是 demo 组件</view>
	</view>
</template>

<script setup lang="ts">
import { APP_CONFIG } from '@/config/data';
</script>

<style>

</style>
