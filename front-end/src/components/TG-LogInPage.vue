<template>
  <div class="login">
    <div class="login_wrap" method="POST">
      <ul class="menu_wrap">
        <li class="menu_item" rule="tab" :class="[{ login_on : loginOn },{signin_on : !signinOff}]">로그인</li>
        <li class="menu_item" rule="tab" :class="[{ login_on: !loginOn }, { signin_on : signinOff}]">회원가입</li>
      </ul>
      <form class="login_idpw_wrap">
        <input class="login_form" type="text" placeholder="ID를 입력해 주세요" name="username" required>
        <input class="login_form" type="password" placeholder="비밀번호를 입력해 주세요" name="password" required>
        <input class="login_submit" type="submit" value="로그인">
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loginOn:true,
      signinOff : true
    }
  }
}
</script>

<style>
.login{
  background-color: #4981f8;
  margin:0;
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.login_wrap{
  background-color: white;
  border-radius: 5px;
  display: flex;
  flex-direction: column;
  width: 470px;
  height: 300px;
}

.menu_wrap{
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  list-style: none;
  margin: 0;
  padding: 0;
  background-color: whitesmoke;
  border-radius: 5px;
}

.menu_item{
  display: block;
  height: 30px;
  width: 50%;
  text-align: center;
  padding-top: 10px;
  
}
.login_on{
  background-color: white;
  border-top-left-radius: 5px;
  border-top-right-radius: 5px;
}

.signin_on{
  background-color: whitesmoke;
  border-top-left-radius: 5px;
  border-top-right-radius: 5px;
}

.login_idpw_wrap{
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px
}

.login_form{
  border: 1px solid gainsboro;
  border-radius: 5px;
  width: 95%;
  height: 40px;
  margin: 0 10px 10px 10px;
  padding: 10px;
}

.login_submit{
  border: 0px solid black;
  border-radius: 5px;
  background: #4981f8;
  color:white;
  width: 100%;
  height: 60px;
  padding: 10px;
  margin-top: 15px;
}
</style>