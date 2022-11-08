package Service

import (
	"fmt"
	"goserver/entity"
	"goserver/mysql"
)

func Selectallbook() *[]entity.Book {
	db := mysql.Mysqlopen() //获取数据库连接
	rows, err := db.Query("SELECT * FROM book ")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	pbook := []entity.Book{}
	for rows.Next() {
		var pboo entity.Book
		rows.Scan(&pboo.Id, &pboo.Name, &pboo.Author, &pboo.Content, &pboo.Src, &pboo.DocSrc)
		pbook = append(pbook, pboo)
	}
	rows.Close() //关闭指针
	db.Close()
	return &pbook //返回数组
}

func AddBook(pbook *entity.Book) bool {
	db := mysql.Mysqlopen()
	stmt, err := db.Prepare("INSERT into book(id,name,author,content,src,docsrc) values (?,?,?,?,?,?)")
	if err != nil {
		return false
	}
	rs, err := stmt.Exec(pbook.Id, pbook.Name, pbook.Author, pbook.Content, pbook.Src, pbook.DocSrc)
	if err != nil {
		return false
	}
	fmt.Println(rs)
	if err != nil {
		return false
	}
	defer stmt.Close()
	db.Close()
	return true
}

func ReviseBook(pbook *entity.Book) bool {
	db := mysql.Mysqlopen()
	stmt, err := db.Prepare("update book set name=?,author=?,content=?,src=?,docsrc=? where id=?")
	if err != nil {
		return false
	}
	rs, err := stmt.Exec(pbook.Name, pbook.Author, pbook.Content, pbook.Src, pbook.DocSrc, pbook.Id)
	if err != nil {
		return false
	}
	fmt.Println(rs)
	if err != nil {
		return false
	}
	defer stmt.Close()
	db.Close()
	return true
}

func DelBook(id *int) bool {
	db := mysql.Mysqlopen()
	stmt, err := db.Prepare("delete  from book where id=?")
	if err != nil {
		return false
	}
	rs, err := stmt.Exec(id)
	if err != nil {
		return false
	}
	fmt.Println(rs)
	if err != nil {
		return false
	}
	defer stmt.Close()
	db.Close()
	return true
}
