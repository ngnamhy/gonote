import { postApi } from '@/entities/post/api/postApi'
import { usePostStore } from '@/entities/post/model/postStore'

export async function createPost(postData) {
  const postStore = usePostStore()
  
  try {
    const result = await postStore.createPost(postData)
    return { success: true, data: result }
  } catch (err) {
    return { 
      success: false, 
      error: err.response?.data?.message || 'Failed to create post' 
    }
  }
}
