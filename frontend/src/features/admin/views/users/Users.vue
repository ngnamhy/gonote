<script setup>
import { ref, onMounted, computed } from 'vue'
import api from '@/shared/services/api'
import { useAuthStore } from '@/features/auth/stores/authStore'

const authStore = useAuthStore()
const users = ref([])
const showCreateForm = ref(false)
const editingUser = ref(null)
const disablingUser = ref(null)
const form = ref({ email: '', username: '', password: '', confirm_password: '', avatar: '' })
const error = ref('')
const uploadingAvatar = ref(false)
const avatarPreview = ref('')

const currentUserId = computed(() => authStore.user?.id)

async function loadUsers() {
  const res = await api.get('/users')
  users.value = res.data
}

async function handleAvatarUpload(event) {
  const file = event.target.files[0]
  if (!file) return
  
  // Validate file type
  if (!file.type.startsWith('image/')) {
    error.value = 'Please select an image file'
    return
  }
  
  // Validate file size (max 5MB)
  if (file.size > 5 * 1024 * 1024) {
    error.value = 'Image size must be less than 5MB'
    return
  }
  
  try {
    uploadingAvatar.value = true
    error.value = ''
    
    const formData = new FormData()
    formData.append('image', file)
    
    const response = await api.post('/upload/avatar', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    
    form.value.avatar = response.data.url
    avatarPreview.value = response.data.url
  } catch (err) {
    error.value = 'Failed to upload image'
    console.error('Upload error:', err)
  } finally {
    uploadingAvatar.value = false
  }
}

async function saveUser() {
  try { 
    if (editingUser.value) {
      await api.put(`/users/${editingUser.value.id}`, {
        email: form.value.email,
        username: form.value.username,
        avatar: form.value.avatar
      })
    } else {
      await api.post('/users', form.value)
    }
    error.value = ""
    closeForm()
    loadUsers()
  }
  catch (err) { 
    if (err.response) {
      const backendMessage = err.response.data.message || 
                             err.response.data.error || 
                             JSON.stringify(err.response.data)
      error.value = backendMessage
    } else {
      error.value = 'Error connecting to server'
    }
    console.error(err)
  }
}

function showDisableDialog(user) {
  disablingUser.value = user
}

async function confirmDisable() {
  try {
    await api.put(`/users/${disablingUser.value.id}/disable`)
    disablingUser.value = null
    loadUsers()
  } catch (err) {
    console.error('Error disabling user:', err)
  }
}

function cancelDisable() {
  disablingUser.value = null
}

function editUser(user) {
  editingUser.value = user
  form.value = { email: user.email, username: user.username, password: '', confirm_password: '', avatar: user.avatar || '' }
  avatarPreview.value = user.avatar || ''
}

function closeForm() {
  showCreateForm.value = false
  editingUser.value = null
  form.value = { email: '', username: '', password: '', confirm_password: '', avatar: '' }
  avatarPreview.value = ''
  error.value = ''
}

onMounted(loadUsers)
</script>

<template>
  <div>
    <!-- Header -->
    <div class="flex items-center justify-between mb-8">
      <div>
        <h1 class="text-4xl font-bold text-white mb-2">Users Management</h1>
        <p class="text-lg" style="color: var(--color-text-tertiary)">
          Manage user accounts and permissions
        </p>
      </div>
      <button 
        @click="showCreateForm = true"
        class="px-6 py-3 rounded-xl font-semibold transition-all hover:scale-105"
        style="background-color: var(--color-accent); color: white"
      >
        <i class="pi pi-plus mr-2"></i>
        Add User
      </button>
    </div>

    <!-- Users Table -->
    <div 
      class="rounded-2xl border overflow-hidden"
      style="background-color: var(--color-bg-secondary); border-color: var(--color-border-default)"
    >
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr style="background-color: var(--color-bg-tertiary)">
              <th class="px-6 py-4 text-left text-sm font-bold text-white">ID</th>
              <th class="px-6 py-4 text-left text-sm font-bold text-white">Avatar</th>
              <th class="px-6 py-4 text-left text-sm font-bold text-white">Email</th>
              <th class="px-6 py-4 text-left text-sm font-bold text-white">Username</th>
              <th class="px-6 py-4 text-left text-sm font-bold text-white">Role</th>
              <th class="px-6 py-4 text-left text-sm font-bold text-white">Status</th>
              <th class="px-6 py-4 text-left text-sm font-bold text-white">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr 
              v-for="user in users" 
              :key="user.id"
              class="border-t transition-colors hover:bg-opacity-50"
              style="border-color: var(--color-border-default)"
              :style="{ 'background-color': 'transparent' }"
            >
              <td class="px-6 py-4 text-white">{{ user.id }}</td>
              <td class="px-6 py-4">
                <div class="w-10 h-10 rounded-full overflow-hidden" style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))">
                  <img v-if="user.avatar" :src="user.avatar" alt="Avatar" class="w-full h-full object-cover" />
                  <div v-else class="w-full h-full flex items-center justify-center">
                    <i class="pi pi-user text-white"></i>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4" style="color: var(--color-text-secondary)">{{ user.email }}</td>
              <td class="px-6 py-4 text-white font-semibold">{{ user.username }}</td>
              <td class="px-6 py-4">
                <span 
                  class="px-3 py-1 rounded-full text-xs font-semibold"
                  :style="user.role === 'admin' 
                    ? 'background-color: var(--color-accent-muted); color: var(--color-accent)' 
                    : 'background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)'"
                >
                  {{ user.role || 'user' }}
                </span>
              </td>
              <td class="px-6 py-4">
                <span 
                  class="px-3 py-1 rounded-full text-xs font-semibold"
                  :style="user.is_active !== false
                    ? 'background-color: rgba(34, 197, 94, 0.1); color: rgb(34, 197, 94)' 
                    : 'background-color: rgba(239, 68, 68, 0.1); color: rgb(239, 68, 68)'"
                >
                  {{ user.is_active !== false ? 'Active' : 'Disabled' }}
                </span>
              </td>
              <td class="px-6 py-4">
                <div class="flex gap-2">
                  <button 
                    @click="editUser(user)"
                    class="px-4 py-2 rounded-lg font-semibold text-sm transition-all hover:scale-105"
                    style="background-color: var(--color-accent-muted); color: var(--color-accent)"
                  >
                    <i class="pi pi-pencil mr-1"></i>
                    Edit
                  </button>
                  <button 
                    @click="showDisableDialog(user)"
                    :disabled="user.id === currentUserId"
                    class="px-4 py-2 rounded-lg font-semibold text-sm transition-all"
                    :class="user.id === currentUserId ? 'opacity-50 cursor-not-allowed' : 'hover:scale-105'"
                    style="background-color: rgba(239, 68, 68, 0.1); color: rgb(239, 68, 68)"
                  >
                    <i class="pi pi-ban mr-1"></i>
                    Disable
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Modal Form -->
    <div 
      v-if="showCreateForm || editingUser" 
      class="fixed inset-0 flex items-center justify-center z-50"
      style="background-color: rgba(0, 0, 0, 0.7)"
      @click.self="closeForm"
    >
      <div 
        class="rounded-2xl p-8 max-w-md w-full mx-4"
        style="background-color: var(--color-bg-secondary)"
      >
        <h3 class="text-2xl font-bold text-white mb-6">
          {{ editingUser ? 'Edit User' : 'Create New User' }}
        </h3>
        
        <div class="space-y-4">
          <!-- Avatar Upload -->
          <div>
            <label class="block text-sm font-semibold mb-2 text-white">Avatar</label>
            <div class="flex items-center gap-4">
              <div class="w-20 h-20 rounded-full overflow-hidden" style="background: linear-gradient(135deg, var(--color-accent), var(--color-cinnamon))">
                <img v-if="avatarPreview" :src="avatarPreview" alt="Avatar preview" class="w-full h-full object-cover" />
                <div v-else class="w-full h-full flex items-center justify-center">
                  <i class="pi pi-user text-white text-2xl"></i>
                </div>
              </div>
              <div class="flex-1">
                <label 
                  class="px-4 py-2 rounded-lg font-semibold text-sm cursor-pointer transition-all hover:scale-105 inline-flex items-center gap-2"
                  style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
                >
                  <i class="pi pi-upload"></i>
                  {{ uploadingAvatar ? 'Uploading...' : 'Choose Image' }}
                  <input 
                    type="file" 
                    @change="handleAvatarUpload" 
                    accept="image/*" 
                    class="hidden"
                    :disabled="uploadingAvatar"
                  />
                </label>
                <p class="text-xs mt-2" style="color: var(--color-text-muted)">Max size: 5MB. JPG, PNG, GIF</p>
              </div>
            </div>
          </div>
          
          <div>
            <label class="block text-sm font-semibold mb-2 text-white">Email</label>
            <input 
              v-model="form.email" 
              type="email"
              class="w-full px-4 py-3 rounded-xl border outline-none transition-all"
              style="background-color: var(--color-bg-tertiary); border-color: var(--color-border-default); color: white"
            />
          </div>
          
          <div>
            <label class="block text-sm font-semibold mb-2 text-white">Username</label>
            <input 
              v-model="form.username" 
              class="w-full px-4 py-3 rounded-xl border outline-none transition-all"
              style="background-color: var(--color-bg-tertiary); border-color: var(--color-border-default); color: white"
            />
          </div>
          
          <div v-if="!editingUser">
            <label class="block text-sm font-semibold mb-2 text-white">Password</label>
            <input 
              v-model="form.password" 
              type="password"
              class="w-full px-4 py-3 rounded-xl border outline-none transition-all"
              style="background-color: var(--color-bg-tertiary); border-color: var(--color-border-default); color: white"
            />
          </div>
          
          <div v-if="!editingUser">
            <label class="block text-sm font-semibold mb-2 text-white">Confirm Password</label>
            <input 
              v-model="form.confirm_password" 
              type="password"
              class="w-full px-4 py-3 rounded-xl border outline-none transition-all"
              style="background-color: var(--color-bg-tertiary); border-color: var(--color-border-default); color: white"
            />
          </div>
          
          <p v-if="error" class="text-sm font-semibold" style="color: rgb(239, 68, 68)">
            {{ error }}
          </p>
        </div>
        
        <div class="flex gap-3 mt-6">
          <button 
            @click="closeForm"
            class="flex-1 px-6 py-3 rounded-xl font-semibold transition-all hover:scale-105"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            Cancel
          </button>
          <button 
            @click="saveUser"
            class="flex-1 px-6 py-3 rounded-xl font-semibold transition-all hover:scale-105"
            style="background-color: var(--color-accent); color: white"
          >
            Save
          </button>
        </div>
      </div>
    </div>

    <!-- Disable Confirmation Modal -->
    <div 
      v-if="disablingUser" 
      class="fixed inset-0 flex items-center justify-center z-50"
      style="background-color: rgba(0, 0, 0, 0.7)"
      @click.self="cancelDisable"
    >
      <div 
        class="rounded-2xl p-8 max-w-md w-full mx-4"
        style="background-color: var(--color-bg-secondary)"
      >
        <div class="flex items-center gap-4 mb-6">
          <div 
            class="w-12 h-12 rounded-full flex items-center justify-center"
            style="background-color: rgba(239, 68, 68, 0.1)"
          >
            <i class="pi pi-exclamation-triangle text-2xl" style="color: rgb(239, 68, 68)"></i>
          </div>
          <h3 class="text-2xl font-bold text-white">Disable User</h3>
        </div>
        
        <p class="text-lg mb-2" style="color: var(--color-text-secondary)">
          Are you sure you want to disable this user?
        </p>
        <p class="font-semibold text-white mb-6">
          {{ disablingUser.username }} ({{ disablingUser.email }})
        </p>
        
        <div class="flex gap-3">
          <button 
            @click="cancelDisable"
            class="flex-1 px-6 py-3 rounded-xl font-semibold transition-all hover:scale-105"
            style="background-color: var(--color-bg-tertiary); color: var(--color-text-tertiary)"
          >
            Cancel
          </button>
          <button 
            @click="confirmDisable"
            class="flex-1 px-6 py-3 rounded-xl font-semibold transition-all hover:scale-105"
            style="background-color: rgb(239, 68, 68); color: white"
          >
            Disable
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
tr:hover {
  background-color: var(--color-bg-tertiary) !important;
}
</style>