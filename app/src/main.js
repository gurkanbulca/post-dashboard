import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import {store} from "./store.js"
import {router} from "./router"
import vueCookies from "vue-cookies"

Vue.config.productionTip = false

Vue.use(vueCookies)




new Vue({
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')



