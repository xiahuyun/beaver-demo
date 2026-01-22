import type { QRCodeData } from '@/src/types/utils/qrcode/qrcode';

/**
 * 处理扫码结果并导航
 * @param qrContent 二维码内容
 */
export const handleScanResult = (qrContent: string) => {
	try {
		const data: QRCodeData = JSON.parse(qrContent);
		console.log("二维码数据", data);
		if (!data.action) {
			console.error("无效的二维码格式");
			return;
		}
		
		if (data.expireAt !== -1 && Date.now() > data.timestamp + data.expireAt * 60 * 1000) {
			console.error('二维码已过期', 2000, 'none');
			return;
		}
		
		switch (data.action) {
			case 'addFriend': {
				console.log("add friend");
				break;
			}
			case 'joinGroup': {
				console.log("join group");
				break;
			}
			default:
				console.error("未知的二维码类型");
		}
	} catch (error) {
		console.error("二维码异常", error);
	}
};
