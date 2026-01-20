export function getLocal(key: string): string {
	let value = uni.getStorageSync(key)
	if (value) {
		return value
	} else {
		return value
	}
}

export function setLocal(key: string, value: string): void {
	uni.setStorageSync(key, value)
}

export function removeLocal(key: string): void {
	uni.removeStorageSync(key)
}
