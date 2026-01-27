import { getPlatform } from '@/src/utils/platform/platform';
import { previewOnlineFileApi } from '@/src/api/file';

// 图片缓存映射 
const imageCache = new Map<string, string>();

// 文件获取Promise映射 - 防止重复下载/预加载
const filePromises = new Map<string, Promise<string>>();

/**
 * 获取图片（主要接口）
 */
export async function getImage(fileName: string): Promise<string> {
	const platform = getPlatform();
	
	// 检查缓存
	if (imageCache.has(fileName)) {
		const cachedPath = imageCache.get(fileName)!;
		return cachedPath;
	}
	
	  
	// 检查是否正在获取中
	if (filePromises.has(fileName)) {
		console.log('图片正在获取中，等待完成:', fileName);
		return await filePromises.get(fileName)!;
	}

	// 创建获取Promise
	const getPromise = (async () => {
		try {
			if (platform === 'h5') {
				// H5环境：直接返回URL
				const url = previewOnlineFileApi(fileName);
				imageCache.set(fileName, url);
				return url;
			} else {
				// APP环境：下载到本地
				// const localPath = await downloadFile(fileName);
				//imageCache.set(fileName, localPath);
				console.log("download file to local");
				return "localPath";
			}
		} catch (error) {
			console.error('获取图片失败:', error);
			return previewOnlineFileApi(fileName);
		}
	})();
	
	// 保存Promise
	filePromises.set(fileName, getPromise);

	try {
		const result = await getPromise;
		return result;
	} finally {
		// 清理Promise
		filePromises.delete(fileName);
	}
}
