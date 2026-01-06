<template>
  <form @submit.prevent="handleSubmit" class="space-y-5">
    <!-- Error message -->
    <div v-if="error" class="p-3 bg-red-500/10 border border-red-500 rounded-lg">
      <p class="text-red-500 text-sm">{{ error }}</p>
    </div>

    <!-- Title -->
    <div>
      <label for="title" class="block font-semibold mb-2 text-sm" style="color: var(--color-accent)">
        Title
      </label>
      <InputText 
        id="title"
        v-model="form.title" 
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
        v-model="form.content" 
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
        @click="$emit('cancel')"
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
</template>

<script setup>
import { ref } from 'vue'
import InputText from 'primevue/inputtext'
import Textarea from 'primevue/textarea'
import Button from 'primevue/button'
import { createPost } from '../model/createPost'

const emit = defineEmits(['success', 'cancel'])

const form = ref({
  title: '',
  content: ''
})

const submitting = ref(false)
const error = ref(null)

async function handleSubmit() {
  error.value = null

  if (!form.value.title.trim() || !form.value.content.trim()) {
    error.value = 'Please fill in both title and content'
    return
  }

  submitting.value = true

  const result = await createPost(form.value)

  submitting.value = false

  if (result.success) {
    emit('success', result.data)
  } else {
    error.value = result.error
  }
}
</script>
