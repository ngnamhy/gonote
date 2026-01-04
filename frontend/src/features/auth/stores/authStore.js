import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    access_token: null,
  }),

  getters: {
    isLoggedIn: state => !!state.access_token,
    isAuthenticated: state => !!state.access_token,
    isAdmin: state => state.user?.role === 'admin', 
  },

  actions: {
    setAccessToken(access_token) {
      this.access_token = access_token
      if (access_token) {
        localStorage.setItem('access_token', access_token)
      } else {
        localStorage.removeItem('access_token')
      }
    },

    setUser(user) {
      this.user = user
      if (user) {
        localStorage.setItem('auth_user', JSON.stringify(user))
      } else {
        localStorage.removeItem('auth_user')
      }
    },

    clearAuth() {
      this.access_token = null
      this.user = null
      localStorage.removeItem('access_token')
      localStorage.removeItem('auth_user')
    },

    logout() {
      this.clearAuth()
    },

    loadFromStorage() {
      const token = localStorage.getItem('access_token')
      const user = localStorage.getItem('auth_user')

      if (token) {
        this.access_token = token
      }
      
      if (user) {
        try {
          this.user = JSON.parse(user)
        } catch (e) {
          console.error('Failed to parse user data:', e)
        }
      }
    }
  },
})
