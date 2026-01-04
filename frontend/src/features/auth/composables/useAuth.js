import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { authService } from '../api/authService'
import { useAuthStore } from '../stores/authStore'

export function useAuth() {
  const router = useRouter()
  const authStore = useAuthStore()
  const loading = ref(false)
  const errors = ref({})

  const register = async (formData) => {
    loading.value = true
    errors.value = {}

    try {
      const data = await authService.register(formData)
      
      if (data.access_token && data.user) {
        authStore.setAccessToken(data.access_token)
        authStore.setUser(data.user)
        router.push('/')
      } else {
        router.push('/login')
      }
      
      return { success: true }
    } catch (error) {
      errors.value = error.response?.data?.errors || {
        general: error.response?.data?.message || 'Registration failed'
      }
      return { success: false, errors: errors.value }
    } finally {
      loading.value = false
    }
  }

  const login = async (credentials) => {
    loading.value = true
    errors.value = {}

    try {
      const data = await authService.login(credentials)
      authStore.setAccessToken(data.access_token)
      authStore.setUser(data.user)
      
      // Redirect based on role
      if (data.user?.role === 'admin') {
        router.push('/admin')
      } else {
        router.push('/feed')
      }
      
      return { success: true }
    } catch (error) {
      errors.value = error.response?.data?.errors || {
        general: error.response?.data?.message || 'Login failed'
      }
      return { success: false, errors: errors.value }
    } finally {
      loading.value = false
    }
  }

  const logout = async () => {
    try {
      await authService.logout()
    } catch (error) {
      console.error('Logout failed:', error)
    } finally {
      authStore.clearAuth()
      router.push('/login')
    }
  }

  return {
    loading,
    errors,
    register,
    login,
    logout
  }
}
