<template>
  <form @submit.prevent="handleSubmit" class="space-y-5">
    <!-- Error message -->
    <div v-if="error" class="p-3 bg-red-500/10 border border-red-500 rounded-lg">
      <p class="text-red-500 text-sm">{{ error }}</p>
    </div>

    <!-- Name -->
    <div>
      <label for="name" class="block font-semibold mb-2 text-sm">Full Name</label>
      <InputText 
        id="name" 
        v-model="form.name" 
        class="w-full" 
        placeholder="John Doe"
        :disabled="loading"
      />
    </div>

    <!-- Username -->
    <div>
      <label for="username" class="block font-semibold mb-2 text-sm">Username</label>
      <InputText 
        id="username" 
        v-model="form.username" 
        class="w-full" 
        placeholder="johndoe"
        :disabled="loading"
      />
    </div>

    <!-- Email -->
    <div>
      <label for="email" class="block font-semibold mb-2 text-sm">Email</label>
      <InputText 
        id="email" 
        v-model="form.email" 
        type="email"
        class="w-full" 
        placeholder="you@example.com"
        :disabled="loading"
      />
    </div>

    <!-- Password -->
    <div>
      <label for="password" class="block font-semibold mb-2 text-sm">Password</label>
      <Password 
        id="password" 
        v-model="form.password" 
        class="w-full" 
        toggleMask 
        :feedback="false"
        placeholder="••••••••"
        :disabled="loading"
      />
    </div>

    <!-- Confirm Password -->
    <div>
      <label for="confirmPassword" class="block font-semibold mb-2 text-sm">Confirm Password</label>
      <Password 
        id="confirmPassword" 
        v-model="form.confirmPassword" 
        class="w-full" 
        toggleMask 
        :feedback="false"
        placeholder="••••••••"
        :disabled="loading"
      />
    </div>

    <Button 
      type="submit" 
      label="Create Account" 
      class="w-full py-3.5 text-lg font-semibold rounded-xl" 
      :loading="loading"
    />
  </form>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import InputText from 'primevue/inputtext'
import Password from 'primevue/password'
import Button from 'primevue/button'
import { useAuthStore } from '../model/authStore'
import { validateEmail } from '@/shared/lib/utils'

const router = useRouter()
const authStore = useAuthStore()

const form = ref({
  name: '',
  username: '',
  email: '',
  password: '',
  confirmPassword: ''
})

const loading = ref(false)
const error = ref(null)

async function handleSubmit() {
  error.value = null

  // Validation
  if (!form.value.name.trim()) {
    error.value = 'Name is required'
    return
  }

  if (!form.value.username.trim()) {
    error.value = 'Username is required'
    return
  }

  if (!validateEmail(form.value.email)) {
    error.value = 'Invalid email format'
    return
  }

  if (!form.value.password || form.value.password.length < 6) {
    error.value = 'Password must be at least 6 characters'
    return
  }

  if (form.value.password !== form.value.confirmPassword) {
    error.value = 'Passwords do not match'
    return
  }

  loading.value = true

  const success = await authStore.register(form.value)

  loading.value = false

  if (success) {
    router.push('/feed')
  } else {
    error.value = authStore.error
  }
}
</script>

<style scoped>
:deep(.p-password) {
  width: 100%;
}

:deep(.p-password .p-password-input) {
  width: 100%;
}
</style>
