/**
 * 接口的公共请求头
 */
export const getCommonHeader = () => {
  const token = getLocal('token');
  const deviceId = uni.getStorageSync('uni_device_id');
  return {
    token: token,
    deviceId: deviceId,
    version: appVersion()
  }
}

export const getLocal = (keys: string) =>{
  let value = uni.getStorageSync(keys)
  if (value) {
      return JSON.parse(value)
  } else {
      return value
  }
}

/**
 * 获取APP版本号
 */
export const appVersion = (): string => {
  try {
    // APP环境下，优先使用plus.runtime.version
    if (plus && plus.runtime && plus.runtime.version) {
      return plus.runtime.version;
    }
  } catch (error) {
    console.warn('获取APP版本号失败:', error);
  }  
  return '';
};
