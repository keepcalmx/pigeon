import { createApp } from 'vue'
import App from './App.vue'

const app = createApp(App)

// 状态管理
import { createPinia } from 'pinia'
app.use(createPinia())

// 路由管理
import router from './router'
app.use(router)

// font awesome
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { far } from '@fortawesome/free-regular-svg-icons'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { fab } from '@fortawesome/free-brands-svg-icons'

library.add(far, fas, fab)
app.component('font-awesome-icon', FontAwesomeIcon)

// vue-toastification
import Toast, { PluginOptions } from 'vue-toastification'
import 'vue-toastification/dist/index.css'

const toastOptions: PluginOptions = {
  // pass
}
app.use(Toast, toastOptions)

app.mount('#app')
