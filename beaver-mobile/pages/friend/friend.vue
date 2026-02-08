<template>
	<view class="container">
		 <BeaverHeader 
			title="好友"
			:show-back="false"
		/>
		
		<!-- 索引栏 -->
		<view class="index-bar">
			<view 
				v-for="(item, index) in indexList" 
				:key="index" 
				class="index-item" 
				:class="{ 'active': currentIndex === item }"
				@click="scrollToSection(item)"
			>
				{{ item }}
			</view>
		</view>
	</view>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useFriendStore } from '@/src/pinia/friend/friend';

const friendStore = useFriendStore();

const friendList = computed(() => friendStore.friendList);

const indexList = computed(() => {
	const letters = ['↑'];
	// 从好友列表中获取所有的首字母并去重
	const uniqueLetters = [...new Set(
		friendList.value.map(friend => {
		// 获取名称的首字母并转为大写
		const firstChar = friend.nickname.charAt(0).toUpperCase();
		// 判断是否为英文字母
		return /[A-Z]/.test(firstChar) ? firstChar : '#';
	})
	)].sort();

	return [...letters, ...uniqueLetters];
});

</script>

<style lang="scss" scoped>

.container {
  background-color: #FFFFFF;
  color: #2D3436;
  font-size: 14px;
  line-height: 1.5;
  overflow-x: hidden;
  position: relative;
  min-height: 100vh;
}

/* 索引栏 */
.index-bar {
  position: fixed;
  top: 50%;
  right: 5px;
  transform: translateY(-50%);
  display: flex;
  flex-direction: column;
  z-index: 90;
}

</style>
