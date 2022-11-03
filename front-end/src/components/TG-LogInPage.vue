<template>
  <div class="login">
    <div class="login_wrap" v-show="signinOff" method="POST">
      <ul class="menu_wrap">
        <li class="menu_item" rule="tab" :class="[{ login_on: signinOff },{signin_on : !signinOff}]">로그인</li>
        <li class="menu_item" rule="tab" :class="[{ login_on: !signinOff }, { signin_on : signinOff}]">회원가입</li>
      </ul>
      <form class="login_idpw_wrap">
        <input class="login_form" type="text" placeholder="ID를 입력해 주세요" name="user_email" required>
        <input class="login_form" type="password" placeholder="비밀번호를 입력해 주세요" name="password" required>
        <input class="login_submit" type="submit" value="로그인">
      </form>
    </div>
    <div class="signin_wrap" v-show="!signinOff" method="POST">
      <ul class="menu_wrap">
        <li class="menu_item" rule="tab" :class="[{ login_on: signinOff },{signin_on : !signinOff}]">로그인</li>
        <li class="menu_item" rule="tab" :class="[{ login_on: !signinOff }, { signin_on : signinOff}]">회원가입</li>
      </ul>
      <form class="signin_from_wrap">
        <h3 class="signin_name">이름</h3>
        <input class="signin_form" type="text" placeholder="이름을 입력해 주세요" name="username" required>
        <h3 class="signin_email">이메일</h3>
        <input class="signin_form" type="text" placeholder="이메일을 입력해 주세요" name="user_email" required>
        <h3 class="signin_email">비밀번호</h3>
        <input class="signin_form" type="password" placeholder="영문자, 숫자, 특수문자 포함 최소 8-20자" name="password" required>
        <input class="signin_form" type="password" placeholder="비밀번호를 확인해 주세요" name="password_check" required>
        <input class="signin_submit" type="submit" value="회원가입">
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
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

.signin_wrap{
  background-color: white;
    border-radius: 5px;
    display: flex;
    flex-direction: column;
    width: 470px;
    height: 600px;
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

.login_idpw_wrap{
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px
}

.signin_from_wrap{
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

.signin_form{
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

.signin_submit{
  border: 0px solid black;
    border-radius: 5px;
    background: #4981f8;
    color: white;
    width: 100%;
    height: 60px;
    padding: 10px;
    margin-top: 20px;
}

.login_on{
  background-color: white;
  border-top-left-radius: 5px;
  border-top-right-radius: 5px;
  font-weight: 600;
}

.signin_on{
  background-color: whitesmoke;
  border-top-left-radius: 5px;
  border-top-right-radius: 5px;
  font-weight: 600;
}
</style>\

<style scoped>
h3{
  margin: 10px 0px;
  font-size: 17px;
}
</style>