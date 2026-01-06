<template>
  <div class="min-h-screen flex p-8 gap-8" style="background-color: #141204">
    <!-- Left: Gradient background -->
    <div class="hidden lg:flex lg:w-1/2 flex-col justify-between gradient-panel text-white p-12 rounded-[2rem]">
      <router-link to="/" class="flex items-center gap-3 hover:opacity-80 transition-opacity w-fit">
        <i class="pi pi-sparkles text-4xl"></i>
        <span class="text-2xl font-bold">gono.</span>
      </router-link>

      <div class="max-w-md">
        <p class="text-sm mb-2 opacity-90">Join us today</p>
        <h1 class="text-4xl font-semibold leading-tight">
          Create your account and start sharing with your team
        </h1>
      </div>

      <div></div>
    </div>

    <!-- Right: Form -->
    <div class="flex-1 flex items-center justify-center p-4">
      <div class="w-full max-w-md p-10 rounded-3xl">
        <div class="text-center mb-8">
          <i class="pi pi-sparkles text-5xl mb-4" style="color: #E8985E"></i>
          <h2 class="text-3xl font-bold text-white mb-2">Create Account</h2>
          <p class="text-sm" style="color: #A9714B">
            Join gono and collaborate with your team
          </p>
        </div>

        <RegisterForm />

        <p class="text-center text-sm mt-6" style="color: #A9714B">
          Already have an account?
          <router-link to="/login" class="font-semibold hover:underline" style="color: #E8985E">
            Sign in
          </router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import RegisterForm from '@/features/auth/ui/RegisterForm.vue'
</script>

<style scoped>
.gradient-panel {
  background: linear-gradient(135deg, #141204 0%, #262A10 40%, #54442B 100%);
  box-shadow: 0 20px 60px rgba(20, 18, 4, 0.6);
}
</style>
