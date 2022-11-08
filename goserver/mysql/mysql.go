package mysql

import (
	"database/sql"
	"fmt"
)

var DriverName = "mysql"  //数据库类型
var Host = "localhost"    //主机
var Port = "3306"         //端口
var Database = "goandvue" //数据库名称
var Username = "root"     //用户名称
var Password = "123456"   //密码
var Charset = "utf8"      //编码方式  记住不能utf-8

func Mysqlopen() *sql.DB {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		Username,
		Password,
		Host,
		Port,
		Database,
		Charset)
	/*	sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/information?charset=utf8")
	 */
	db, err := sql.Open(DriverName, args)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
