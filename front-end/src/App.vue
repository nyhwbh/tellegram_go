<template>
  <div class="eoullim">
    <div class="eoullim_login">
      <LogInPage v-show="loggedIn"></LogInPage>
    </div>
    <div class="eoullim_todo"></div>
    <div clclass="eoullim_calendar"></div>
    <div class="eoullim_chatting">  
      <ChattingApp v-show="!loggedIn"></ChattingApp>
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
      loggedIn: false
    }
  }


}
</script>

<style>
body{
  margin: 0;
}
</style>