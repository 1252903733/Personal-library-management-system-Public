<template>

<div>

  <embed :src="url" style="width: 100%;height: 100%"/>

</div>
</template>


<script>
import {ref,reactive,toRefs} from 'vue'

const upload= ref()
import {get,postupload,post} from "@/util/service";

export default {
  setup() {
    const data = reactive({
      url:'C:\\Users\\Administrator\\Desktop\\论文模板.pdf',

    })





    return{
      ...toRefs(data),
    }
    //
  }
}



</script>