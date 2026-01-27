<template>
	<image
		:src="imageSrc"
		:mode="mode"
		:lazy-load="lazyLoad"
		:class="imageClass"
		:style="imageStyle"
		@load="$emit('load', $event)"
		@error="$emit('error', $event)"
		@click="$emit('click', $event)"
	/>
</template>

<script setup lang="ts">
import { ref, watch, computed } from 'vue';
import { getFile} from '@/src/cache/file';

const imageSrc = ref('');

// 检查当前平台
const isH5 = computed(() => {
	// #ifdef H5
	return true;
});

const props = defineProps({
	fileName: {
		type: String,
		required: true,
		default: "heihei"
	},
	mode: {
		type: String,
		default: 'aspectFill'
	},
	lazyLoad: {
		type: Boolean,
		default: true
	},
	imageClass: {
		type: String,
		default: ''
	},
	imageStyle: {
		type: Object,
		default: () => ({})
	}
})

const emit = defineEmits(['load', 'error', 'click']);

const updateImage = async () => {
	if (!props.fileName) {
		imageSrc.value = '';
		return;
	}

	try {
		imageSrc.value = await getFile(props.fileName);
	} catch (error) {
		console.error('加载图片失败:', error);
		// 降级到直接URL
		imageSrc.value = previewOnlineFileApi(props.fileName);
	}
};

// 监听 fileName 变化
watch(() => props.fileName, (newFileName, oldFileName) => {
	if (newFileName !== oldFileName) {
		updateImage();
	}
}, { immediate: true });

</script>

<style>
</style>