
<script setup>
import { ref, onMounted } from 'vue'
import api from '@/services/api'

const users = ref([])
const showCreateForm = ref(false)
const editingUser = ref(null)
const form = ref({ email: '', name: '', password: '' })

async function loadUsers() {
  const res = await api.get('/users')
  users.value = res.data
}

async function saveUser() {
  if (editingUser.value) {
    await api.put(`/users/${editingUser.value.id}`, form.value)
  } else {
    await api.post('/users', form.value)
  }
  closeForm()
  loadUsers()
}

async function deleteUser(id) {
  if (confirm('Xóa thật không?')) {
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
    <h2>Quản lý Người dùng</h2>
    <button @click="showCreateForm = true" class="btn btn-success mb-3">Thêm người dùng</button>

    <table class="table table-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Email</th>
          <th>Tên</th>
          <th>Hành động</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td>{{ user.id }}</td>
          <td>{{ user.email }}</td>
          <td>{{ user.name }}</td>
          <td>
            <button @click="editUser(user)" class="btn btn-warning btn-sm">Sửa</button>
            <button @click="deleteUser(user.id)" class="btn btn-danger btn-sm ms-2">Xóa</button>
          </td>
        </tr>
      </tbody>
    </table>

    <!-- Form thêm/sửa (bạn có thể tách ra component riêng sau) -->
    <div v-if="showCreateForm || editingUser" class="modal fade show d-block" style="background: rgba(0,0,0,0.5)">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5>{{ editingUser ? 'Sửa' : 'Thêm' }} người dùng</h5>
          </div>
          <div class="modal-body">
            <div class="mb-3">
              <label>Email</label>
              <input v-model="form.email" class="form-control" />
            </div>
            <div class="mb-3">
              <label>Tên</label>
              <input v-model="form.name" class="form-control" />
            </div>
            <div class="mb-3" v-if="!editingUser">
              <label>Mật khẩu</label>
              <input v-model="form.password" type="password" class="form-control" />
            </div>
          </div>
          <div class="modal-footer">
            <button @click="closeForm" class="btn btn-secondary">Hủy</button>
            <button @click="saveUser" class="btn btn-primary">Lưu</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
