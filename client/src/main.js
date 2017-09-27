// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import router from './router'
import vueResource from 'vue-resource'
import Vuetify from 'vuetify'
import App from './App'
import 'vuetify/dist/vuetify.min.css'
import 'vue-awesome/icons'
import Icon from 'vue-awesome/components/Icon'
import 'vue-material-design-icons/styles.css'
import MenuIcon from 'vue-material-design-icons/magnify.vue'

Vue.component('magnify-icon', MenuIcon)
Vue.component('icon', Icon)
Vue.use(vueResource)
Vue.use(Vuetify)
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
