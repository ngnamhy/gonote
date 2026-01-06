import { setupAuthInterceptor, setupResponseInterceptor } from '@/shared/api/client'
import { useAuthStore } from '@/features/auth/model/authStore'
import { useUserStore } from '@/entities/user/model/userStore'

export function setupApiInterceptors() {
  const authStore = useAuthStore()
  const userStore = useUserStore()

  // Load auth data from storage
  authStore.loadFromStorage()

  // Setup request interceptor
  setupAuthInterceptor(() => authStore.token)

  // Setup response interceptor
  setupResponseInterceptor(() => {
    authStore.clearAuth()
    window.location.href = '/login'
  })
}
