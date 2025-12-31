<template>
  <div>
    <!-- Ở đây bạn có thể để navbar user khác nếu muốn -->
    <!-- Ví dụ: navbar công khai với logo, search, login/register... -->
    
    <!-- Nếu chưa có navbar user thì để trống cũng được -->
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark shadow-sm">
      <div class="container-fluid">
        <router-link to="/" class="navbar-brand fw-bold">
          <i class="bi bi-journal-text me-2"></i>Gonote
        </router-link>
        <!-- Thêm các link cho user nếu cần -->
      </div>
    </nav>

    <div class="container mt-4">
      <router-view></router-view>
    </div>
  </div>
</template>

<script setup>
// Không cần gì đặc biệt ở đây
</script>