<template>
  <div id="app">
    <div class="log_in" v-show="loggedIn">
      <LogInPage></LogInPage>
    </div>
    <div class="chatting_app">
      <ChattingApp></ChattingApp>
    </div>
  </div>
</template>

<script>
import LogInPage from './components/TG-LogInPage.vue'
import ChattingApp from './components/TG-ChattingPage.vue'
export default {
  components:{
    LogInPage,
    ChattingApp
  },
  data(){
    return{
      loggedIn: true
    }
  }


}
</script>

<style>

</style>