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
