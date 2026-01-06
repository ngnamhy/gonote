import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    currentUser: null,
    isAuthenticated: false
  }),

  getters: {
    user: state => state.currentUser,
    userName: state => state.currentUser?.name || '',
    userEmail: state => state.currentUser?.email || '',
    userRole: state => state.currentUser?.role || 'user',
    isAdmin: state => state.currentUser?.role === 'admin'
  },

  actions: {
    setUser(userData) {
      this.currentUser = userData
      this.isAuthenticated = true
      
      if (userData) {
        localStorage.setItem('authUser', JSON.stringify(userData))
      }
    },

    updateUser(updates) {
      if (this.currentUser) {
        this.currentUser = { ...this.currentUser, ...updates }
        localStorage.setItem('authUser', JSON.stringify(this.currentUser))
      }
    },

    clearUser() {
      this.currentUser = null
      this.isAuthenticated = false
      localStorage.removeItem('authUser')
    },

    loadUserFromStorage() {
      const userData = localStorage.getItem('authUser')
      if (userData) {
        try {
          this.currentUser = JSON.parse(userData)
          this.isAuthenticated = true
        } catch (e) {
          console.error('Failed to parse user data:', e)
          this.clearUser()
        }
      }
    }
  }
})
