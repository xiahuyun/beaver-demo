import { v4 as uuidv4 } from 'uuid';
import { getCommonHeader } from './common';

export interface IResponseSuccessData<T> {
  code: number
  message: string
  token: string
  result: T
}

export interface IResponseErrorData {
  code: number
  message: string
  token: string
}

export interface IRequestConfig {
  method: 'GET' | 'POST' | 'DELETE'
  url: string
  data?: object
  header?: object
}

export const request = <T> (config: IRequestConfig): Promise<IResponseSuccessData<T>> => {
	return new Promise((resolve, reject) => {
		const requestId = uuidv4();
		
		const requestConfig  = {
			requestId,
			...config?.header || {},
			'content-type': 'application/json;charset=UTF-8', 
			...getCommonHeader(),
			'Beaver-User-Id': 'xiahuyun043@126.com',
		}

		uni.request({
			method: config.method,
			url: config.url,
			data: config.data,
			header: requestConfig,
			success: (res: any) => {
				if (res.statusCode === 200) {
					const data = res.data as IResponseSuccessData<T>;
					data.code = 0;
					resolve(data);
					console.log("result success: ", res);
				} else {
					resolve({
						code: -1,
						message: res.data,
					} as IResponseSuccessData<T>);
					console.log("result failed: ", res);
				}
			},
			fail: (err: UniApp.GeneralCallbackResult) => {
				console.log("request fail with data: ", err);
				reject({
					code: -1,
					message: err.errMsg
				} as IResponseErrorData);
			}
		})
	})
}