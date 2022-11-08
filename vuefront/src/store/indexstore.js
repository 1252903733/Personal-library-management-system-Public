import {createStore} from 'vuex'
//需要解决刷新不会初始化
//每次刷新也会初始化
//可以用在路由守卫哪里  每次url浏览state都会初始化这样UserRole就会为空
const store=createStore({
    //全局状态初始数值
    state: {
        UserRole: localStorage.getItem('userrole'),
        CurrentStatusInformation:localStorage.getItem('currentstatusinformation'),
        Id:localStorage.getItem('id'),
    },
    //计算state 获取对应的数值
    getters:{
    UserRoleStatus(state)
    {
        return state.UserRole
    }
    },
    //更新状态的方法 更新state的唯一方法 commit mutation
    mutations:{
        setUserRole(state,parameter)
        {
            state.UserRole=parameter
        }
    },
    //进行异步操作 可以返回promise 更改数据传递到mutation是里面去更改
    actions:{
        setUserRolePromise(context,parameter) {
        return  new Promise((resolve,reject)=>{
          context.commit("setUserRole",parameter)
        })

    }
},
    //数据比较多时分模块
    modules:{}
})
export  default store