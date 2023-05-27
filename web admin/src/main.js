import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import SocketIO from './plugins/io'

const app = createApp(App)

app.use(SocketIO, {
  connection: 'ws://localhost/admin'
})

app.mount('#app')
