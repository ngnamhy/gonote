<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/shared/services/api'
import InputText from 'primevue/inputtext'
import Textarea from 'primevue/textarea'
import Button from 'primevue/button'

const router = useRouter()

const newPost = ref({
  title: '',
  content: ''
})

const submitting = ref(false)
const error = ref(null)

async function submitPost() {
  if (!newPost.value.title.trim() || !newPost.value.content.trim()) {
    error.value = 'Please fill in both title and content'
    return
  }

  try {
    submitting.value = true
    error.value = null
    
    await api.post('/posts', {
      title: newPost.value.title,
      content: newPost.value.content
    })
    
    // Redirect to feed after successful post
    router.push('/feed')
  } catch (err) {
    error.value = err.response?.data?.message || 'Failed to create post'
    console.error('Error creating post:', err)
  } finally {
    submitting.value = false
  }
}

function cancel() {
  router.push('/feed')
}
</script>

<template>
  <div class="min-h-full p-6" style="background-color: var(--color-bg-primary)">
    <div class="max-w-4xl mx-auto">
      <!-- Header -->
      <div class="mb-6">
        <div class="flex items-center gap-3 mb-2">
          <button 
            @click="cancel"
            class="p-2 rounded-lg transition-all hover:bg-opacity-80"
            style="color: var(--color-text-muted)"
          >
            <i class="pi pi-arrow-left"></i>
          </button>
          <h1 class="text-3xl font-bold text-white">Create New Post</h1>
        </div>
        <p style="color: var(--color-text-muted)">Share your thoughts with the community</p>
      </div>

      <!-- Form -->
      <div 
        class="rounded-xl p-6"
        style="background-color: var(--color-bg-secondary)"
      >
        <!-- Error message -->
        <div v-if="error" class="mb-4 p-3 bg-red-500/10 border border-red-500 rounded-lg">
          <p class="text-red-500 text-sm">{{ error }}</p>
        </div>

        <form @submit.prevent="submitPost" class="space-y-5">
          <!-- Title -->
          <div>
            <label for="title" class="block font-semibold mb-2 text-sm" style="color: var(--color-accent)">
              Title
            </label>
            <InputText 
              id="title"
              v-model="newPost.title" 
              class="w-full" 
              placeholder="Enter your post title"
              :disabled="submitting"
              size="large"
            />
          </div>

          <!-- Content -->
          <div>
            <label for="content" class="block font-semibold mb-2 text-sm" style="color: var(--color-accent)">
              Content
            </label>
            <Textarea 
              id="content"
              v-model="newPost.content" 
              class="w-full" 
              rows="12"
              placeholder="What's on your mind? Share your thoughts, ideas, or updates..."
              :disabled="submitting"
            />
          </div>

          <!-- Actions -->
          <div class="flex gap-3 justify-end pt-4">
            <Button 
              type="button"
              label="Cancel" 
              severity="secondary"
              @click="cancel"
              :disabled="submitting"
              size="large"
            />
            <Button 
              type="submit"
              label="Publish Post" 
              :loading="submitting"
              size="large"
              icon="pi pi-send"
            />
          </div>
        </form>
      </div>

      <!-- Tips -->
      <div 
        class="mt-4 rounded-xl p-4"
        style="background-color: var(--color-bg-secondary)"
      >
        <h3 class="font-semibold text-white mb-2 flex items-center gap-2">
          <i class="pi pi-info-circle" style="color: var(--color-accent)"></i>
          Tips for a great post
        </h3>
        <ul class="space-y-1 text-sm" style="color: var(--color-text-muted)">
          <li>• Keep your title clear and concise</li>
          <li>• Provide context and details in the content</li>
          <li>• Be respectful and constructive</li>
          <li>• Use proper formatting for better readability</li>
        </ul>
      </div>
    </div>
  </div>
</template>
