<template>
  <v-app>
    <main>
      <v-parallax src="/static/doc-images/background.jpg" style="height:100vh; margin-top: 0">
        <v-layout align-center justify-center>
          <form v-on:submit.prevent="">
            <input type="text" v-model="text" placeholder="Enter any names or categories" class="white" style="color:black;height:50px;font-size:150%;" size="50">
            <input type="image" src="/static/icon-images/magnifying-glass-icon.png" style="margin-left: 10px; margin-right: 10px; height:32px;width:32px;">
          </form>
        </v-layout>
      </v-parallax>
    </main>
    <v-footer>
    </v-footer>
  </v-app>
</template>

<script>
import Test from './Test'
import Users from './Users'
import headertab from './HeaderTab'
export default {
  name: 'homepage',
  data () {
    return {
      text: ''
    }
  },
  components: {
    Test, Users, headertab
  }
}
</script>
