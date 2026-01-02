<script setup>
import { ref } from 'vue'
import InputText from 'primevue/inputtext'
import Password from 'primevue/password'
import Button from 'primevue/button'
import { useRouter } from 'vue-router'

const router = useRouter()
const loading = ref(false)
const form = ref({
  username: '',
  email: '',
  password: '',
  confirm_password: ''
})
const errors = ref({})

function handleSubmit() {
  // Xử lý đăng ký ở đây (gọi API, validate, etc.)
  loading.value = true
  setTimeout(() => {
    loading.value = false
    router.push('/login')
  }, 2000)
}
</script>

<style scoped>
.gradient-panel {
  background: linear-gradient(135deg, #54442B 0%, #A9714B 50%, #E8985E 100%);
  box-shadow: 0 20px 60px rgba(84, 68, 43, 0.6);
}

/* Fix password toggle icon position */
:deep(.p-password) {
  width: 100%;
}

:deep(.p-password .p-password-input) {
  width: 100%;
  padding-right: 3rem;
}

:deep(.p-password .p-icon-field-right > .p-input-icon) {
  right: 1rem;
}
</style>

<template>
  <div class="min-h-screen flex flex-row-reverse p-8 gap-8" style="background-color: #141204">
    <!-- Phần phải: Gradient background (đảo ngược so với login) -->
    <div class="hidden lg:flex lg:w-5/12 flex-col justify-between gradient-panel text-white p-12 rounded-[2rem]">
      <div class="flex items-center gap-3">
        <i class="pi pi-users text-4xl"></i>
        <span class="text-2xl font-bold">gono.</span>
      </div>

      <div class="max-w-md">
        <p class="text-sm mb-2 opacity-90">Join our community</p>
        <h1 class="text-4xl font-semibold leading-tight">
          Start your journey to better productivity today
        </h1>
      </div>

      <div class="flex gap-2 text-sm opacity-75">
        <span>✓ Free to start</span>
        <span>•</span>
        <span>✓ No credit card</span>
      </div>
    </div>

    <!-- Phần trái: Form đăng ký -->
    <div class="flex-1 flex items-center justify-center p-4">
      <div class="w-full max-w-lg">
        <div class="mb-8">
          <i class="pi pi-user-plus text-4xl mb-4 block" style="color: #E8985E"></i>
          <h2 class="text-4xl font-bold text-white mb-2">Join Gono</h2>
          <p class="text-base" style="color: #A9714B">
            Create your account and unlock your productivity potential.
          </p>
        </div>

        <form @submit.prevent="handleSubmit">
          <!-- 2 cột cho Username và Email -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-5">
            <div>
              <label for="username" class="block font-semibold mb-2 text-sm" style="color: #E8985E">Username</label>
              <InputText 
                id="username" 
                v-model="form.username" 
                class="w-full" 
                placeholder="your_username"
                :invalid="!!errors.username" 
              />
              <small v-if="errors.username" class="text-red-500 text-sm mt-1">{{ errors.username }}</small>
            </div>

            <div>
              <label for="email" class="block font-semibold mb-2 text-sm" style="color: #E8985E">Email</label>
              <InputText 
                id="email" 
                v-model="form.email" 
                type="email" 
                class="w-full" 
                placeholder="you@example.com"
                :invalid="!!errors.email" 
              />
              <small v-if="errors.email" class="text-red-500 text-sm mt-1">{{ errors.email }}</small>
            </div>
          </div>

          <!-- Password -->
          <div class="mb-5">
            <label for="password" class="block font-semibold mb-2 text-sm" style="color: #E8985E">Password</label>
            <Password 
              id="password" 
              v-model="form.password" 
              class="w-full" 
              toggleMask 
              :feedback="true"
              placeholder="Create a strong password" 
              :invalid="!!errors.password" 
            />
            <small v-if="errors.password" class="text-red-500 text-sm mt-1">{{ errors.password }}</small>
          </div>

          <!-- Confirm Password -->
          <div class="mb-6">
            <label for="confirm_password" class="block font-semibold mb-2 text-sm" style="color: #E8985E">Confirm Password</label>
            <Password 
              id="confirm_password" 
              v-model="form.confirm_password" 
              class="w-full" 
              toggleMask 
              :feedback="false"
              placeholder="Confirm your password" 
              :invalid="!!errors.confirm_password" 
            />
            <small v-if="errors.confirm_password" class="text-red-500 text-sm mt-1">{{ errors.confirm_password }}</small>
          </div>

          <!-- Terms & Conditions -->
          <div class="mb-6 text-sm" style="color: #A9714B">
            By signing up, you agree to our 
            <a href="#" class="font-semibold hover:underline" style="color: #E8985E">Terms of Service</a> 
            and 
            <a href="#" class="font-semibold hover:underline" style="color: #E8985E">Privacy Policy</a>
          </div>

          <Button 
            type="submit" 
            label="Create Account" 
            class="w-full py-3.5 text-lg font-semibold rounded-xl" 
            style="color: var(--color-text-primary); background-color: var(--color-bg-tertiary); border-color: var(--color-border)"
            :loading="loading" 
          />

          <!-- Sign in link -->
          <p class="text-center text-sm mt-6" style="color: #A9714B">
            Already have an account?
            <router-link to="/login" class="font-semibold hover:underline" style="color: #E8985E">
              Sign in
            </router-link>
          </p>
        </form>
      </div>
    </div>
  </div>
</template>
