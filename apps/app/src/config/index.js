/*
 * *************
 * 是否为测试环境
 *
 * 生产环境: false
 * 测试环境: true
 * *************
 */
export const DEV_ENV = true

// 控制台打印http request详情
export const CONSOLE_REQUEST_ENABLE = false

// axios 默认配置
export const AXIOS_DEFAULT_CONFIG = {
    timeout: 30000,
    maxContentLength: 2000,
    headers: {}
}

// API 默认配置
export const API_DEFAULT_CONFIG = {
    prodBaseURL: '',
    mockBaseURL: '', // 前端和服务端端口号不一致通过devServer解决跨域问题因而URL为空
    mock: DEV_ENV,
    debug: true,
    sep: '/'
}
