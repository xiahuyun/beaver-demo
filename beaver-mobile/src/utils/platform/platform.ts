/**
 * 检查当前平台
 */
export const getPlatform = () => {
  // #ifdef H5
  return 'h5';
  // #endif
}
