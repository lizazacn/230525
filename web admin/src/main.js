import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import SocketIO from './plugins/io'

app.use(SocketIO, {
  connection: 'ws://localhost/admin'
})

createApp(App).mount('#app')
