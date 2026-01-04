<template>
  <div class="flex h-screen overflow-hidden" style="background-color: var(--color-bg-primary)">
    <!-- Sidebar -->
    <aside 
      class="w-64 flex flex-col hidden md:flex"
      style="background-color: var(--color-bg-secondary)"
    >
      <!-- Logo Section -->
      <div class="p-6 pb-4">
        <div class="flex items-center gap-3">
          <div 
            class="flex items-center justify-center w-10 h-10 rounded-xl text-xl"
            style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
          >
            <i class="pi pi-shield text-white"></i>
          </div>
          <span class="text-2xl font-bold text-white">Admin</span>
        </div>
        <!-- Soft divider -->
        <div class="mt-4 h-px" style="background: linear-gradient(to right, transparent, var(--color-border-default), transparent); opacity: 0.3"></div>
      </div>

      <!-- Navigation Menu -->
      <nav class="flex-1 p-4 space-y-2">
        <router-link 
          v-for="item in menuItems" 
          :key="item.to"
          :to="item.to"
          v-slot="{ isExactActive }"
        >
          <div 
            class="flex items-center gap-3 px-4 py-3 rounded-xl cursor-pointer transition-all group"
            :class="isExactActive ? 'active-menu-item' : 'menu-item'"
          >
            <i 
              :class="item.icon" 
              class="text-xl transition-colors"
              :style="isExactActive ? 'color: var(--color-accent)' : 'color: var(--color-text-muted)'"
            ></i>
            <span 
              class="font-semibold transition-colors"
              :style="isExactActive ? 'color: var(--color-text-primary)' : 'color: var(--color-text-tertiary)'"
            >
              {{ item.label }}
            </span>
          </div>
        </router-link>
      </nav>

      <!-- Footer: User Profile & Logout -->
      <div class="p-4">
        <!-- Soft divider -->
        <div class="mb-4 h-px" style="background: linear-gradient(to right, transparent, var(--color-border-default), transparent); opacity: 0.3"></div>
        
        <!-- User Profile -->
        <div class="flex items-center gap-3 mb-3 px-4 py-2">
          <div 
            class="flex items-center justify-center w-10 h-10 rounded-full text-lg"
            style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
          >
            <i class="pi pi-user text-white"></i>
          </div>
          <div class="flex flex-col">
            <span class="font-bold text-white">{{ user?.username || "Admin" }}</span>
            <span class="text-xs" style="color: var(--color-text-muted)">{{ user?.role || "Administrator" }}</span>
          </div>
        </div>
        
        <!-- Logout Button -->
        <div 
          @click="logout"
          class="flex items-center gap-3 px-4 py-3 rounded-xl cursor-pointer transition-all hover:bg-opacity-80"
          style="background-color: var(--color-bg-tertiary)"
        >
          <i class="pi pi-sign-out text-xl" style="color: var(--color-text-muted)"></i>
          <span class="font-semibold" style="color: var(--color-text-tertiary)">Logout</span>
        </div>
      </div>
    </aside>

    <!-- Main Content Area -->
    <div class="flex-1 flex flex-col overflow-hidden">
      <!-- Top Bar (optional - có thể thêm breadcrumb, search, notifications) -->
      <!-- <header 
        class="px-6 py-4"
      >
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
          <i class="pi pi-bars text-xl md:hidden cursor-pointer" style="color: var(--color-text-secondary)"></i>
          <h2 class="text-xl font-bold text-white">Admin Panel</h2>
        </div>
        <div class="flex items-center gap-4">
          <button 
            class="w-10 h-10 rounded-full flex items-center justify-center transition-all hover:scale-110"
            style="background-color: var(--color-bg-tertiary)"
          >
            <i class="pi pi-bell" style="color: var(--color-accent)"></i>
          </button>
        </div>
        </div> -->
        <!-- Soft divider -->
        <!-- <div class="mt-4 h-px" style="background: linear-gradient(to right, transparent, var(--color-border-default), transparent); opacity: 0.3"></div>
      </header> -->

      <!-- Main Content -->
      <main class="flex-1 overflow-hidden">
        <div class="h-full overflow-y-auto p-6" style="background-color: var(--color-bg-secondary)">
          <router-view />
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/features/auth/stores/authStore'

const router = useRouter()
const authStore = useAuthStore()
const user = ref(authStore.user)

function logout() {
  authStore.logout()
  router.push('/login')
}

const menuItems = ref([
  {
    label: 'Dashboard',
    icon: 'pi pi-home',
    to: '/admin'
  },
  {
    label: 'Users',
    icon: 'pi pi-users',
    to: '/admin/users'
  },
  {
    label: 'Posts',
    icon: 'pi pi-book',
    to: '/admin/posts'
  },
])
</script>

<style scoped>
.menu-item {
  background-color: transparent;
}

.menu-item:hover {
  background-color: var(--color-bg-tertiary);
}

.active-menu-item {
  background: linear-gradient(90deg, var(--color-accent-muted), transparent);
  position: relative;
}

.active-menu-item::before {
  content: '';
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 3px;
  height: 70%;
  background-color: var(--color-accent);
  border-radius: 0 2px 2px 0;
}

.active-menu-item:hover {
  background: linear-gradient(90deg, var(--color-accent-muted), transparent);
}
</style>