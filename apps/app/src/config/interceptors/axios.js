import { CONSOLE_REQUEST_ENABLE, CONSOLE_RESPONSE_ENABLE } from '../index.js'

export function requestSuccessFunc (requestObj) {
    CONSOLE_REQUEST_ENABLE && console.log('requestInterceptorFunc',
        'url: ' + requestObj.url, requestObj)

    // 自定义拦截请求
    //
    requestObj.headers['Access-Control-Allow-Origin'] = '*'
    return requestObj
}

export function requestFailFunc (requestError) {
    // 自定义发送请求失败日志、断网、请求发送监控等
    // ..
    
    return Promise.reject(requestError)
}

export function responseSuccessFunc (responseObj) {
    // 自定义响应成功逻辑，全局拦截接口，根据不同业务做不同处理，响应成功监控等
    const resData = responseObj.data
    let resp = {
        code: resData.code,
        msg: resData.msg,
    }
    if(typeof resData.data == 'string') {
        resp.data = JSON.parse(resData.data)
    } else {
        resp.data = resData.data
    }
    return resp
}

export function responseFailFunc (responseError) {
    // 响应失败，可根据 responseError.message 和 responseError.response.status 来做监控处理
    // ..
    
    return Promise.reject(requestError)
}
