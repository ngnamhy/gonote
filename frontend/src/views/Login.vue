<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const form = ref({
  email: '',
  password: '',
})

const loading = ref(false)
const error = ref('')

async function handleLogin() {
  loading.value = true
  error.value = ''

  try {
    await authStore.login(form.value)

    // await authStore.fetchUser()
    console.log(form.value)

    if (authStore.isAdmin) {
      router.push('/admin')
    } else {
      router.push('/') 
    }

  } catch (err) {
    error.value = err.response?.data?.message || 'Wrong username or password'
  } finally {
    loading.value = false
  }
}
</script>
<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <div class="card shadow">
          <div class="card-body">
            <h3 class="text-center mb-4">Sign in Gonote</h3>

            <form @submit.prevent="handleLogin">
              <div class="mb-3">
                <label>Username</label>
                <input v-model="form.username" class="form-control" required />
              </div>
              <div class="mb-3">
                <label>Password</label>
                <input v-model="form.password" type="password" class="form-control" required />
              </div>

              <button type="submit" class="btn btn-primary w-100" :disabled="loading">
                {{ loading ? 'Loading...' : 'Sign in' }}
              </button>

              <p class="text-danger mt-3 text-center" v-if="error">{{ error }}</p>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

