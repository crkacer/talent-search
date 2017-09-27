<template>
  <v-toolbar class="cyan" style="margin-bottom: 0;">
    <v-toolbar-title>Talent Search</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-toolbar-items class="hidden-sm-and-down" v-for="button in buttons" :key="button.text">
      <v-btn flat :href=button.url class="button">{{ button.text }}</v-btn>
    </v-toolbar-items>
    <v-toolbar-items class="hidden-sm-and-down" v-for="dropdown in dropdowns" :key="dropdown.text">
      <v-menu open-on-hover offset-y>
      <v-btn flat slot="activator">{{ dropdown.text }}</v-btn>
      <v-list class="deep-orange lighten-4">
        <v-list-tile v-for="content in contents" :key="content.title" @click="">
          <v-list-tile-title>{{ content.title }}</v-list-tile-title>
        </v-list-tile>
      </v-list>
    </v-menu>
    </v-toolbar-items>
  </v-toolbar>
</template>
<script>
export default {
  name: 'headertab',
  data () {
    return {
      buttons: [
        {
          text: 'Home',
          url: '/'
        },
        {
          text: 'Register',
          url: '/users'
        }],
      dropdowns: [
        {
          text: 'Category'
        }],
      contents: [
        {
          title: 'Singer'
        },
        {
          title: 'Dancer'
        }]
    }
  }
}
</script>

<style scoped>
.button{
  color:black;
}
</style>