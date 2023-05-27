import { createApp } from 'vue'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import App from './App.vue'
import router from './router'
import 'element-plus/dist/index.css'
import SocketIO from './plugins/socketIo'

const app = createApp(App)
app.use(router)

app.use(SocketIO, {
  connection: '/admin'
})

// 注册elementplus图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.mount('#main')
