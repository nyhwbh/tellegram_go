<template>
    <div class="work_board">
        <todoApp></todoApp>
        <div clclass="eoullim_calendar"></div>
    </div>
</template>

<script>
import todoApp from './co-work/TG-TodoApp.vue'
export default {
    components:{
        todoApp,
    }
}
</script>

<style>
.work_board{
    background-color: bisque;
    position: absolute;
    padding: 30px;
    top: 7%;
    left: 3%;
    height: 85%;
    width: 91%;
    border-radius: 30px;
}
</style>