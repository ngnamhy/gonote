import axios from 'axios'
import { useAuthStore } from '@/features/auth/stores/authStore'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1', 
  withCredentials: true, 
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
      const url = error.config?.url
      if (url && url.includes('/login')) { 
        return Promise.reject(error)
      }
      const authStore = useAuthStore()
      authStore.logout() 
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export default api
