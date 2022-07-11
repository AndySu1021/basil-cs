import { createApp } from 'vue'
import App from './App.vue'
import 'element-plus/dist/index.css'
import store from './store';

createApp(App).use(store).mount('#app')
