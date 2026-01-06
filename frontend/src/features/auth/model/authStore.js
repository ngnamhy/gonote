import { defineStore } from 'pinia'
import { authApi } from '../api/authApi'
import { useUserStore } from '@/entities/user/model/userStore'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    accessToken: null,
    loading: false,
    error: null
  }),

  getters: {
    isAuthenticated: state => !!state.accessToken,
    token: state => state.accessToken
  },

  actions: {
    async login(credentials) {
      this.loading = true
      this.error = null

      try {
        const response = await authApi.login(credentials)
        const { accessToken, user } = response

        // Save token
        this.setAccessToken(accessToken)

        // Update user entity
        const userStore = useUserStore()
        userStore.setUser(user)

        return true
      } catch (err) {
        this.error = err.response?.data?.message || 'Login failed'
        return false
      } finally {
        this.loading = false
      }
    },

    async register(userData) {
      this.loading = true
      this.error = null

      try {
        const response = await authApi.register(userData)
        const { accessToken, user } = response

        // Save token
        this.setAccessToken(accessToken)

        // Update user entity
        const userStore = useUserStore()
        userStore.setUser(user)

        return true
      } catch (err) {
        this.error = err.response?.data?.message || 'Registration failed'
        return false
      } finally {
        this.loading = false
      }
    },

    async logout() {
      try {
        await authApi.logout()
      } catch (err) {
        console.error('Logout error:', err)
      } finally {
        this.clearAuth()
      }
    },

    setAccessToken(token) {
      this.accessToken = token
      if (token) {
        localStorage.setItem('accessToken', token)
      } else {
        localStorage.removeItem('accessToken')
      }
    },

    clearAuth() {
      this.accessToken = null
      localStorage.removeItem('accessToken')

      const userStore = useUserStore()
      userStore.clearUser()
    },

    loadFromStorage() {
      const token = localStorage.getItem('accessToken')
      if (token) {
        this.accessToken = token
      }

      const userStore = useUserStore()
      userStore.loadUserFromStorage()
    }
  }
})
