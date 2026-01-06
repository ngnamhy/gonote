import apiClient from '@/shared/api/client'

export const authApi = {
  /**
   * Register new user
   */
  async register(data) {
    const response = await apiClient.post('/auth/register', {
      username: data.username,
      email: data.email,
      name: data.name,
      password: data.password,
      confirmPassword: data.confirmPassword
    })
    return response.data
  },

  /**
   * Login user
   */
  async login(credentials) {
    const response = await apiClient.post('/auth/login', credentials)
    return response.data
  },

  /**
   * Logout user
   */
  async logout() {
    await apiClient.post('/auth/logout')
  },

  /**
   * Get current authenticated user
   */
  async getCurrentUser() {
    const response = await apiClient.get('/auth/me')
    return response.data
  }
}
