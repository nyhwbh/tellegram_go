<template>
  <div>
    <LogInPage v-show="loggedIn"></LogInPage>
    <ChattingApp v-show="!loggedIn"></ChattingApp>
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
body{
  margin: 0;
}
</style>