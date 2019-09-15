<template>
  <div id="app">
    <el-header style="padding:0px"><Navbar v-if="!$store.state.login"></Navbar></el-header>
    <router-view/>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar'
import NavbarAdmin from '@/components/NavbarAdmin'
import store from "@/store";
export default {
  name: 'App',
  data(){
    return{
      showNav:false
    };
  },
  store,
  components: {
    Navbar,
    NavbarAdmin,
  },
}
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height:100%;
}
#nav {
  padding: 30px;
  a {
    font-weight: bold;
    color: #2c3e50;
    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
