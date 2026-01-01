
<script setup>
import { ref, onMounted } from 'vue'
import api from '@/services/api'

const users = ref([])
const showCreateForm = ref(false)
const editingUser = ref(null)
const form = ref({ email: '', name: '', password: '' })
const error = ref('')

async function loadUsers() {
  const res = await api.get('/users')
  users.value = res.data
}

async function saveUser() {
  try { 
    if (editingUser.value) {
      await api.put(`/users/${editingUser.value.id}`, form.value)
    } else {
      await api.post('/users', form.value)
    }
    error.value = ""; 
    closeForm()
    loadUsers()
  }
  catch (err) { 
    if (err.response) {
      const backendMessage = err.response.data.message || 
                             err.response.data.error || 
                             JSON.stringify(err.response.data); 

      error.value = backendMessage; 
    } else {
      error.value = 'Error connecting to server';
    }

    console.error(err);
  }

}

async function deleteUser(id) {
  if (confirm('delete user ' + id + ' ?')) {
    await api.delete(`/users/${id}`)
    loadUsers()
  }
}

function editUser(user) {
  editingUser.value = user
  form.value = { email: user.email, name: user.name, password: '' }
}

function closeForm() {
  showCreateForm.value = false
  editingUser.value = null
  form.value = { email: '', name: '', password: '' }
}

onMounted(loadUsers)
</script>
<template>
  <div>
    <h2>Users Management</h2>
    <button @click="showCreateForm = true" class="btn btn-success mb-3">Add user</button>

    <table class="table table-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Email</th>
          <th>Username</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td>{{ user.id }}</td>
          <td>{{ user.email }}</td>
          <td>{{ user.username }}</td>
          <td>
            <button @click="editUser(user)" class="btn btn-warning btn-sm">Modify</button>
            <button @click="deleteUser(user.id)" class="btn btn-danger btn-sm ms-2">Remove</button>
          </td>
        </tr>
      </tbody>
    </table>

    <div v-if="showCreateForm || editingUser" class="modal fade show d-block" style="background: rgba(0,0,0,0.5)">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5>{{ editingUser ? 'Modify' : 'Remove' }} user</h5>
          </div>
          <div class="modal-body">
            <div class="mb-3">
              <label>Email</label>
              <input v-model="form.email" class="form-control" />
            </div>
            <div class="mb-3">
              <label>Username</label>
              <input v-model="form.username" class="form-control" />
            </div>
            <div class="mb-3" v-if="!editingUser">
              <label>Password</label>
              <input v-model="form.password" type="password" class="form-control" />
            </div>
            <div class="mb-3" v-if="!editingUser">
              <label>Confirm Password</label>
              <input v-model="form.confirm_password" type="password" class="form-control" />
            </div>
          </div>
          <div class="modal-footer">
            <button @click="closeForm" class="btn btn-secondary">Cancel</button>
            <button @click="saveUser" class="btn btn-primary">Save</button>
             <p v-if="error" class="text-danger mt-3">{{ error }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
