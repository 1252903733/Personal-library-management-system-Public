package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goserver/Service"
	"goserver/entity"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func Getbookdata(r *gin.Engine) {
	r.GET("/getbookdata", func(ctx *gin.Context) {
		fmt.Println("开始执行")
		//password是指针类型
		pbook := Service.Selectallbook()
		ctx.JSON(http.StatusOK, pbook)
		return
	})
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	fmt.Println("寻找文件中")
	src, err := file.Open()
	if err != nil {
		fmt.Println("文件寻找失败")
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		fmt.Println("文件寻找失败1")
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}

// 获取前端图片为你教案
func Getfile(r *gin.Engine) {
	r.GET("/getfile", func(ctx *gin.Context) {
		Id := ctx.Query("Id1")  //管理编号 不能使用Name不然接受不到
		Src := ctx.Query("Src") //管理员姓名*/
		filename := Id + "." + Src
		ctx.File("./img/" + filename)
	})
}

func GetDoc(r *gin.Engine) {
	r.GET("/getdoc", func(ctx *gin.Context) {
		Id := ctx.Query("Id1")        //管理编号 不能使用Name不然接受不到
		DocSrc := ctx.Query("DocSrc") //管理员姓名*/
		filename := Id + "." + DocSrc
		fmt.Println("获取文档", filename)
		ctx.File("./doc/" + filename)
	})
}

func GetDow(r *gin.Engine) {
	r.GET("/getdow", func(ctx *gin.Context) {
		Id := ctx.Query("Id1")        //管理编号 不能使用Name不然接受不到
		DocSrc := ctx.Query("DocSrc") //管理员姓名*/
		filename := Id + "." + DocSrc
		fmt.Println("获取文档", filename)
		ctx.File("./doc/" + filename)
	})
}
func ReviseBook(r *gin.Engine) {
	r.POST("/revise", func(ctx *gin.Context) {
		Id := ctx.PostForm("Id1")
		Name := ctx.PostForm("Name1")
		Content := ctx.PostForm("Content")
		Src := ctx.PostForm("Src")
		DocSrc := ctx.PostForm("DocSrc")
		Author := ctx.PostForm("Author")
		srcfile, _ := ctx.FormFile("srcfile")
		docsrcfile, _ := ctx.FormFile("docsrcfile")
		fmt.Println("编号添加" + Id)
		fmt.Println("名称添加" + Name)
		fmt.Println("内容添加" + Content)
		fmt.Println("作者添加1" + Author)
		fmt.Println(Src)
		fmt.Println(DocSrc)
		if srcfile == nil {
			fmt.Println("图片文件没有修改")
		} else {
			fmt.Println("生成图片")
			ctx.SaveUploadedFile(srcfile, "./img/"+Id+"."+Src)
		}
		if docsrcfile == nil {
			fmt.Println("文件不存在")
		} else {
			ctx.SaveUploadedFile(docsrcfile, "./doc/"+Id+"."+DocSrc)
		}
		pbook := entity.Book{}
		pbook.Id, _ = strconv.Atoi(Id)
		pbook.Name = Name
		pbook.Content = Content
		pbook.Src = Src
		pbook.DocSrc = DocSrc
		pbook.Author = Author
		is := Service.ReviseBook(&pbook)
		if is {
			ctx.String(http.StatusOK, success)
		} else {
			ctx.String(http.StatusOK, fail)
		}

	})
}

func AddBook(r *gin.Engine) {
	r.POST("/add", func(ctx *gin.Context) {
		Id := ctx.PostForm("Id1")
		Name := ctx.PostForm("Name1")
		Content := ctx.PostForm("Content")
		Src := ctx.PostForm("Src")
		DocSrc := ctx.PostForm("DocSrc")
		Author := ctx.PostForm("Author")
		srcfile, _ := ctx.FormFile("srcfile")
		docsrcfile, _ := ctx.FormFile("docsrcfile")
		fmt.Println("编号添加" + Id)
		fmt.Println("名称添加" + Name)
		fmt.Println("内容添加" + Content)
		fmt.Println("作者添加" + Author)
		fmt.Println(Src)
		fmt.Println(DocSrc)
		fmt.Println(srcfile.Filename)
		fmt.Println(docsrcfile.Filename)
		ctx.SaveUploadedFile(srcfile, "./img/"+Name+"."+Src)
		ctx.SaveUploadedFile(docsrcfile, "./doc/"+Name+"."+DocSrc)
		pbook := entity.Book{}
		pbook.Id, _ = strconv.Atoi(Id)
		pbook.Name = Name
		pbook.Content = Content
		pbook.Src = Src
		pbook.DocSrc = DocSrc
		pbook.Author = Author
		is := Service.AddBook(&pbook)
		if is {
			ctx.String(http.StatusOK, success)
		} else {
			ctx.String(http.StatusOK, fail)
		}
	})
}

func DelBook(r *gin.Engine) {
	r.POST("/delbook", func(ctx *gin.Context) {
		Id, err := strconv.Atoi(ctx.PostForm("Id1"))
		if err != nil {
			fmt.Println(Id)
		}

		is := Service.DelBook(&Id)
		if is {
			ctx.String(http.StatusOK, success)
		} else {
			ctx.String(http.StatusOK, fail)
		}

	})

}
