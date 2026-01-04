import api from '@/shared/services/api'

export const authService = {
  async register(data) {
    const response = await api.post('/auth/register', {
      username: data.username,
      email: data.email,
      password: data.password,
      confirm_password: data.confirm_password
    })
    return response.data
  },

  async login(credentials) {
    const response = await api.post('/auth/login', credentials)
    return response.data
  },

  async logout() {
    await api.post('/auth/logout')
  },

  async getCurrentUser() {
    const response = await api.get('/auth/me')
    return response.data
  }
}
