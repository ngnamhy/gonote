import './app/styles/assets/main.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config'
import Aura from '@primeuix/themes/aura'
import 'primeicons/primeicons.css'

import App from './app/App.vue'
import router from './app/router'
import { setupApiInterceptors } from './app/config/api'

const app = createApp(App)

// Setup Pinia
const pinia = createPinia()
app.use(pinia)

// Setup API interceptors after pinia is ready
setupApiInterceptors()

// Setup Router
app.use(router)

// Setup PrimeVue
app.use(PrimeVue, {
  theme: {
    preset: Aura
  }
})

// Mount app
app.mount('#app')

