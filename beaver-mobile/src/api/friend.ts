import { request } from '../utils/request/request'
import type { 
	ISearchRes,
	ISearchReq,
	IFriendInfoReq,
	IFriendInfo,
	IAddFriendReq,
	IAddFriendRes,
} from '@/src/types/ajax/friend'

/**
 * @description: 搜索好友
 */
export const searchFriendApi = (data: ISearchReq) => {
	return request<ISearchRes>({
		method: 'GET',
		data: data,
		url: `/api/friend/search`
	})
}

/**
 * @description: 获取好友信息
 */
export const friendInfoApi = (data: IFriendInfoReq) => {
  return request<IFriendInfo>({
    data: data,
    method: 'GET',
    url: `/api/friend/friend_info`
  })
}

/**
 * @description: 添加好友
 */
export const addFriendApi = (data: IAddFriendReq) => {
  return request<IAddFriendRes>({
    data: data,
    method: 'POST',
    url: `/api/friend/add_friend`
  })
}
