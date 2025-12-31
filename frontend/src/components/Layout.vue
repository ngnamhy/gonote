<template>
  <div>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark shadow-sm">
      <div class="container-fluid">
        <a class="navbar-brand fw-bold" href="#">
          <i class="bi bi-journal-text me-2"></i>Gonote Admin
        </a>

        <button
          class="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarNav"
        >
          <span class="navbar-toggler-icon"></span>
        </button>

        <div class="collapse navbar-collapse" id="navbarNav">
          <ul class="navbar-nav me-auto">
            <li class="nav-item">
              <router-link to="/" class="nav-link" active-class="active">
                <i class="bi bi-house-door"></i> Dashboard
              </router-link>
            </li>
            <li class="nav-item">
              <router-link to="/users" class="nav-link" active-class="active">
                <i class="bi bi-people"></i> Users
              </router-link>
            </li>
            <li class="nav-item">
              <router-link to="/posts" class="nav-link" active-class="active">
                <i class="bi bi-file-post"></i> Posts
              </router-link>
            </li>
          </ul>

          <!-- Phần bên phải: Đăng xuất -->
          <div class="d-flex">
            <button @click="logout" class="btn btn-outline-danger">
              <i class="bi bi-box-arrow-right"></i> Đăng xuất
            </button>
          </div>
        </div>
      </div>
    </nav>

    <!-- Nội dung trang con sẽ hiện ở đây -->
    <div class="container mt-4">
      <router-view></router-view>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { removeToken } from '@/stores/auth'

const router = useRouter()

function logout() {
  removeToken()
  router.push('/login')
}
</script>