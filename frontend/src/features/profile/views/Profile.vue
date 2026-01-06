<script setup>
import { ref, onMounted, computed } from 'vue'
import { useAuthStore } from '@/features/auth/stores/authStore'
import api from '@/shared/services/api'
import InputText from 'primevue/inputtext'
import Textarea from 'primevue/textarea'
import Button from 'primevue/button'

const authStore = useAuthStore()
const user = computed(() => authStore.user)

const profile = ref({
  name: '',
  username: '',
  email: '',
  bio: '',
  avatar: ''
})

const loading = ref(false)
const saving = ref(false)
const error = ref(null)
const successMessage = ref(null)

// Helper function to get user initials
function getInitials(name) {
  if (!name) return 'U'
  return name
    .split(' ')
    .map(word => word[0])
    .join('')
    .toUpperCase()
    .slice(0, 2)
}

async function fetchProfile() {
  try {
    loading.value = true
    error.value = null
    
    // Get profile from auth store or API
    if (user.value) {
      profile.value = {
        name: user.value.name || '',
        username: user.value.username || '',
        email: user.value.email || '',
        bio: user.value.bio || '',
        avatar: user.value.avatar || ''
      }
    }
  } catch (err) {
    error.value = err.message || 'Failed to load profile'
    console.error('Error fetching profile:', err)
  } finally {
    loading.value = false
  }
}

async function updateProfile() {
  if (!profile.value.name.trim() || !profile.value.email.trim()) {
    error.value = 'Name and email are required'
    return
  }

  try {
    saving.value = true
    error.value = null
    successMessage.value = null
    
    const response = await api.put('/users/profile', {
      name: profile.value.name,
      username: profile.value.username,
      email: profile.value.email,
      bio: profile.value.bio
    })
    
    authStore.updateUser(response.data.data || response.data)
    
    successMessage.value = 'Profile updated successfully!'
    
    setTimeout(() => {
      successMessage.value = null
    }, 3000)
  } catch (err) {
    error.value = err.response?.data?.message || 'Failed to update profile'
    console.error('Error updating profile:', err)
  } finally {
    saving.value = false
  }
}

onMounted(() => {
  fetchProfile()
})
</script>

<template>
  <div class="min-h-full p-6" style="background-color: var(--color-bg-primary)">
    <div class="max-w-4xl mx-auto">
      <!-- Header -->
      <div class="mb-6">
        <h1 class="text-3xl font-bold text-white mb-2">Profile Settings</h1>
        <p style="color: var(--color-text-muted)">Manage your account information</p>
      </div>

      <!-- Loading State -->
      <div 
        v-if="loading"
        class="rounded-xl p-8 text-center"
        style="background-color: var(--color-bg-secondary)"
      >
        <i class="pi pi-spinner pi-spin text-3xl" style="color: var(--color-accent)"></i>
        <p class="mt-4" style="color: var(--color-text-muted)">Loading profile...</p>
      </div>

      <!-- Profile Form -->
      <div v-else class="space-y-4">
        <!-- Avatar Section -->
        <div 
          class="rounded-xl p-6"
          style="background-color: var(--color-bg-secondary)"
        >
          <h2 class="text-xl font-bold text-white mb-4">Avatar</h2>
          <div class="flex items-center gap-6">
            <div 
              class="w-24 h-24 rounded-full flex items-center justify-center text-3xl flex-shrink-0"
              style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))"
            >
              <span class="text-white font-bold">{{ getInitials(profile.name) }}</span>
            </div>
            <div>
              <p class="text-white font-semibold mb-1">Profile Picture</p>
              <p class="text-sm mb-3" style="color: var(--color-text-muted)">
                Upload a new avatar or use your initials
              </p>
              <Button 
                label="Upload Image" 
                severity="secondary"
                icon="pi pi-upload"
                size="small"
                disabled
              />
            </div>
          </div>
        </div>

        <!-- Profile Information -->
        <div 
          class="rounded-xl p-6"
          style="background-color: var(--color-bg-secondary)"
        >
          <h2 class="text-xl font-bold text-white mb-4">Profile Information</h2>

          <!-- Success message -->
          <div v-if="successMessage" class="mb-4 p-3 bg-green-500/10 border border-green-500 rounded-lg">
            <div class="flex items-center gap-2">
              <i class="pi pi-check-circle text-green-500"></i>
              <p class="text-green-500 text-sm">{{ successMessage }}</p>
            </div>
          </div>

          <!-- Error message -->
          <div v-if="error" class="mb-4 p-3 bg-red-500/10 border border-red-500 rounded-lg">
            <div class="flex items-center gap-2">
              <i class="pi pi-exclamation-circle text-red-500"></i>
              <p class="text-red-500 text-sm">{{ error }}</p>
            </div>
          </div>

          <form @submit.prevent="updateProfile" class="space-y-5">
            <!-- Name -->
            <div>
              <label for="name" class="block font-semibold mb-2 text-sm" style="color: var(--color-accent)">
                Full Name *
              </label>
              <InputText 
                id="name"
                v-model="profile.name" 
                class="w-full" 
                placeholder="Enter your full name"
                :disabled="saving"
                size="large"
              />
            </div>

            <!-- Username -->
            <div>
              <label for="username" class="block font-semibold mb-2 text-sm" style="color: var(--color-accent)">
                Username
              </label>
              <InputText 
                id="username"
                v-model="profile.username" 
                class="w-full" 
                placeholder="@username"
                :disabled="saving"
                size="large"
              />
            </div>

            <!-- Email -->
            <div>
              <label for="email" class="block font-semibold mb-2 text-sm" style="color: var(--color-accent)">
                Email Address *
              </label>
              <InputText 
                id="email"
                v-model="profile.email" 
                type="email"
                class="w-full" 
                placeholder="your@email.com"
                :disabled="saving"
                size="large"
              />
            </div>

            <!-- Bio -->
            <div>
              <label for="bio" class="block font-semibold mb-2 text-sm" style="color: var(--color-accent)">
                Bio
              </label>
              <Textarea 
                id="bio"
                v-model="profile.bio" 
                class="w-full" 
                rows="4"
                placeholder="Tell us about yourself..."
                :disabled="saving"
              />
            </div>

            <!-- Actions -->
            <div class="flex gap-3 justify-end pt-4">
              <Button 
                type="button"
                label="Cancel" 
                severity="secondary"
                @click="fetchProfile"
                :disabled="saving"
                size="large"
              />
              <Button 
                type="submit"
                label="Save Changes" 
                :loading="saving"
                size="large"
                icon="pi pi-check"
              />
            </div>
          </form>
        </div>

        <!-- Account Settings -->
        <div 
          class="rounded-xl p-6"
          style="background-color: var(--color-bg-secondary)"
        >
          <h2 class="text-xl font-bold text-white mb-4">Account Settings</h2>
          
          <div class="space-y-3">
            <button 
              class="w-full flex items-center justify-between p-4 rounded-lg transition-all hover:bg-opacity-80"
              style="background-color: var(--color-bg-tertiary)"
            >
              <div class="flex items-center gap-3">
                <i class="pi pi-lock text-xl" style="color: var(--color-accent)"></i>
                <div class="text-left">
                  <p class="text-white font-semibold">Change Password</p>
                  <p class="text-sm" style="color: var(--color-text-muted)">Update your password</p>
                </div>
              </div>
              <i class="pi pi-chevron-right" style="color: var(--color-text-muted)"></i>
            </button>

            <button 
              class="w-full flex items-center justify-between p-4 rounded-lg transition-all hover:bg-opacity-80"
              style="background-color: var(--color-bg-tertiary)"
            >
              <div class="flex items-center gap-3">
                <i class="pi pi-shield text-xl" style="color: var(--color-accent)"></i>
                <div class="text-left">
                  <p class="text-white font-semibold">Privacy Settings</p>
                  <p class="text-sm" style="color: var(--color-text-muted)">Control your privacy preferences</p>
                </div>
              </div>
              <i class="pi pi-chevron-right" style="color: var(--color-text-muted)"></i>
            </button>

            <button 
              class="w-full flex items-center justify-between p-4 rounded-lg transition-all hover:bg-opacity-80"
              style="background-color: var(--color-bg-tertiary)"
            >
              <div class="flex items-center gap-3">
                <i class="pi pi-bell text-xl" style="color: var(--color-accent)"></i>
                <div class="text-left">
                  <p class="text-white font-semibold">Notification Preferences</p>
                  <p class="text-sm" style="color: var(--color-text-muted)">Manage your notifications</p>
                </div>
              </div>
              <i class="pi pi-chevron-right" style="color: var(--color-text-muted)"></i>
            </button>
          </div>
        </div>

        <!-- Danger Zone -->
        <div 
          class="rounded-xl p-6 border"
          style="background-color: var(--color-bg-secondary); border-color: #ef4444"
        >
          <h2 class="text-xl font-bold text-red-500 mb-4">Danger Zone</h2>
          
          <div class="space-y-3">
            <button 
              class="w-full flex items-center justify-between p-4 rounded-lg transition-all hover:bg-red-500/10"
              style="background-color: var(--color-bg-tertiary); border: 1px solid #ef4444"
            >
              <div class="flex items-center gap-3">
                <i class="pi pi-trash text-xl text-red-500"></i>
                <div class="text-left">
                  <p class="text-red-500 font-semibold">Delete Account</p>
                  <p class="text-sm" style="color: var(--color-text-muted)">Permanently delete your account and all data</p>
                </div>
              </div>
              <i class="pi pi-chevron-right text-red-500"></i>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
