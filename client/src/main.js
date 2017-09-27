<<<<<<< HEAD
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
=======
// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import VueRouter from 'vue-router'
import vueResource from 'vue-resource'
import App from './App'
import Users from './components/Users'
import Test from './components/Test'

Vue.use(vueResource)
Vue.use(VueRouter)
const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [
    {path: '/', component: Users},
    {path: '/test', component: Test}
  ]
})
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  router,
  template: `
  <div id="app">
    <ul>
      <li><router-link to="/">Users</router-link></li>
      <li><router-link to="/test">Test</router-link></li>
    </ul>
    <router-view></router-view>
  <div>
  `
}).$mount('#app')
>>>>>>> 34c575aa9cdab914bb59a357fbf5d4356afa0c67
