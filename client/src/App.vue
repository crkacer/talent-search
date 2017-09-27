<<<<<<< HEAD
<template>
  <div class = "app">
    <headertab></headertab>
    <router-view></router-view>
  </div>
</template>

<script>
import headertab from './components/HeaderTab'
import homepage from './components/HomePage'
export default {
  name: 'app',
  components: {
    headertab, homepage
  }
}
</script>
<style>

</style>
=======
<template>
  <div id="app">
    <users></users>
  </div>
</template>

<script>
import Test from './components/Test'
import Users from './components/Users'
export default {
  name: 'app',
  components: {
    Test, Users
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
>>>>>>> 34c575aa9cdab914bb59a357fbf5d4356afa0c67
