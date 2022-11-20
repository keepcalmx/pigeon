import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios'
import { useRouter } from 'vue-router'
import { baseURL } from './common'

const http = axios.create({
  baseURL: `http://${baseURL}`,
  timeout: 3000,
  headers: { Accept: 'application/json' },
  responseType: 'json',
  responseEncoding: 'UTF-8'
})

// 请求拦截器
http.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    if (config.url == '/ping' || config.url == '/register' || config.url == '/login') {
      return config
    }
    const token = localStorage.getItem('token') || ''
    return {
      ...config,
      headers: {
        'Authorization': `Bearer ${token}`
      }
    }
  },
  (error: AxiosError) => {
    Promise.reject(error)
  }
)

// 响应拦截器
// http.interceptors.response.use(
//   (response: AxiosResponse) => {
//     const { data, config } = response
//     if (data.code === 419) {
//       localStorage.removeItem('token')
//       const router = useRouter()
//       router.replace({
//         path: '/login'
//       })
//       return Promise.reject(data)
//     }
//     if (data.code && data.code !== 200) {
//       return Promise.reject(data)
//     }
//     return data
//   },
//   (error: AxiosError) => {
//     console.error(error)
//   }
// )

export { http }
