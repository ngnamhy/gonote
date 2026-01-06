<template>
  <div 
    class="rounded-xl overflow-hidden cursor-pointer transition-all hover:ring-2"
    style="background-color: var(--color-bg-secondary); --tw-ring-color: var(--color-accent)"
    @click="$emit('click', post.postID)"
  >
    <!-- Post Header -->
    <div class="p-4">
      <div class="flex items-center gap-3">
        <UserAvatar 
          :name="post.author.name"
          size="md"
        />
        <div class="flex-1 min-w-0 flex items-center">
          <div class="flex items-center gap-2 flex-wrap">
            <span class="font-semibold text-white">{{ post.author.name }}</span>
            <span class="text-xs" style="color: var(--color-text-muted)">@{{ post.author.username }}</span>
            <span class="text-xs" style="color: var(--color-text-muted)">•</span>
            <span class="text-xs" style="color: var(--color-text-muted)">{{ formattedTime }}</span>
          </div>
        </div>
      </div>

      <!-- Post Content -->
      <div class="mt-3 mb-3">
        <h1 v-if="post.title" class="font-bold text-lg text-white mb-2">{{ post.title }}</h1>
        <p class="line-clamp-3" style="color: var(--color-text-secondary)">{{ truncatedContent }}</p>
      </div>

      <!-- Post Actions -->
      <div 
        class="flex items-center gap-2 p-2 rounded-lg w-fit"
        style="background-color: var(--color-bg-tertiary)"
        @click.stop
      >
        <button 
          @click="handleUpvote"
          class="flex items-center gap-1.5 px-2.5 py-1.5 rounded-md transition-all"
          :style="upvoteStyle"
        >
          <i :class="upvoteIcon" class="text-sm"></i>
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
</template>

<script setup>
import { computed } from 'vue'
import UserAvatar from '@/entities/user/ui/UserAvatar.vue'
import { formatTimestamp, truncateText } from '@/shared/lib/utils'

const props = defineProps({
  post: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['click', 'upvote'])

const formattedTime = computed(() => formatTimestamp(props.post.createdAt))
const truncatedContent = computed(() => truncateText(props.post.content, 200))

const upvoteIcon = computed(() => 
  props.post.hasUpvoted ? 'pi pi-arrow-up-right' : 'pi pi-arrow-up'
)

const upvoteStyle = computed(() => 
  props.post.hasUpvoted 
    ? 'background-color: var(--color-accent-muted); color: var(--color-accent)'
    : 'color: var(--color-text-muted); background-color: transparent'
)

function handleUpvote() {
  emit('upvote', props.post.postID)
}
</script>
