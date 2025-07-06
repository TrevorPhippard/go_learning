import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import './style.css' // Import global styles
import './assets/css/tailwind.css' // Import global styles

const pinia = createPinia()

createApp(App).use(router).use(pinia).mount('#app')
