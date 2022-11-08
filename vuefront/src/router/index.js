import {createRouter, createWebHashHistory, createWebHistory} from 'vue-router'
import store from '../store/indexstore'
// 注意components与component的使用方法
const routers=[
        {
            path: "/",
            name: "login",//登录页面 首页即登录页面 映射地址
            component:()=>import('../components/Login.vue')//实际地址
        },
    {
        path: "/home",
        name: "home",//登录页面 首页即登录页面 映射地址
        component:()=>import('../components/Home.vue')//实际地址
    },
    {
        path: "/ceshi",
        name: "ceshi",//登录页面 首页即登录页面 映射地址
        component:()=>import('../components/Ceshi.vue')//实际地址
    },
]
/*mode:"hash",
    history: createWebHashHistory(process.env.BASE_URL),*/
const router=createRouter({
        history: createWebHashHistory(),
        routes: routers,

        /*
        model:'hash',
            routes:routers,
            base: process.env.BASE_URL,
        }
        */
    }
)

//路由守卫前置操作  跳转页面之前都会触发这个操作 主要判断是否已经登录  相当于拦截器
router.beforeEach((to, from, next) => {

      //如果进入的是登录界面 则允许你进入  之所以加return是因为每次next的时候会触发路由前置操作 每次next都会触发route.beforeEach这样子会无限循环  next是回调函数
    //每次执行next都会执行route.beforeEach
        console.log("store",store.state.UserRole)
       if(to.path=="/")
       {
           next()
           return
       }

         //如果没有登录 返回登录页面
/*        if(store.state.UserRole===null)
       {
           alert("未登录");
           next("/")
           return
       }*/
next()
}
)


export default router