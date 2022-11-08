<template>
<router-view></router-view>
</template>

<script>
import {ref,reactive,toRefs} from 'vue'

const upload= ref()
import {get,postupload,post} from "@/util/service";

export default {
  setup() {
    const data = reactive({

    })





    return{
      ...toRefs(data),
    }
    //
  }
}



</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
