<template>
  <div 
    :class="['user-avatar', sizeClass]"
    :style="avatarStyle"
  >
    <span v-if="!image" class="initials">{{ initials }}</span>
    <img v-else :src="image" :alt="name" />
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { getUserInitials } from '@/shared/lib/utils'

const props = defineProps({
  name: {
    type: String,
    default: ''
  },
  image: {
    type: String,
    default: ''
  },
  size: {
    type: String,
    default: 'md',
    validator: (val) => ['sm', 'md', 'lg', 'xl'].includes(val)
  }
})

const initials = computed(() => getUserInitials(props.name))

const sizeClass = computed(() => `size-${props.size}`)

const avatarStyle = computed(() => {
  if (props.image) return {}
  return {
    background: 'linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))'
  }
})
</script>

<style scoped>
.user-avatar {
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 9999px;
  overflow: hidden;
  flex-shrink: 0;
}

.user-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.initials {
  color: white;
  font-weight: 700;
}

.size-sm {
  width: 2rem;
  height: 2rem;
}

.size-sm .initials {
  font-size: 0.75rem;
}

.size-md {
  width: 2.5rem;
  height: 2.5rem;
}

.size-md .initials {
  font-size: 0.875rem;
}

.size-lg {
  width: 3rem;
  height: 3rem;
}

.size-lg .initials {
  font-size: 1rem;
}

.size-xl {
  width: 6rem;
  height: 6rem;
}

.size-xl .initials {
  font-size: 1.5rem;
}
</style>
