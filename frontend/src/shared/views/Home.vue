<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/features/auth/stores/authStore'

const router = useRouter()
const authStore = useAuthStore()
</script>

<template>
  <div class="min-h-screen flex items-center justify-center p-8" style="background-color: #141204">
    <div class="max-w-4xl w-full text-center">
      <!-- Logo -->
      <div class="flex items-center justify-center gap-4 mb-8">
        <div 
          class="flex items-center justify-center w-16 h-16 rounded-2xl text-3xl"
          style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
        >
          <i class="pi pi-book text-white"></i>
        </div>
        <span class="text-5xl font-bold text-white">gono<span style="color: var(--color-accent)">.</span></span>
      </div>

      <!-- Hero Section -->
      <div class="mb-12">
        <h1 class="text-5xl md:text-6xl font-bold text-white mb-6 leading-tight">
          Your Team's Social<br />
          <span style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon)); -webkit-background-clip: text; -webkit-text-fill-color: transparent;">
            Workplace Hub
          </span>
        </h1>
        <p class="text-xl mb-8 max-w-2xl mx-auto" style="color: var(--color-text-secondary)">
          Connect, collaborate, and stay in sync with your team. Share updates, celebrate wins, and build stronger workplace relationships - all in one place.
        </p>
      </div>

      <!-- CTA Buttons -->
      <div class="flex gap-4 justify-center mb-16">
        <router-link to="/register">
          <button 
            class="px-8 py-4 rounded-xl font-semibold text-lg transition-all hover:scale-105 shadow-lg"
            style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon)); color: white"
          >
            <i class="pi pi-user-plus mr-2"></i>
            Get Started
          </button>
        </router-link>
        <router-link to="/login">
          <button 
            class="px-8 py-4 rounded-xl font-semibold text-lg transition-all hover:bg-opacity-10"
            style="border: 2px solid var(--color-accent); color: var(--color-accent)"
          >
            <i class="pi pi-sign-in mr-2"></i>
            Sign In
          </button>
        </router-link>
      </div>

      <!-- Features -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-4xl mx-auto">
        <div 
          class="p-6 rounded-xl"
          style="background-color: var(--color-bg-secondary)"
        >
          <div 
            class="w-14 h-14 rounded-xl flex items-center justify-center mx-auto mb-4"
            style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
          >
            <i class="pi pi-users text-2xl text-white"></i>
          </div>
          <h3 class="font-bold text-lg text-white mb-2">Team Communities</h3>
          <p class="text-sm" style="color: var(--color-text-muted)">
            Join teams and departments to stay connected with your colleagues
          </p>
        </div>

        <div 
          class="p-6 rounded-xl"
          style="background-color: var(--color-bg-secondary)"
        >
          <div 
            class="w-14 h-14 rounded-xl flex items-center justify-center mx-auto mb-4"
            style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
          >
            <i class="pi pi-comment text-2xl text-white"></i>
          </div>
          <h3 class="font-bold text-lg text-white mb-2">Real-time Feed</h3>
          <p class="text-sm" style="color: var(--color-text-muted)">
            Share updates, announcements, and engage with your team instantly
          </p>
        </div>

        <div 
          class="p-6 rounded-xl"
          style="background-color: var(--color-bg-secondary)"
        >
          <div 
            class="w-14 h-14 rounded-xl flex items-center justify-center mx-auto mb-4"
            style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
          >
            <i class="pi pi-bell text-2xl text-white"></i>
          </div>
          <h3 class="font-bold text-lg text-white mb-2">Stay Informed</h3>
          <p class="text-sm" style="color: var(--color-text-muted)">
            Never miss important updates with smart notifications
          </p>
        </div>
      </div>

      <!-- Footer -->
      <p class="mt-12 text-sm" style="color: var(--color-text-muted)">
        Built for teams who work better together
      </p>
    </div>
  </div>
</template>
