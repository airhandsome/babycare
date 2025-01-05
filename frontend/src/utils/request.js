import axios from 'axios'

// 创建axios实例
const request = axios.create({
    baseURL: 'http://localhost:8080/api/v1',  // API的base_url
    timeout: 15000,  // 请求超时时间
    headers: {
        'Content-Type': 'application/json',
        'X-Client-Type': 'web-frontend',
    },
    withCredentials: true,
    // 自定义参数序列化
    paramsSerializer: {
        serialize: (params) => {
            const searchParams = new URLSearchParams()
            for (const key in params) {
                searchParams.append(key, params[key])
            }
            return searchParams.toString()
        }
    }
})

// 请求拦截器
request.interceptors.request.use(
    config => {
        // 移除URL末尾的斜杠
        if (config.url.endsWith('/')) {
            config.url = config.url.slice(0, -1)
        }
        return config
    },
    error => {
        console.error('Request error:', error)
        return Promise.reject(error)
    }
)

// 响应拦截器
request.interceptors.response.use(
    response => {
        return response.data
    },
    error => {
        console.error('Response error:', error)
        return Promise.reject(error)
    }
)

export default request 