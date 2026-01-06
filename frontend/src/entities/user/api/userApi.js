import apiClient from '@/shared/api/client'

export const userApi = {
  /**
   * Get user profile
   */
  getProfile() {
    return apiClient.get('/users/profile')
  },

  /**
   * Update user profile
   */
  updateProfile(data) {
    return apiClient.put('/users/profile', data)
  },

  /**
   * Get user by ID
   */
  getById(userId) {
    return apiClient.get(`/users/${userId}`)
  }
}
