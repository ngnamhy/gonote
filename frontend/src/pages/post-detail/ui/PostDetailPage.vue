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
        v-if="postStore.loading"
        class="rounded-xl p-8 text-center"
        style="background-color: var(--color-bg-secondary)"
      >
        <i class="pi pi-spinner pi-spin text-3xl" style="color: var(--color-accent)"></i>
        <p class="mt-4" style="color: var(--color-text-muted)">Loading post...</p>
      </div>

      <!-- Error State -->
      <div 
        v-else-if="postStore.error"
        class="rounded-xl p-8 text-center"
        style="background-color: var(--color-bg-secondary)"
      >
        <i class="pi pi-exclamation-circle text-3xl" style="color: #ef4444"></i>
        <p class="mt-4 text-white font-semibold">{{ postStore.error }}</p>
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
        <div class="p-6">
          <div class="flex items-start gap-3 mb-4">
            <UserAvatar 
              :name="post.author.name"
              size="lg"
            />
            <div class="flex-1 min-w-0">
              <div class="flex items-center gap-2 flex-wrap">
                <span class="font-semibold text-white text-lg">{{ post.author.name }}</span>
                <span class="text-sm" style="color: var(--color-text-muted)">@{{ post.author.username }}</span>
              </div>
              <span class="text-sm" style="color: var(--color-text-muted)">{{ formattedTime }}</span>
            </div>
          </div>

          <h1 v-if="post.title" class="text-3xl font-bold text-white mb-4">{{ post.title }}</h1>

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
            class="flex items-center gap-2 px-4 py-2 rounded-lg transition-all"
            :style="upvoteStyle"
          >
            <i :class="upvoteIcon"></i>
            <span class="font-semibold">{{ post.upvotes }} Upvotes</span>
          </button>
          <button 
            class="flex items-center gap-2 px-4 py-2 rounded-lg transition-all"
            style="color: var(--color-text-muted); background-color: var(--color-bg-tertiary)"
          >
            <i class="pi pi-comment"></i>
            <span class="font-semibold">{{ post.comments }} Comments</span>
          </button>
          <button 
            class="flex items-center gap-2 px-4 py-2 rounded-lg transition-all"
            style="color: var(--color-text-muted); background-color: var(--color-bg-tertiary)"
          >
            <i class="pi pi-share-alt"></i>
            <span class="font-semibold">Share</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { usePostStore } from '@/entities/post/model/postStore'
import UserAvatar from '@/entities/user/ui/UserAvatar.vue'
import { formatTimestamp } from '@/shared/lib/utils'

const route = useRoute()
const router = useRouter()
const postStore = usePostStore()

const post = computed(() => postStore.currentPost)
const formattedTime = computed(() => post.value ? formatTimestamp(post.value.createdAt) : '')

const upvoteIcon = computed(() => 
  post.value?.hasUpvoted ? 'pi pi-arrow-up-right' : 'pi pi-arrow-up'
)

const upvoteStyle = computed(() => 
  post.value?.hasUpvoted 
    ? 'background-color: var(--color-accent); color: white'
    : 'color: var(--color-text-muted); background-color: var(--color-bg-tertiary)'
)

function toggleUpvote() {
  if (post.value) {
    post.value.hasUpvoted = !post.value.hasUpvoted
    post.value.upvotes += post.value.hasUpvoted ? 1 : -1
  }
}

function goBack() {
  router.push('/feed')
}

async function fetchPost() {
  await postStore.fetchPostById(route.params.id)
}

onMounted(() => {
  fetchPost()
})
</script>
