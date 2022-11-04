<template>
  <div class="chatting">
    <div class="chatting_list">
      <div class="chatting_menu" v-show="!menuOn">
        <section class="chatting_menu_search">
          <i class="fa-solid fa-bars"></i>
        </section>
        <section class="chatting_menu_search">
          <input class="chatting_list_search " type="text" placeholder="검색">
        </section>
      </div>
      <div class="chatting_list_friends">
        <ul>
          <li>이영훈</li>
          <li>한연수</li>
          <li>장주영</li>
        </ul>
      </div>
    </div>
    <div class="chatting_list_menu" v-show="menuOn">
    </div>
    <div class="chatting_room">
      <div class="chatting_room_none">
        <section>대화할 방을 선택하세요.</section>
      </div>
      <div class="chatting_room_single">

      </div>
      <div class="chatting_room_multi">

      </div>
    </div>

  </div>
</template>

<script>
export default {
  data() {
    return {
      menuOn : false
    }
  }
}
</script>

<style>
.chatting{
  display: flex;
  flex-direction: row;
  margin: 0;
  width: 100vw;
  height: 100vh;
}

.chatting_list{
  display: flex;
  flex-direction: column;
  background-color: #102024;
  width: 35%;
}

.chatting_menu{
  display: flex;
  flex-direction: row;
  padding: 10px 0px;
}

.chatting_menu:nth-child(1) {flex-grow: 0;}
.chatting_menu:nth-child(2) {flex-grow: 1;}

.chatting_menu_search{
  color: gray;
  font-size: 20px;
  margin-left: 20px;
}

.chatting_menu_search > input{
  flex-shrink: 0;
  margin-right: 10px;
  border-radius: 3px;
  height: 25px;
  background-color: grey;
  border: 0px solid grey;
  color:#0a1416
}

.chatting_list_search{
  width: 100%;
}

.chatting_room{
  background-color: #0a1416;
  width: 65%;
}

.chatting_room_none{
  display: flex;
  justify-content: center;
  align-items: center;
}

.chatting_room_none > section {
  display: block;
  width: 180px;
  height: 25px;
  padding : 3px 10px;
  color : white;
  background-color: #112d34;
  border-radius: 25px;
}
</style>