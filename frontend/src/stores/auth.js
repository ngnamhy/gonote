import { defineStore } from 'pinia'
import api from '@/services/api'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    access_token: null,
  }),

  getters: {
    isLoggedIn: state => !!state.token,
    isAdmin: state => state.user.role === 'admin', 
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

    async login(credentials) {
      const response = await api.post('/auth/login', credentials)
      const { access_token, user } = response.data
      console.log("USER: " + user)
      this.setAccessToken(access_token)
      this.setUser(user)
    },

    async fetchUser() {
      const response = await api.get('/me') 
      this.setUser(response.data.user)
    },

    logout() {
      this.token = null
      this.user = null
      localStorage.removeItem('access_token')
    },

    loadFromStorage() {
      const token = localStorage.getItem('access_token')

      if (token) {
        this.access_token = token
      }
    }
  },
})