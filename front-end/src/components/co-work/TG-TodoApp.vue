<template>
  <div class="todoApp">
    <div class="board">
      <ListWrapper></ListWrapper>
      <div class="addListWrapper">
        <AddListController></AddListController>
      </div>
    </div>

  </div>
</template>

<script>
import AddListController from './todo/TG-AddListController'
import ListWrapper from './todo/TG-ListWrapper'
export default {
  components:{
    AddListController,
    ListWrapper
  }
}
</script>

<style>
.todoAPP{
  margin: 0;
  width: 100%;
  height: 100%;
}
.board{
    display: flex;
    flex-direction: row;
    overflow-x: auto;
}
.addListWrapper {
  flex-shrink: 0;
  width: 272px;
}

a {
  cursor: pointer;
}

textarea:focus {
  outline: none;
}
</style>