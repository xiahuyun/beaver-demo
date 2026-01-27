<template>
	<view class="page-layout">
		<view v-if="showBackground" class="background-layer" :class="backgroundType" :style="{
			height: backgroundHeight + 'rpx',
			pointerEvents: 'none'
		}"></view>
		
		<!-- Header 组件 -->
		<BeaverHeader v-if="showHeader" :mode="headerMode" :title="title" :left-icon="leftIcon" :right-icon="rightIcon"
			:show-back="showBack" :background="headerBackground" @back="handleGoBack">
			<template #left v-if="$slots.left">
				<slot name="left"></slot>
			</template>
			<template #right v-if="$slots.right">
				<slot name="right"></slot>
			</template>
		</BeaverHeader>

		<!-- 前置内容区域 -->
		<view v-if="$slots.before" class="before-content" :style="beforeContentStyle">
			<slot name="before"></slot>
		</view>
		
		<!-- 内容区域 -->
		<view class="content-wrapper" >
			<!-- 滚动内容 -->
			<scroll-view class="scroll-content" scroll-y="true" scroll-x="false"
				scroll-top="0" scroll-left="0" scroll-into-view=""
			    scroll-with-animation="true" enable-back-to-top="true"
			    show-scrollbar="false" :style="scrollContentStyle">
			    <slot></slot>
			</scroll-view>
		</view>
		<!-- 后置内容区域 -->
		<view v-if="$slots.after" class="after-content" :style="afterContentStyle">
		  <slot name="after"></slot>
		</view>
	</view>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import arrowBackIcon from '@/static/common/arrow-back.svg';
const props = defineProps({
	showBackground: {
		type: Boolean,
		default: false
	},
	backgroundType: {
		type: String as () => 'gradient' | 'solid' | 'none',
		default: 'gradient'
	},
	backgroundHeight: {
		type: Number,
		default: 240 // 120px -> 240rpx
	},
	showHeader: {
		type: Boolean,
		default: true
	},
	headerMode: {
		type: String as () => 'fixed' | 'static' | 'transparent',
		default: 'static'
	},
	title: {
		type: String,
		default: ''
	},
	leftIcon: {
		type: String,
		default: arrowBackIcon
	},
	rightIcon: {
		type: String,
		default: ''
	},
	showBack: {
		type: Boolean,
		default: true
	},
	headerBackground: {
		type: String,
		default: 'transparent'
	},
	beforeHeight: {
		type: Number,
		default: 0
	},
	afterHeight: {
		type: Number,
		default: 0
	}
})

const emit = defineEmits(['back', 'right-click']);
const handleGoBack = () => {
	emit('back');
};

// 前置内容样式
const beforeContentStyle = computed(() => {
	return {
		height: props.beforeHeight + 'rpx'
	};
});

// 滚动内容样式
const scrollContentStyle = computed(() => {
});

// 后置内容样式
const afterContentStyle = computed(() => {
	return {
	};
});

</script>

<style lang="scss" scoped>
.page-layout {
	height: 100vh;
	position: relative;
	background-color: #F9FAFB;
	display: flex;
	flex-direction: column;
	overflow: hidden;
}

/* 背景装饰层 */
.background-layer {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  z-index: 0;
  pointer-events: none;

  &.gradient {
    background: linear-gradient(180deg, rgba(255, 125, 69, 0.1) 0%, rgba(255, 255, 255, 0) 100%);
  }

  &.solid {
    background-color: #F8F9FA;
  }
}

/* 内容包装器 */
.content-wrapper {
	flex: 1;
	position: relative;
	z-index: 1;
}

/* 滚动内容 */
.scroll-content {
  width: 100%;
  box-sizing: border-box;
}

/* 后置内容区域 */
.after-content {
  flex-shrink: 0;
}

</style>