// src/utils/api.js  (hoặc đường dẫn bạn đang đặt)

import axios from 'axios'
import { useAuthStore } from '@/stores/auth'

const api = axios.create({
  baseURL: 'http://localhost:8080/api/v1', 
})

api.interceptors.request.use(config => {
  const authStore = useAuthStore()
  const access_token = authStore.access_token

  if (access_token) {
    config.headers.Authorization = `Bearer ${access_token}`
  }
  return config
})

api.interceptors.response.use(
  response => response,
  error => {
    if (error.response?.status === 401) {
      const authStore = useAuthStore()
      authStore.logout() // xóa token, user
      window.location.href = '/login' // hoặc dùng router.push nếu inject được
    }
    return Promise.reject(error)
  }
)

export default api