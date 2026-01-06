<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import api from '@/shared/services/api'
import Button from 'primevue/button'

const route = useRoute()
const router = useRouter()

const post = ref(null)
const loading = ref(false)
const error = ref(null)

// Helper function to get user initials
function getInitials(name) {
  return name
    .split(' ')
    .map(word => word[0])
    .join('')
    .toUpperCase()
    .slice(0, 2)
}

// Helper function to format timestamp
function formatTimestamp(timestamp) {
  const date = new Date(timestamp)
  const now = new Date()
  const diffMs = now - date
  const diffMins = Math.floor(diffMs / 60000)
  const diffHours = Math.floor(diffMs / 3600000)
  const diffDays = Math.floor(diffMs / 86400000)

  if (diffMins < 60) return `${diffMins} minutes ago`
  if (diffHours < 24) return `${diffHours} hours ago`
  return `${diffDays} days ago`
}

async function fetchPost() {
  try {
    loading.value = true
    error.value = null
    const response = await api.get(`/posts/${route.params.id}`)
    post.value = {
      ...response.data.data,
      hasUpvoted: false
    }
  } catch (err) {
    error.value = err.message || 'Failed to load post'
    console.error('Error fetching post:', err)
  } finally {
    loading.value = false
  }
}

function toggleUpvote() {
  if (post.value.hasUpvoted) {
    post.value.upvotes--
  } else {
    post.value.upvotes++
  }
  post.value.hasUpvoted = !post.value.hasUpvoted
}

function goBack() {
  router.push('/feed')
}

onMounted(() => {
  fetchPost()
})
</script>

<template>
  <div class="min-h-full p-6" style="background-color: var(--color-bg-primary)">
    <div class="max-w-4xl mx-auto">
      <!-- Back button -->
      <button 
        @click="goBack"
        class="mb-4 flex items-center gap-2 px-4 py-2 rounded-lg transition-all hover:bg-opacity-80"
        style="color: var(--color-text-muted); background-color: var(--color-bg-secondary)"
      >
        <i class="pi pi-arrow-left"></i>
        <span>Back to Feed</span>
      </button>

      <!-- Loading State -->
      <div 
        v-if="loading"
        class="rounded-xl p-8 text-center"
        style="background-color: var(--color-bg-secondary)"
      >
        <i class="pi pi-spinner pi-spin text-3xl" style="color: var(--color-accent)"></i>
        <p class="mt-4" style="color: var(--color-text-muted)">Loading post...</p>
      </div>

      <!-- Error State -->
      <div 
        v-else-if="error"
        class="rounded-xl p-8 text-center"
        style="background-color: var(--color-bg-secondary)"
      >
        <i class="pi pi-exclamation-circle text-3xl" style="color: #ef4444"></i>
        <p class="mt-4 text-white font-semibold">{{ error }}</p>
        <button 
          @click="fetchPost"
          class="mt-4 px-6 py-2 rounded-lg font-semibold transition-all hover:bg-opacity-80"
          style="background-color: var(--color-accent); color: white"
        >
          Try Again
        </button>
      </div>

      <!-- Post Content -->
      <div 
        v-else-if="post && post.author"
        class="rounded-xl overflow-hidden"
        style="background-color: var(--color-bg-secondary)"
      >
        <!-- Post Header -->
        <div class="p-6">
          <div class="flex items-start gap-3 mb-4">
            <div 
              class="w-12 h-12 rounded-full flex items-center justify-center flex-shrink-0"
              style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
            >
              <span class="text-white font-bold">{{ getInitials(post.author.name) }}</span>
            </div>
            <div class="flex-1 min-w-0">
              <div class="flex items-center gap-2 flex-wrap">
                <span class="font-semibold text-white text-lg">{{ post.author.name }}</span>
                <span class="text-sm" style="color: var(--color-text-muted)">@{{ post.author.username }}</span>
              </div>
              <span class="text-sm" style="color: var(--color-text-muted)">{{ formatTimestamp(post.createdAt) }}</span>
            </div>
          </div>

          <!-- Title -->
          <h1 class="text-3xl font-bold text-white mb-4">{{ post.title }}</h1>

          <!-- Content -->
          <div class="prose max-w-none">
            <p class="text-lg leading-relaxed whitespace-pre-wrap" style="color: var(--color-text-secondary)">{{ post.content }}</p>
          </div>
        </div>

        <!-- Post Actions -->
        <div 
          class="flex items-center gap-4 px-6 py-4 border-t"
          style="background-color: var(--color-bg-secondary); border-color: var(--color-border-default)"
        >
          <button 
            @click="toggleUpvote"
            class="flex items-center gap-2 px-4 py-2 rounded-lg transition-all hover:bg-opacity-80"
            :style="post.hasUpvoted ? 'background-color: var(--color-accent-muted); color: var(--color-accent)' : 'color: var(--color-text-muted)'"
          >
            <i :class="post.hasUpvoted ? 'pi pi-arrow-up-right' : 'pi pi-arrow-up'"></i>
            <span class="font-semibold">{{ post.upvotes }}</span>
            <span class="text-sm">Upvotes</span>
          </button>
          <button 
            class="flex items-center gap-2 px-4 py-2 rounded-lg transition-all hover:bg-opacity-80"
            style="color: var(--color-text-muted)"
          >
            <i class="pi pi-comment"></i>
            <span class="font-semibold">{{ post.comments }}</span>
            <span class="text-sm">Comments</span>
          </button>
          <button 
            class="flex items-center gap-2 px-4 py-2 rounded-lg transition-all hover:bg-opacity-80"
            style="color: var(--color-text-muted)"
          >
            <i class="pi pi-share-alt"></i>
            <span class="text-sm">Share</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
