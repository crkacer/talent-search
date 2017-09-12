<template>
  <div class="users">
    <h1>Users</h1>
    <input type="text" v-model="newUser.name" placeholder="Enter Name">
    <br />
    <input type="text" v-model="newUser.email" placeholder="Enter Email">
    <br />
    <input type="submit" value="Submit">    
    <ul>
        <li v-for="user in users">
            {{user.name}} : {{user.email}}
        </li>
    </ul>
  </div>
</template>
<script>
export default {
  name: 'users',
  data () {
    return {
      newUser: {},
      users: [
        {
          name: 'John Doe',
          email: 'jdoe@gmail.com',
          contacted: false
        },
        {
          name: 'Jane Doe',
          email: 'janedoe@gmail.com',
          contacted: true
        },
        {
          name: 'Huy Dam',
          email: 'hdame@gmail.com',
          contacted: false
        }
      ]
    }
  },
  methods: {
    addUser: function (e) {
      this.users.push({
          name: this.newUser.name,
          email: this.newUser.email,
          contacted: false
      })
      console.log('add')
      e.preventDefault()
    }
  },
  computed: {
  }
}
</script>
<style scoped>

</style>