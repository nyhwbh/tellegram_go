<template>
  <div class="eoullim">
    <div class="eoullim_login" v-show="!loggedIn">
      <LogInPage></LogInPage>
    </div>
    <div class="eoullim_app" v-show="loggedIn">
      <WorkComponent></WorkComponent>
      <div class="eoullim_chatting" v-show="chattingOn">
        <ChattingApp></ChattingApp>
      </div>
      <div class="eoullin_chatting_off" v-show="!chattingOn">
        
      </div>
    </div>
  </div>
</template>

<script>
import LogInPage from './components/TG-LogInPage.vue'
import ChattingApp from './components/TG-ChattingPage.vue'
import WorkComponent from './components/TG-WorkComponent.vue'
export default {
  components:{
    LogInPage,
    ChattingApp,
    WorkComponent
  },
  data(){
    return{
      loggedIn: true,
      chattingOn : false,
    }
  }


}
</script>

<style>
body{
  margin: 0;
}
.eoullim_app{

}
.eoullim_chatting{
  position: absolute;
  right: 0;
}
</style>