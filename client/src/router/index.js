import Vue from 'vue'
import VueRouter from 'vue-router'
import Users from '../components/Users'
import Test from '../components/Test'
import HomePage from '../components/HomePage'

Vue.use(VueRouter)

export default new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [
    {path: '/', component: HomePage},
    {path: '/test', component: Test},
    {path: '/users', component: Users}
  ]
})
