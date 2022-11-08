import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from "element-plus";//ElementPlus不能加括号
import 'element-plus/dist/index.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap'
import router from './router/index'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'//导入图标样式
const app=createApp(App)
//element icon注册
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.use(ElementPlus)
app.use(router)
app.mount('#app')
