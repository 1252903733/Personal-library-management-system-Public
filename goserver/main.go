package main //表明这个是main包

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"goserver/Config"
	"goserver/Controller"
	"goserver/entity"
)

// 定义一个结构体对象  属性首字母要大写 不然返回的json是一个空 但是会显示有多少个数据

type multuadmin struct {
	Admin []entity.Admin /*`json:"admin"`*/
}

/*func bookdata(r *gin.Engine) {
	r.GET("/bookdata", func(ctx *gin.Context) {
		fmt.Println("正在请求图书信息")
	})
}
*/

func login(rows *sql.Rows, id *int, password *string) bool {
	var padmin entity.Admin
	for rows.Next() {
		rows.Scan(&padmin.Id, &padmin.Password)
		if padmin.Id == *id && padmin.Password == *password {
			return true
		}
	}
	rows.Close()
	return false
}

func main() {
	r := gin.Default() //返回一个默认的路由引擎
	r.Use(Config.Core())
	fmt.Println("执行")
	Controller.Adminexits(r)
	Controller.Getbookdata(r)
	Controller.Getfile(r)
	Controller.GetDoc(r)
	Controller.GetDow(r)
	Controller.AddBook(r)
	Controller.ReviseBook(r)
	Controller.DelBook(r)
	r.Run(":8019")
	//1.get会返回json  但是post不会返回
	//2.fmt.Println("开始访问") 如果是post请求这个不会输出
	//3.ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "返回前端数据"}) 返回前端json 只有 get请求才能返回
	//4.使用http://localhost:8012/login?id=2&password=甘世涛可以把参数传递过去 这样就可以输出这个参数 同时后端也能获取这个参数
	/*5.
	这个是封装数组
	pmultuadmin := new(multuadmin)
		var pamin Admin
		for rows.Next() {
			rows.Scan(&pamin.Id, &pamin.Password)
			fmt.Println(&pamin.Id, &pamin.Password)
			pmultuadmin.Admin = append(pmultuadmin.Admin, Admin{pamin.Id, pamin.Password})
	}*/
	//6.if else  else要做{右边
	//7.rows.Close() 释放内存 一旦释放  rows的内容为空  不能再次遍历
	//8.gin发送json数据 会自动生成0 1  2 3 4ctx.JSON(http.StatusOK,)

	//获取管理是否存在

	/*
		r.GET("/bookdata", func(ctx *gin.Context) {
			fmt.Println("正在请求图书信息")
		})

		r.GET("/login", func(ctx *gin.Context) {
			fmt.Println("正在请求登录信息")
			db := mysqlopen()                      //获取数据库连接
			rows := admindata(db)                  //获取数据
			id, _ := strconv.Atoi(ctx.Query("id")) //管理编号
			password := ctx.Query("password")      //管理员姓名

			if login(rows, &id, &password) {
				ctx.String(http.StatusOK, success)
			} else {
				ctx.String(http.StatusOK, "shibaia")
			}
		})
		r.Run(":8012") //这个8012是端口号*/

}

/*func shishi() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/information?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}

	type info struct {
		id     int    `db:"id"`
		name   string `db:"name"`
		author string `db:"author"`
	}
	rows, err := db.Query("SELECT * FROM student")

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var s info
		err = rows.Scan(&s.id, &s.name, &s.author)
		fmt.Println(s)
	}
	rows.Close()
}*/

/*
// go获取端口号
func GetPort() (int, error) {
	address, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, nil
	}
	listen, err := net.ListenTCP("tcp", address)
	if err != nil {
		return 0, nil
	}
	defer listen.Close()

	return listen.Addr().(*net.TCPAddr).Port, nil
}

// 获取数据库信息
func databaseconnection() *gorm.DB {
	driverName := "mysql"     //数据库类型
	host := "localhost"       //主机
	port := "3306"            //端口
	database := "information" //数据库名称
	username := "root"        //用户名称
	password := "123456"      //密码
	charset := "utf-8"        //编码方式
	args := fmt.Sprintf("$s:%s@tcp(%s:%s)?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/information?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}

	type info struct {
		id     int    `db:"id"`
		name   string `db:"name"`
		author string `db:"author"`
	}
	rows, err := db.Query("SELECT * FROM student")

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var s info
		err = rows.Scan(&s.id, &s.name, &s.author)
		fmt.Println(s)
	}
	rows.Close()
}
*/
