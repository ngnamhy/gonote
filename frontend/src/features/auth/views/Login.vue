<script setup>
import { ref } from 'vue'
import InputText from 'primevue/inputtext'
import Password from 'primevue/password'
import Button from 'primevue/button'
import { useAuth } from '../composables/useAuth'
import { useValidation } from '@/shared/composables/useValidation'

const { login, loading, errors } = useAuth()
const { validateLoginForm } = useValidation()

const form = ref({
  email: '',
  password: ''
})

async function handleSubmit() {
  const validationErrors = validateLoginForm(form.value)
  if (Object.keys(validationErrors).length > 0) {
    errors.value = validationErrors
    return
  }

  // Call login from composable
  await login({
    email: form.value.email,
    password: form.value.password
  })
}

function socialLogin(provider) {
  console.log('Login with', provider)
  // Xử lý OAuth ở đây
}
</script>

<style scoped>
.gradient-panel {
  background: linear-gradient(135deg, #141204 0%, #262A10 40%, #54442B 100%);
  box-shadow: 0 20px 60px rgba(20, 18, 4, 0.6);
}
</style>

<template>
  <div class="min-h-screen flex p-8 gap-8" style="background-color: #141204">
    <!-- Phần trái: Gradient background -->
    <div class="hidden lg:flex lg:w-1/2 flex-col justify-between gradient-panel text-white p-12 rounded-[2rem]">
      <router-link to="/" class="flex items-center gap-3 hover:opacity-80 transition-opacity w-fit">
        <i class="pi pi-sparkles text-4xl"></i>
        <span class="text-2xl font-bold">gono.</span>
      </router-link>

      <div class="max-w-md">
        <p class="text-sm mb-2 opacity-90">You can easily</p>
        <h1 class="text-4xl font-semibold leading-tight">
          Get access your personal hub for clarity and productivity
        </h1>
      </div>

      <div></div>
    </div>

    <!-- Phần phải: Form -->
    <div class="flex-1 flex items-center justify-center p-4">
      <div class="w-full max-w-md p-10 rounded-3xl">
        <div class="text-center mb-8">
          <i class="pi pi-sparkles text-5xl mb-4" style="color: #E8985E"></i>
          <h2 class="text-3xl font-bold text-white mb-2">Welcome Back</h2>
          <p class="text-sm">
            Access your tasks, notes, and projects anytime, anywhere - and keep everything flowing in one place.
          </p>
        </div>

        <!-- General error message -->
        <div v-if="errors.general" class="mb-4 p-3 bg-red-500/10 border border-red-500 rounded-lg">
          <p class="text-red-500 text-sm">{{ errors.general }}</p>
        </div>

        <form @submit.prevent="handleSubmit">
          <!-- Email -->
          <div class="mb-5">
            <label for="email" class="block font-semibold mb-2 text-sm">Email</label>
            <InputText 
              id="email" 
              v-model="form.email" 
              type="email"
              class="w-full" 
              placeholder="you@example.com"
              :invalid="!!errors.email" 
            />
            <small v-if="errors.email" class="text-red-500 text-sm mt-2">{{ errors.email }}</small>
          </div>

          <!-- Password -->
          <div class="mb-5">
            <label for="password" class="block font-semibold mb-2 text-sm">Password</label>
            <Password 
              id="password" 
              v-model="form.password" 
              class="w-full" 
              toggleMask 
              :feedback="false"
              placeholder="••••••••" 
              :invalid="!!errors.password" 
            />
            <small v-if="errors.password" class="text-red-500 text-sm mt-2">{{ errors.password }}</small>
          </div>

          <Button 
            type="submit" 
            label="Get Started" 
            class="w-full py-3.5 text-lg font-semibold rounded-xl mt-2" 
            style="color: var(--color-text-primary); 
              background-color: var(--color-bg-tertiary); 
              border: 0px;" 
            onmouseenter="this.style.backgroundColor='#54442B'"
            onmouseleave="this.style.backgroundColor='#262A10'" 
            :loading="loading" 
          />

          <!-- Divider -->
          <div class="flex items-center my-7 gap-4">
            <div class="flex-1 h-px" style="background-color: #54442B"></div>
            <span class="text-sm" style="color: #A9714B">or continue with</span>
            <div class="flex-1 h-px" style="background-color: #54442B"></div>
          </div>

          <!-- Social buttons -->
          <div class="grid grid-cols-1 gap-3 mb-6">
            <button 
              type="button" 
              class="p-3 rounded-xl transition-colors flex items-center justify-center"
              style="border: 0px; background-color: #262A10"
              onmouseenter="this.style.backgroundColor='#54442B'" 
              onmouseleave="this.style.backgroundColor='#262A10'"
              @click="socialLogin('google')">
              <img src="https://www.google.com/favicon.ico" alt="Google" class="h-6 w-auto" />
            </button>
          </div>

          <!-- Sign up link -->
          <p class="text-center text-sm" style="color: #A9714B">
            Don't have an account?
            <router-link to="/register" class="font-semibold hover:underline" style="color: #E8985E">
              Sign up
            </router-link>
          </p>
        </form>
      </div>
    </div>
  </div>
</template>
