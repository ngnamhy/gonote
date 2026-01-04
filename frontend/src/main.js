import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from '@/router'
import { createPinia } from 'pinia'
import { useAuthStore } from '@/features/auth/stores/authStore'
import PrimeVue from 'primevue/config';
import Aura from '@primeuix/themes/aura';
import 'primeicons/primeicons.css'

const app = createApp(App)

app.use(createPinia())  
app.use(router)   

app.use(PrimeVue, {
    theme: {
        preset: Aura
    }
});     


const authStore = useAuthStore()
authStore.loadFromStorage()
app.mount('#app')

