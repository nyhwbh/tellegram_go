<template>
  <div>
    <div class="log_in" v-show="loggedIn">
      <LogInPage></LogInPage>
    </div>
    <div class="chatting_app" v-show="!loggedIn">
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
.log_in{
  display: block;
  z-index: 1;
}
.chattingApp{
  display: block;
  z-index: 0;
}
</style>