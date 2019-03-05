import Vue from 'vue'
import App from './App.vue'
import router from './router/index'
import store from './plugins/store'
import inject from './plugins/inject'
Vue.use(inject)

import './plugins/iview.js'

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
