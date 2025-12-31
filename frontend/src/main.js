import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from '@/router'
import { createPinia } from 'pinia'
import { useAuthStore } from './stores/auth'

const app = createApp(App)

app.use(createPinia())  
app.use(router)         
const authStore = useAuthStore()
authStore.loadFromStorage()
app.mount('#app')

