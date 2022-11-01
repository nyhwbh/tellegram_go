<template>
  <form class="logInBox">
    <div class="logInIDPW">
      <input type="text">
      <input type="password">
      <input type="button">
    </div>
  </form>
</template>

<script>
export default {

}
</script>

<style>
.logInBox{
  width: 280px;
  height: 560px;
  background-color: aquamarine;
}
.logInIDPW{
  display: flex;
  flex-direction: column;
  padding: 20px
}

input{
  border: 0px solid black;
  height: 15px;
  margin-bottom: 10px;
  padding: 10px;
}
</style>