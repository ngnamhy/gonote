import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    accessToken: null,
  }),

  getters: {
    isLoggedIn: state => !!state.accessToken,
    isAuthenticated: state => !!state.accessToken,
    isAdmin: state => state.user?.role === 'admin', 
  },

  actions: {
    setAccessToken(accessToken) {
      this.accessToken = accessToken
      if (accessToken) {
        localStorage.setItem('accessToken', accessToken)
      } else {
        localStorage.removeItem('accessToken')
      }
    },

    setUser(user) {
      this.user = user
      console.log(JSON.stringify(user))
      if (user) {
        localStorage.setItem('authUser', JSON.stringify(user))
      } else {
        localStorage.removeItem('authUser')
      }
    },

    updateUser(updatedData) {
      if (this.user) {
        this.user = { ...this.user, ...updatedData }
        localStorage.setItem('authUser', JSON.stringify(this.user))
      }
    },

    clearAuth() {
      this.accessToken = null
      this.user = null
      localStorage.removeItem('accessToken')
      localStorage.removeItem('authUser')
    },

    logout() {
      this.clearAuth()
    },

    loadFromStorage() {
      const token = localStorage.getItem('accessToken')
      const user = localStorage.getItem('authUser')

      if (token) {
        this.accessToken = token
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
