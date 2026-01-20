import type { 
	IGetEmailCodeReq,
	IGetEmailCodeRes, 
	IRegisterUserReq, 
	IRegisterUserRes,
	ILoginUserReq,
	ILoginUserRes,
} from '@/src/types/ajax/auth'
import { request } from '../utils/request/request'
import { getLocal,removeLocal } from '../utils/local/local'

/**
 * @description: 获取邮箱验证码
 */
export const getEmailCodeApi = (data: IGetEmailCodeReq) => {
  return request<IGetEmailCodeRes>({
    method: 'POST',
    data,
    url: `/api/auth/emailcode`
  })
}

/**
 * @description: 注册用户
 */
export const registerUserApi = (data: IRegisterUserReq) => {
  return request<IRegisterUserRes>({
    method: 'POST',
    data,
    url: `/api/auth/register`
  })
}

/**
 * @description: 用户登录
 */
export const loginUserApi = (data: ILoginUserReq) => {
  return request<ILoginUserRes>({
    method: 'POST',
    data,
    url: `/api/auth/login`
  })
}

/**
 * @description: 用户登录
 */
export const resetPasswordApi = (data: ILoginUserReq) => {
  return request<ILoginUserRes>({
    method: 'POST',
    data,
    url: `/api/auth/resetpassword`
  })
}

/**
 * @description: 用户认证
 */
export const authenticationApi = () => {
	return new Promise((resolve, reject) => {
		const token = getLocal('token');
		
		uni.request({
			method: 'GET',
			url: `/api/auth/authentication`,
			header: {
				'token': token || '',
			},
			success: (res: any) => {
				const data = res.data;
				if (data.code === 1) {
					removeLocal('token');
				}
				
				resolve(data);
			},
			fail: (err: UniApp.GeneralCallbackResult) => {
				console.error('Network error during authentication:', err);
				reject(new Error('Network error'));
			}
		})
	})
}