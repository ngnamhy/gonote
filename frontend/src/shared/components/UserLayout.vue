<template>
  <div class="flex h-screen overflow-hidden" style="background-color: var(--color-bg-primary)">
    <!-- Sidebar -->
    <aside 
      :class="['sidebar', { 'sidebar-open': sidebarOpen }]"
      style="background-color: var(--color-bg-secondary)"
    >
      <!-- Logo Section -->
      <div class="p-6 pb-4">
        <div class="flex items-center gap-3">
          <div 
            class="flex items-center justify-center w-10 h-10 rounded-xl text-xl"
            style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
          >
            <i class="pi pi-book text-white"></i>
          </div>
          <span class="text-2xl font-bold text-white">gono.</span>
        </div>
        <!-- Soft divider -->
        <div class="mt-4 h-px" style="background: linear-gradient(to right, transparent, var(--color-border-default), transparent); opacity: 0.3"></div>
      </div>

      <!-- Navigation Menu -->
      <nav class="p-4 space-y-2">
        <router-link 
          v-for="item in menuItems" 
          :key="item.to"
          :to="item.to"
          v-slot="{ isExactActive }"
          @click="sidebarOpen = false"
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

      <!-- My Teams Section -->
      <div class="px-4 pb-4">
        <!-- Soft divider -->
        <div class="mb-4 h-px" style="background: linear-gradient(to right, transparent, var(--color-border-default), transparent); opacity: 0.3"></div>
        
        <h3 class="font-bold text-white mb-3 px-4 text-sm flex items-center gap-2">
          <i class="pi pi-users" style="color: var(--color-accent)"></i>
          My Teams
        </h3>
        <div class="space-y-2 max-h-64 overflow-y-auto">
          <div 
            v-for="team in myTeams" 
            :key="team.id"
            class="flex items-center gap-3 p-3 rounded-lg cursor-pointer transition-all hover:bg-opacity-80"
            style="background-color: var(--color-bg-tertiary)"
          >
            <div 
              class="w-8 h-8 rounded-lg flex items-center justify-center flex-shrink-0"
              :style="`background-color: ${team.color}20`"
            >
              <i :class="`pi ${team.icon} text-sm`" :style="`color: ${team.color}`"></i>
            </div>
            <div class="flex-1 min-w-0">
              <p class="font-semibold text-sm text-white truncate">{{ team.name }}</p>
              <p class="text-xs" style="color: var(--color-text-muted)">{{ team.members }} members</p>
            </div>
          </div>
        </div>
        <button 
          @click="$router.push('/teams')"
          class="w-full mt-3 px-4 py-2 rounded-lg font-semibold text-sm transition-all hover:bg-opacity-80 flex items-center justify-center gap-2"
          style="background-color: var(--color-bg-tertiary); color: var(--color-accent)"
        >
          <i class="pi pi-plus"></i>
          Join Team
        </button>
      </div>

      <!-- Spacer to push user profile to bottom -->
      <div class="flex-1"></div>

      <!-- User Profile / Auth Section -->
      <div class="p-4">
        <!-- Soft divider -->
        <div class="mb-4 h-px" style="background: linear-gradient(to right, transparent, var(--color-border-default), transparent); opacity: 0.3"></div>
        
        <!-- If logged in -->
        <div v-if="isAuthenticated">
          <div class="flex items-center gap-3 mb-3 px-4 py-2">
            <div 
              class="flex items-center justify-center w-10 h-10 rounded-full text-lg"
              style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
            >
              <i class="pi pi-user text-white"></i>
            </div>
            <div class="flex flex-col">
              <span class="font-bold text-white">{{ user?.username }}</span>
              <span class="text-xs" style="color: var(--color-text-muted)">{{ user?.email }}</span>
            </div>
          </div>
          <div 
            @click="logout"
            class="flex items-center gap-3 px-4 py-3 rounded-xl cursor-pointer transition-all hover:bg-opacity-80"
            style="background-color: var(--color-bg-tertiary)"
          >
            <i class="pi pi-sign-out text-xl" style="color: var(--color-text-muted)"></i>
            <span class="font-semibold" style="color: var(--color-text-tertiary)">Logout</span>
          </div>
        </div>

        <!-- If not logged in -->
        <div v-else class="space-y-2">
          <router-link to="/login" @click="sidebarOpen = false">
            <div 
              class="flex items-center gap-3 px-4 py-3 rounded-xl cursor-pointer transition-all"
              style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
            >
              <i class="pi pi-sign-in text-xl text-white"></i>
              <span class="font-semibold text-white">Login</span>
            </div>
          </router-link>
          <router-link to="/register" @click="sidebarOpen = false">
            <div 
              class="flex items-center gap-3 px-4 py-3 rounded-xl cursor-pointer transition-all hover:bg-opacity-80"
              style="background-color: var(--color-bg-tertiary)"
            >
              <i class="pi pi-user-plus text-xl" style="color: var(--color-accent)"></i>
              <span class="font-semibold" style="color: var(--color-text-tertiary)">Register</span>
            </div>
          </router-link>
        </div>
      </div>
    </aside>

    <!-- Overlay for mobile -->
    <div 
      v-if="sidebarOpen"
      @click="sidebarOpen = false"
      class="fixed inset-0 bg-black bg-opacity-50 z-30 md:hidden"
    ></div>

    <!-- Main Content Area -->
    <div class="flex-1 flex flex-col overflow-hidden">
      <!-- Top Bar -->
      <header 
        class="px-6 py-4 flex items-center justify-between md:justify-end"
        style="background-color: var(--color-bg-secondary)"
      >
        <!-- Mobile menu button -->
        <button 
          @click="sidebarOpen = !sidebarOpen"
          class="md:hidden w-10 h-10 rounded-lg flex items-center justify-center transition-all hover:bg-opacity-80"
          style="background-color: var(--color-bg-tertiary)"
        >
          <i class="pi pi-bars text-xl" style="color: var(--color-text-secondary)"></i>
        </button>

        <!-- Search bar (optional) -->
        <div class="hidden lg:flex items-center gap-2 px-4 py-2 rounded-lg max-w-md" style="background-color: var(--color-bg-tertiary)">
          <i class="pi pi-search" style="color: var(--color-text-muted)"></i>
          <input 
            type="text" 
            placeholder="Search teams, posts, people..."
            class="bg-transparent outline-none flex-1 text-white placeholder-opacity-50"
            style="color: var(--color-text-primary)"
          />
        </div>
      </header>

      <!-- Main Content -->
      <main class="flex-1 overflow-hidden">
        <div class="h-full overflow-y-auto">
          <router-view />
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/features/auth/stores/authStore'

const router = useRouter()
const authStore = useAuthStore()
const sidebarOpen = ref(false)

const isAuthenticated = computed(() => authStore.isAuthenticated)
const user = computed(() => authStore.user)

const myTeams = ref([
  { id: 1, name: 'Engineering', icon: 'pi-code', color: '#3B82F6', members: 24 },
  { id: 2, name: 'Marketing', icon: 'pi-megaphone', color: '#10B981', members: 12 },
  { id: 3, name: 'Design', icon: 'pi-palette', color: '#8B5CF6', members: 8 },
  { id: 4, name: 'General', icon: 'pi-comment', color: 'var(--color-accent)', members: 45 },
])

function logout() {
  authStore.logout()
  router.push('/login')
  sidebarOpen.value = false
}

const menuItems = ref([
  {
    label: 'Feed',
    icon: 'pi pi-home',
    to: '/feed'
  },
  {
    label: 'Teams',
    icon: 'pi pi-users',
    to: '/teams'
  },
  {
    label: 'Notifications',
    icon: 'pi pi-bell',
    to: '/notifications'
  },
])
</script>

<style scoped>
.sidebar {
  @apply w-64 flex flex-col;
  @apply fixed md:static inset-y-0 left-0 z-40;
  @apply transform -translate-x-full md:translate-x-0 transition-transform duration-300;
}

.sidebar-open {
  @apply translate-x-0;
}

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