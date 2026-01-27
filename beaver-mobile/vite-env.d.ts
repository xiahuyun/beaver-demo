/// <reference types="vite/client" />

// 声明图片资源类型
declare module '*.png' {
  const src: string
  export default src
}

declare module '*.jpg' {
  const src: string
  export default src
}

declare module '*.jpeg' {
  const src: string
  export default src
}

declare module '*.gif' {
  const src: string
  export default src
}

declare module '*.svg' {
  const src: string
  export default src
}

// 声明样式文件类型
declare module '*.css'
declare module '*.less'
declare module '*.scss'
declare module '*.sass'

// env.json 类型定义
declare module '@/env.json' {
  const env: {
    baseUrl: string,
	wsUrl: string,
	friendUrl: string
  }
  export default env
  export const baseUrl: string
  export const wsUrl: string
  export const friendUrl: string
}