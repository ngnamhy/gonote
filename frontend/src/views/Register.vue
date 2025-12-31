
<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/services/api'
import { setToken } from '@/stores/auth'

const username = ref('')
const password = ref('')
const error = ref('')
const router = useRouter()

async function login() {
  try {
    const res = await api.post('/auth/login', {
      username: username.value,
      password: password.value
    })
    setToken(res.data.token)  
    router.push('/')
  } catch (err) {
    error.value = 'Wrong username or password'
  }
}
</script>

<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h4>Welcome to Gonote!</h4>
          </div>
          <div class="card-body">
            <form @submit.prevent="login">
              <div class="mb-3">
                <label>Username</label>
                <input v-model="username" type="username" class="form-control" required />
              </div>
              <div class="mb-3">
                <label>Password</label>
                <input v-model="password" type="password" class="form-control" required />
              </div>
              <button type="submit" class="btn btn-primary">Login</button>
              <p v-if="error" class="text-danger mt-3">{{ error }}</p>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
