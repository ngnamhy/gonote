import axios from 'axios'

export const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1', 
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json'
  }
})

// Request interceptor - sẽ được config từ app layer
export function setupAuthInterceptor(getToken) {
  apiClient.interceptors.request.use(config => {
    const token = getToken()
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  })
}

// Response interceptor - sẽ được config từ app layer
export function setupResponseInterceptor(onUnauthorized) {
  apiClient.interceptors.response.use(
    response => response,
    error => {
      if (error.response?.status === 401) {
        const url = error.config?.url
        if (!url?.includes('/login')) {
          onUnauthorized()
        }
      }
      return Promise.reject(error)
    }
  )
}

export default apiClient
