package Service

import (
	"fmt"
	"goserver/entity"
	"goserver/mysql"
)

func Selectalladmin() *[]entity.Admin {
	db := mysql.Mysqlopen() //获取数据库连接
	rows, err := db.Query("SELECT * FROM admin ")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	padmin := []entity.Admin{}
	for rows.Next() {
		var pamin entity.Admin
		rows.Scan(&pamin.Id, &pamin.Password)
		padmin = append(padmin, pamin)
	}
	rows.Close() //关闭指针
	db.Close()
	return &padmin //返回数组
}
