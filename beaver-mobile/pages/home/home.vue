<template>
	<view class="container">
		<!-- 使用通用Header组件 -->
		<BeaverHeader 
			title="消息"
			:show-back="false"
			:right-icon="plusIcon"
			@right-click="showDropdown = !showDropdown"
		></BeaverHeader>
		
		<!-- 遮罩层 -->
		<view class="mask" :class="{ 'show': showDropdown }" @click="showDropdown = false"></view>
	</view>
</template>

<script setup>
import { ref } from 'vue';
import plusIcon from '@/static/common/plus-icon.svg';

const showDropdown = ref(false);

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

.dropdown.show {
  display: block;
  opacity: 1;
  transform: translateY(0);
}

.mask.show {
  display: block;
  opacity: 1;
}

</style>
