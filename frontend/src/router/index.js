import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/features/auth/stores/authStore'

// Layouts
import AdminLayout from '@/shared/components/AdminLayout.vue'
import UserLayout from '@/shared/components/UserLayout.vue'

// Admin views
import AdminDashboard from '@/features/admin/views/Dashboard.vue'
import AdminUsers from '@/features/admin/views/users/Users.vue'
import AdminPosts from '@/features/admin/views/posts/Posts.vue'

// Shared views
import Home from '@/shared/views/Home.vue'
import Teams from '@/shared/views/Teams.vue'
import Notifications from '@/shared/views/Notifications.vue'

// Feed views
import Feed from '@/features/feed/views/Feed.vue'
import Submit from '@/features/feed/views/Submit.vue'
import PostDetail from '@/features/feed/views/PostDetail.vue'

// Profile views
import Profile from '@/features/profile/views/Profile.vue'

const routes = [
  // Landing page (no layout)
  { 
    path: '/', 
    name: 'home',
    component: Home,
    meta: { guest: true }
  },

  {
    path: '/admin',
    component: AdminLayout,           
    children: [
      { path: '', name: 'admin-dashboard', component: AdminDashboard },
      { path: 'users', name: 'admin-users', component: AdminUsers },
      { path: 'posts', name: 'admin-posts', component: AdminPosts },
    ],
    meta: { requiresAuth: true, requiresAdmin: true } 
  },

  {
    path: '/',
    component: UserLayout,         
    children: [
      { path: 'feed', name: 'feed', component: Feed },
      { path: 'submit', name: 'submit', component: Submit },
      { path: 'posts/:id', name: 'post-detail', component: PostDetail },
      { path: 'teams', name: 'teams', component: Teams },
      { path: 'notifications', name: 'notifications', component: Notifications },
      { path: 'profile', name: 'profile', component: Profile },
    ]
  },

  { 
    path: '/login', 
    name: 'login',
    component: () => import('@/features/auth/views/Login.vue'), 
    meta: { guest: true } 
  },
  { 
    path: '/register', 
    name: 'register',
    component: () => import('@/features/auth/views/Register.vue') 
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Navigation guard
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/login')
  } else if (to.meta.requiresAdmin && !authStore.isAdmin) {
    next('/feed')
  } else if (to.meta.guest && authStore.isAuthenticated) {
    next('/feed')
  } else {
    next()
  }
})

export default router