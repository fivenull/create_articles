import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './directives'
import './components'
import VueSweetalert2 from './plugins/vue-sweetalert2'
// 引入插件
import Message from './plugins/message'

import Axios from 'axios'
Vue.prototype.REQUEST = Axios
Vue.prototype.HOST = 'http://127.0.0.1:517/api/v1/'

import api from './api/index'
Vue.prototype.API = api

Vue.use(VueSweetalert2)
Vue.use(Message)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
