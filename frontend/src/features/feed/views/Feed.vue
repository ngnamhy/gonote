<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/features/auth/stores/authStore'
import api from '@/shared/services/api'

const router = useRouter()
const authStore = useAuthStore()

// Posts data from API
const posts = ref([])
const loading = ref(false)
const error = ref(null)

// Trending topics
const trending = ref([
  { tag: 'product-launch', posts: 12 },
  { tag: 'team-building', posts: 8 },
  { tag: 'engineering', posts: 15 },
  { tag: 'design-system', posts: 6 },
])

// Fetch posts from API
async function fetchPosts() {
  try {
    loading.value = true
    error.value = null
    const response = await api.get('/posts')
    posts.value = response.data.map(post => ({
      ...post,
      hasUpvoted: false
    }))
  } catch (err) {
    error.value = err.message || 'Failed to load posts'
    console.error('Error fetching posts:', err)
  } finally {
    loading.value = false
  }
}

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

// Helper function to truncate content
function truncateContent(content, maxLength = 200) {
  if (content.length <= maxLength) return content
  return content.substring(0, maxLength) + '...'
}

function toggleUpvote(post) {
  if (post.hasUpvoted) {
    post.upvotes--
  } else {
    post.upvotes++
  }
  post.hasUpvoted = !post.hasUpvoted
}

function createPost() {
  router.push('/submit')
}

function viewPost(postId) {
  router.push(`/posts/${postId}`)
}

// Load posts on mount
onMounted(() => {
  fetchPosts()
})
</script>

<template>
  <div class="flex justify-between gap-6 p-6" style="background-color: var(--color-bg-primary); min-height: 100%;">
    <!-- Main Feed -->
     <div></div>
    <main class="flex-1 max-w-4xl space-y-4">
      <!-- Create Post Card -->
      <div 
        class="rounded-xl p-4"
        style="background-color: var(--color-bg-secondary)"
      >
        <div class="flex items-center gap-3">
          <div 
            class="w-10 h-10 rounded-full flex items-center justify-center"
            style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
          >
            <i class="pi pi-user text-white"></i>
          </div>
          <button 
            @click="createPost"
            class="flex-1 px-4 py-3 rounded-lg text-left transition-all hover:bg-opacity-80"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-muted)"
          >
            Share something with your team...
          </button>
        </div>
        <div class="flex gap-2 mt-3 pt-3 border-t" style="border-color: var(--color-border-default)">
          <button 
            class="flex-1 py-2 px-3 rounded-lg text-sm font-semibold transition-all hover:bg-opacity-80 flex items-center justify-center gap-2"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            <i class="pi pi-image"></i>
            Image
          </button>
          <button 
            class="flex-1 py-2 px-3 rounded-lg text-sm font-semibold transition-all hover:bg-opacity-80 flex items-center justify-center gap-2"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            <i class="pi pi-link"></i>
            Link
          </button>
          <button 
            class="flex-1 py-2 px-3 rounded-lg text-sm font-semibold transition-all hover:bg-opacity-80 flex items-center justify-center gap-2"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            <i class="pi pi-calendar"></i>
            Event
          </button>
        </div>
      </div>

      <!-- Loading State -->
      <div 
        v-if="loading"
        class="rounded-xl p-8 text-center"
        style="background-color: var(--color-bg-secondary)"
      >
        <i class="pi pi-spinner pi-spin text-3xl" style="color: var(--color-accent)"></i>
        <p class="mt-4" style="color: var(--color-text-muted)">Loading posts...</p>
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
          @click="fetchPosts"
          class="mt-4 px-6 py-2 rounded-lg font-semibold transition-all hover:bg-opacity-80"
          style="background-color: var(--color-accent); color: white"
        >
          Try Again
        </button>
      </div>

      <!-- Posts Feed -->
      <div 
        v-for="post in posts" 
        :key="post.postID"
        class="rounded-xl overflow-hidden cursor-pointer transition-all hover:ring-2"
        style="background-color: var(--color-bg-secondary); --tw-ring-color: var(--color-accent)"
        @click="viewPost(post.postID)"
      >
        <!-- Post Header -->
        <div class="p-4">
          <div class="flex items-center gap-3">
            <div 
              class="w-10 h-10 rounded-full flex items-center justify-center flex-shrink-0"
              style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
            >
              <span class="text-white text-sm font-bold">{{ getInitials(post.author.name) }}</span>
            </div>
            <div class="flex-1 min-w-0 flex items-center">
              <div class="flex items-center gap-2 flex-wrap">
                <span class="font-semibold text-white">{{ post.author.name }}</span>
                <span class="text-xs" style="color: var(--color-text-muted)">@{{ post.author.username }}</span>
                <span class="text-xs" style="color: var(--color-text-muted)">•</span>
                <span class="text-xs" style="color: var(--color-text-muted)">{{ formatTimestamp(post.createdAt) }}</span>
              </div>
            </div>
          </div>

          <!-- Post Content -->
          <div class="mt-3 mb-3">
            <h1 v-if="post.title" class="font-bold text-lg text-white mb-2">{{ post.title }}</h1>
            <p class="line-clamp-3" style="color: var(--color-text-secondary)">{{ truncateContent(post.content) }}</p>
          </div>

          <!-- Post Actions -->
          <div 
            class="flex items-center gap-2 p-2 rounded-lg w-fit"
            style="background-color: var(--color-bg-tertiary)"
            @click.stop
          >
          <button 
            @click="toggleUpvote(post)"
            class="flex items-center gap-1.5 px-2.5 py-1.5 rounded-md transition-all"
            :style="post.hasUpvoted ? 'background-color: var(--color-accent-muted); color: var(--color-accent)' : 'color: var(--color-text-muted); background-color: transparent'"
            style="hover:background-color: var(--color-bg-primary)"
          >
            <i :class="post.hasUpvoted ? 'pi pi-arrow-up-right' : 'pi pi-arrow-up'" class="text-sm"></i>
            <span class="font-semibold text-sm">{{ post.upvotes }}</span>
          </button>
          <button 
            class="flex items-center gap-1.5 px-2.5 py-1.5 rounded-md transition-all hover:bg-opacity-80"
            style="color: var(--color-text-muted)"
          >
            <i class="pi pi-comment text-sm"></i>
            <span class="font-semibold text-sm">{{ post.comments }}</span>
          </button>
          <button 
            class="flex items-center gap-1.5 px-2.5 py-1.5 rounded-md transition-all hover:bg-opacity-80"
            style="color: var(--color-text-muted)"
          >
            <i class="pi pi-share-alt text-sm"></i>
            <span class="font-semibold text-sm">Share</span>
          </button>
          </div>
        </div>
      </div>
    </main>

    <!-- Right Sidebar - Trending & Info -->
    <aside class="hidden xl:block w-80">
      <div class="sticky top-6 space-y-4">
        <!-- Trending Topics -->
        <div 
          class="rounded-xl p-4"
          style="background-color: var(--color-bg-secondary)"
        >
        <h3 class="font-bold text-white mb-4 flex items-center gap-2">
          <i class="pi pi-bolt" style="color: var(--color-accent)"></i>
          Trending Topics
        </h3>
        <div class="space-y-3">
          <div 
            v-for="topic in trending" 
            :key="topic.tag"
            class="flex items-center justify-between p-3 rounded-lg cursor-pointer transition-all hover:bg-opacity-80"
            style="background-color: var(--color-bg-tertiary)"
          >
            <div>
              <p class="font-semibold text-sm" style="color: var(--color-accent)">#{{ topic.tag }}</p>
              <p class="text-xs" style="color: var(--color-text-muted)">{{ topic.posts }} posts</p>
            </div>
            <i class="pi pi-arrow-right" style="color: var(--color-text-muted)"></i>
          </div>
        </div>
      </div>

      <!-- Quick Actions -->
      <div 
        class="rounded-xl p-4"
        style="background-color: var(--color-bg-secondary)"
      >
        <h3 class="font-bold text-white mb-4">Quick Actions</h3>
        <div class="space-y-2">
          <button 
            class="w-full px-4 py-3 rounded-lg font-semibold text-sm transition-all hover:bg-opacity-80 flex items-center gap-3"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            <i class="pi pi-users" style="color: var(--color-accent)"></i>
            Create Team
          </button>
          <button 
            class="w-full px-4 py-3 rounded-lg font-semibold text-sm transition-all hover:bg-opacity-80 flex items-center gap-3"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            <i class="pi pi-calendar" style="color: var(--color-accent)"></i>
            Schedule Event
          </button>
          <button 
            class="w-full px-4 py-3 rounded-lg font-semibold text-sm transition-all hover:bg-opacity-80 flex items-center gap-3"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            <i class="pi pi-megaphone" style="color: var(--color-accent)"></i>
            Make Announcement
          </button>
        </div>
      </div>

      <!-- Team Stats -->
      <div 
        class="rounded-xl p-4"
        style="background-color: var(--color-bg-secondary)"
      >
        <h3 class="font-bold text-white mb-4">Your Activity</h3>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <span class="text-sm" style="color: var(--color-text-muted)">Posts</span>
            <span class="font-bold text-white">24</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-sm" style="color: var(--color-text-muted)">Comments</span>
            <span class="font-bold text-white">156</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-sm" style="color: var(--color-text-muted)">Upvotes Received</span>
            <span class="font-bold text-white">342</span>
          </div>
        </div>
      </div>
      </div>
    </aside>
  </div>
</template>
