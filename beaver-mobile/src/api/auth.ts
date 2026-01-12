import type { 
	IGetEmailCodeReq,
	IGetEmailCodeRes, 
	IRegisterUserReq, 
	IRegisterUserRes,
	ILoginUserReq,
	ILoginUserRes,
} from '@/src/types/ajax/auth'
import { request } from '../utils/request/request'

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
