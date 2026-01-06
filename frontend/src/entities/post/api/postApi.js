import apiClient from '@/shared/api/client'

export const postApi = {
  /**
   * Get all posts
   */
  getPosts(params = {}) {
    return apiClient.get('/posts', { params })
  },

  /**
   * Get post by ID
   */
  getById(postId) {
    return apiClient.get(`/posts/${postId}`)
  },

  /**
   * Create new post
   */
  create(data) {
    return apiClient.post('/posts', data)
  },

  /**
   * Update post
   */
  update(postId, data) {
    return apiClient.put(`/posts/${postId}`, data)
  },

  /**
   * Delete post
   */
  delete(postId) {
    return apiClient.delete(`/posts/${postId}`)
  },

  /**
   * Upvote post
   */
  upvote(postId) {
    return apiClient.post(`/posts/${postId}/upvote`)
  },

  /**
   * Remove upvote
   */
  removeUpvote(postId) {
    return apiClient.delete(`/posts/${postId}/upvote`)
  }
}
