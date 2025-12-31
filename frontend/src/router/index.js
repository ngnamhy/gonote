import { createRouter, createWebHistory } from 'vue-router'
import { getToken } from '@/stores/auth'

import LoginView from '@/views/LoginView.vue'
import DashboardView from '@/views/DashboardView.vue'
import UserList from '@/views/users/UserList.vue'
import PostList from '@/views/posts/PostList.vue'

const routes = [
  { path: '/login', name: 'Login', component: LoginView },
  {
    path: '/',
    component: () => import('@/components/Layout.vue'),
    meta: { requiresAuth: true },
    children: [
      { path: '', name: 'Dashboard', component: DashboardView },
      { path: 'users', name: 'Users', component: UserList },
      { path: 'posts', name: 'Posts', component: PostList },
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Bảo vệ route
router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth && !getToken()) {
    next('/login')
  } else {
    next()
  }
})

export default router