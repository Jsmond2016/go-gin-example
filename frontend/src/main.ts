import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './assets/css/main.css'
import axios from 'axios'

// Configure axios defaults
axios.defaults.headers.common['Content-Type'] = 'application/json'

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.mount('#app')
