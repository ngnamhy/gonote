import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/features/auth/model/authStore'
import { useUserStore } from '@/entities/user/model/userStore'

// Layouts
import UserLayout from '@/shared/components/UserLayout.vue'

const routes = [
  // Landing page
  { 
    path: '/', 
    name: 'home',
    component: () => import('@/shared/views/Home.vue'),
    meta: { guest: true }
  },

  // Auth pages
  { 
    path: '/login', 
    name: 'login',
    component: () => import('@/pages/login/ui/LoginPage.vue'), 
    meta: { guest: true } 
  },
  { 
    path: '/register', 
    name: 'register',
    component: () => import('@/pages/register/ui/RegisterPage.vue'),
    meta: { guest: true }
  },

  // User pages with layout
  {
    path: '/',
    component: UserLayout,         
    meta: { requiresAuth: true },
    children: [
      { 
        path: 'feed', 
        name: 'feed', 
        component: () => import('@/pages/feed/ui/FeedPage.vue')
      },
      { 
        path: 'submit', 
        name: 'submit', 
        component: () => import('@/pages/submit/ui/SubmitPage.vue')
      },
      { 
        path: 'posts/:id', 
        name: 'post-detail', 
        component: () => import('@/pages/post-detail/ui/PostDetailPage.vue')
      },
      { 
        path: 'profile', 
        name: 'profile', 
        component: () => import('@/pages/profile/ui/ProfilePage.vue')
      },
      { 
        path: 'teams', 
        name: 'teams', 
        component: () => import('@/pages/teams/ui/TeamsPage.vue')
      },
      { 
        path: 'notifications', 
        name: 'notifications', 
        component: () => import('@/pages/notifications/ui/NotificationsPage.vue')
      },
    ]
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Navigation guard
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  const userStore = useUserStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/login')
  } else if (to.meta.requiresAdmin && !userStore.isAdmin) {
    next('/feed')
  } else if (to.meta.guest && authStore.isAuthenticated) {
    next('/feed')
  } else {
    next()
  }
})

export default router
