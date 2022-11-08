<template>
  <el-form ref="formRef"  :model="admin"  class="demo-dynamic" >
    <!--prop要跟属性名称对应-->
    <el-form-item prop="id"    :rules="[ { required: true, message: '账号不能为空', trigger: 'blur',}, ]">
      <el-input  v-model="admin.id" ><template #prepend>账号</template></el-input>
    </el-form-item>
    <!--密码输入-->
    <el-form-item prop="password"  class="password-input"  :rules="[ { required: true, message: '密码不能为空', trigger: 'blur',}, ]">
      <el-input type="password" v-model="admin.password" ><template #prepend>密码</template></el-input>
    </el-form-item>
    <el-button type="primary" @click='lgintrigger' class="login-btn">登录</el-button>
  </el-form>
</template>


<script>
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {get,post} from "@/util/service";
import { reactive, onMounted, ref, toRefs } from 'vue'
import { ElMessage} from 'element-plus'

export default {
  setup() {
    const router=useRouter()
    const data = reactive({
     admin:{
      id:"",
       password:""
     },
      page:{
        currentpage:1,//当前页面
        pagesize:5,//每页的数据个数
        total:10//数据总数
      },
    })


    onMounted(async () => {//初始化函数
    })


    const lgintrigger=async () => {
      let res1
      await get("/login", {"id":data.admin.id,"password":data.admin.password}).then(res => {
        res1 = res//获取学生信息 然后进行显示
      })
      if (res1==="success")
      {
        router.push("/home")
      }
      else {
        ElMessage.error("密码错误")
      }

    }

    return{
      lgintrigger,
      ...toRefs(data),

    }
    //
  }
}



</script>