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
        <CreatePostForm 
          @success="handleSuccess"
          @cancel="cancel"
        />
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

<script setup>
import { useRouter } from 'vue-router'
import CreatePostForm from '@/features/create-post/ui/CreatePostForm.vue'

const router = useRouter()

function handleSuccess() {
  router.push('/feed')
}

function cancel() {
  router.push('/feed')
}
</script>
