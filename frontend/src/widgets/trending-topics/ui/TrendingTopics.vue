<template>
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
          <p class="font-semibold text-white">#{{ topic.tag }}</p>
          <p class="text-xs" style="color: var(--color-text-muted)">{{ topic.posts }} posts</p>
        </div>
        <i class="pi pi-arrow-right text-sm" style="color: var(--color-text-muted)"></i>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const trending = ref([
  { tag: 'product-launch', posts: 12 },
  { tag: 'team-building', posts: 8 },
  { tag: 'engineering', posts: 15 },
  { tag: 'design-system', posts: 6 },
])
</script>
