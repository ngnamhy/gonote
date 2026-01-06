import { defineStore } from 'pinia'
import { postApi } from '../api/postApi'

export const usePostStore = defineStore('post', {
  state: () => ({
    posts: [],
    currentPost: null,
    loading: false,
    error: null
  }),

  getters: {
    allPosts: state => state.posts,
    getPostById: state => (id) => state.posts.find(p => p.postID === id)
  },

  actions: {
    async fetchPosts() {
      this.loading = true
      this.error = null
      
      try {
        const response = await postApi.getPosts()
        this.posts = response.data.map(post => ({
          ...post,
          hasUpvoted: false
        }))
      } catch (err) {
        this.error = err.message || 'Failed to fetch posts'
        throw err
      } finally {
        this.loading = false
      }
    },

    async fetchPostById(postId) {
      this.loading = true
      this.error = null
      
      try {
        const response = await postApi.getById(postId)
        this.currentPost = {
          ...response.data.data,
          hasUpvoted: false
        }
      } catch (err) {
        this.error = err.message || 'Failed to fetch post'
        throw err
      } finally {
        this.loading = false
      }
    },

    async createPost(postData) {
      try {
        const response = await postApi.create(postData)
        return response.data
      } catch (err) {
        throw err
      }
    },

    toggleUpvote(postId) {
      const post = this.posts.find(p => p.postID === postId)
      if (post) {
        if (post.hasUpvoted) {
          post.upvotes--
        } else {
          post.upvotes++
        }
        post.hasUpvoted = !post.hasUpvoted
      }
    }
  }
})
