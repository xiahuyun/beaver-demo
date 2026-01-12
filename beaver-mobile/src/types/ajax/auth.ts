// 获取邮箱验证码请求
export interface IGetEmailCodeReq {
	email: string;
	type: string; // 验证码类型：register(注册)，reset(重置密码)，login(登录)
}

// 获取邮箱验证码响应
export interface IGetEmailCodeRes {
	code: number;
	message: string;
}

// 注册用户请求
export interface IRegisterUserReq {
	email: string;
	phone: string;
	password: string;
	verificationCode: string;
	type: string; // 用户类型：email(邮箱注册)，phone(手机注册)
}

// 注册用户响应
export interface IRegisterUserRes {
	message: string;
}

// 用户登录请求
export interface ILoginUserReq {
	email: string;
	phone: string;
	password: string;
	type: string; // 用户类型：email(邮箱注册)，phone(手机注册)
}

// 用户登录响应
export interface ILoginUserRes {
	code: number;
	message: string;
}
