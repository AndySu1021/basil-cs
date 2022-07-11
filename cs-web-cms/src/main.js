import Vue from 'vue'

// A modern alternative to CSS resets
import 'normalize.css/normalize.css'

// Element ui usage
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import locale from 'element-ui/lib/locale/lang/zh-TW'

Vue.use(ElementUI, { locale })
Vue.use(ElementUI, {
  i18n: (key, value) => i18n.t(key, value)
})

// localization usage
import i18n from '@/lang'

// global css config
import '@/styles/index.scss'

import App from './App'
import store from './store'
import router from './router'

import '@/icons' // icon
import '@/filter' // Vue global filters
import '@/permission' // permission control
import '@/directive' // Vue global directive
import '@/utils/global' // Vue global function

window.i18n = i18n
Vue.config.productionTip = false

new Vue({
  el: '#app',
  i18n,
  router,
  store,
  render: h => h(App)
})
