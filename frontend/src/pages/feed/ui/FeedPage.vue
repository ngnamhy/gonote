<template>
  <div class="flex justify-between gap-6 p-6" style="background-color: var(--color-bg-primary); min-height: 100%;">
    <div></div>
    
    <main class="flex-1 max-w-4xl space-y-4">
      <!-- Create Post Card -->
      <div 
        class="rounded-xl p-4"
        style="background-color: var(--color-bg-secondary)"
      >
        <div class="flex items-center gap-3">
          <UserAvatar :name="userName" size="md" />
          <button 
            @click="createPost"
            class="flex-1 px-4 py-3 rounded-lg text-left transition-all hover:bg-opacity-80"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-muted)"
          >
            Share something with your team...
          </button>
        </div>
      </div>

      <!-- Loading State -->
      <div 
        v-if="postStore.loading"
        class="rounded-xl p-8 text-center"
        style="background-color: var(--color-bg-secondary)"
      >
        <i class="pi pi-spinner pi-spin text-3xl" style="color: var(--color-accent)"></i>
        <p class="mt-4" style="color: var(--color-text-muted)">Loading posts...</p>
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
          @click="fetchPosts"
          class="mt-4 px-6 py-2 rounded-lg font-semibold transition-all hover:bg-opacity-80"
          style="background-color: var(--color-accent); color: white"
        >
          Try Again
        </button>
      </div>

      <!-- Posts Feed -->
      <PostCard 
        v-for="post in postStore.allPosts" 
        :key="post.postID"
        :post="post"
        @click="viewPost"
        @upvote="handleUpvote"
      />
    </main>

    <!-- Right Sidebar -->
    <aside class="hidden xl:block w-80">
      <div class="sticky top-6">
        <TrendingTopics />
      </div>
    </aside>
  </div>
</template>

<script setup>
import { onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { usePostStore } from '@/entities/post/model/postStore'
import { useUserStore } from '@/entities/user/model/userStore'
import UserAvatar from '@/entities/user/ui/UserAvatar.vue'
import PostCard from '@/widgets/post-card/ui/PostCard.vue'
import TrendingTopics from '@/widgets/trending-topics/ui/TrendingTopics.vue'

const router = useRouter()
const postStore = usePostStore()
const userStore = useUserStore()

const userName = computed(() => userStore.userName)

function createPost() {
  router.push('/submit')
}

function viewPost(postId) {
  router.push(`/posts/${postId}`)
}

function handleUpvote(postId) {
  postStore.toggleUpvote(postId)
}

async function fetchPosts() {
  await postStore.fetchPosts()
}

onMounted(() => {
  fetchPosts()
})
</script>
