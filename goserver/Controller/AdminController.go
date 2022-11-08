package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"goserver/Service"
	"net/http"
	"strconv"
)

// 为了安全  请求后端连接一次数据  不是一次连接数据库 多次执行sql语句
var success = "success"
var fail = "fail"

func Adminexits(r *gin.Engine) {
	r.GET("/login", func(ctx *gin.Context) {
		fmt.Println("开始执行")
		id, _ := strconv.Atoi(ctx.Query("id")) //管理编号
		password := ctx.Query("password")      //管理员姓名
		//password是指针类型
		fmt.Println("获取账号和密码为", id, password)
		padmin := Service.Selectalladmin()
		for i := 0; i < len(*padmin); i++ {
			fmt.Println((*padmin)[i].Id)
			fmt.Println((*padmin)[i].Password)
			if (*padmin)[i].Id == id && (*padmin)[i].Password == password {
				ctx.JSON(http.StatusOK, success)
				return
			}
		}
		ctx.JSON(http.StatusOK, fail)
		return
	})
}
