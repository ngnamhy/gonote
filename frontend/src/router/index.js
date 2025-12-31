import { createRouter, createWebHistory } from 'vue-router'


import AdminLayout from '@/components/AdminLayout.vue'
import DefaultLayout from '@/components/DefaultLayout.vue'


import AdminDashboard from '@/views/admin/Dashboard.vue'
import AdminUsers from '@/views/admin/users/Users.vue'
import AdminPosts from '@/views/admin/posts/Posts.vue'

import Home from '@/views/Home.vue'

const routes = [
  {
    path: '/admin',
    component: AdminLayout,           
    children: [
      { path: '', name: 'admin-dashboard', component: AdminDashboard },
      { path: 'users', name: 'admin-users', component: AdminUsers },
      { path: 'posts', name: 'admin-posts', component: AdminPosts },
    ],
    meta: { requiresAuth: true, isAdmin: true } 
  },

  {
    path: '/',
    component: DefaultLayout,         
    children: [
      { path: '', name: 'home', component: Home },
      // { path: 'note/:id', name: 'note-detail', component: NoteDetail },
      // { path: 'profile', name: 'user-profile', component: UserProfile },
      // tất cả route user khác
    ]
  },

  { path: '/login', component: () => import('@/views/Login.vue'), meta: { guest: true } },
  { path: '/register', component: () => import('@/views/Register.vue') },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router