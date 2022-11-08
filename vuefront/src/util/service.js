import axios from "axios"
import {ElLoading, ElMessage} from 'element-plus'
import {config} from "@vue/cli-plugin-eslint/eslintOptions";

let loadingObj=null
//使用axios创作一个实例
const Service =axios.create({
    //下面属性是公共属性
    timeout:8000,//超时设置9秒
    baseURL:"http://localhost:8019",//controller前缀
 /*   headers不设置代表可以上传任何类型*/
/*    headers://请求类型
        {
        "Content-type":"application/json;chartset=utf-8"*/
         /*   "Content-Type":"*!/!*;chartset=utf-8"*/
          /*   "Content-Type":"multipart/form-data"*/
       /*     "Content-Type":"multipart/form-data"*/
      /*  }*/
})



// //请求拦截 -增加loading 对请求统一处理
Service.interceptors.request.use(config=>{
    loadingObj=ElLoading.service({
        lock:true,
        text:'Loading',
        background:'rgba(0,0,0,0.7)',
    })
    return config
})




//响应拦截 对返回值做统一处理
Service.interceptors.response.use(response=> {
loadingObj.close()
    // data是获取请求额数据
    const data=response.data
    console.log("请求数据为",data)
    console.log("请求数据格式为",response)
    if(data===null){//如果请求返回的数值为空
       ElMessage.error("服务器请求错误信息")
        return data
    }
return data
},error=>{
  loadingObj.close()
    ElMessage({
        message:"服务器错误",
        type:"error",
        duration:2000
    })
})

//config是多个参数
export const post=(url,config)=>{
    return Service(
        {
            url:url,
            ...config,
            method:"post",
            data:config,
            params:config,

        }
    )
}

export const get=(url,config)=>{
    return Service(
        {   url:url,
            ...config,
            method:"get",
            data:config,
            params:config,//参数名称要与getmapping里面的接受参数名称一样

        }
    )
}



//文件类型 可以将字符串参数上传
export  const postupload=(url,config)=>
{
    return Service(
        {   url:url,
            ...config,
            method:"post",
            data:config,
            params:config,//参数名称要与getmapping里面的接受参数名称一样
            headers://请求类型
                {
                  "Content-Type":"multipart/form-data"//这个是针对图片和文件的的
                    /*     "Content-Type":"multipart/form-data"*/
                }
        }
    )
}

//获取文件  返回blog类型
export  const postfile=(url,config)=>
{
    return Service(
        {   url:url,
            ...config,
            method:"post",
            data:config,
            params:config,//参数名称要与getmapping里面的接受参数名称一样
            responseType:'blob',
            headers://请求类型
                {
                    "Content-type": "application/json;chartset=utf-8"
                }
        }
    )
}

//获取文件  返回blog类型
export  const getfile=(url,config)=>
{
    return Service(
        {   url:url,
            ...config,
            method:"get",
            data:config,
            params:config,//参数名称要与getmapping里面的接受参数名称一样
            responseType:'blob',
            headers://请求类型
                {
                    "Content-type": "application/json;chartset=utf-8"
                }
        }
    )
}
