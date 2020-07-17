package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BookMagesMentSystem
func main() {
	gin.SetMode(gin.DebugMode)
	//  程序启动 就连接数据库
	err := initDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("template/**/*")
	//  查看所有书籍数据
	r.GET("/book/list", bookListHandler)
	// 返回一个页面给用户填写新增的书籍信息
	r.GET("/book/new", newBookhandler)
	r.POST("/book/new", createnewBookhandler)
	r.GET("/book/detele", deleteBookhandler)
	r.Run()
}

func bookListHandler(c *gin.Context) {
	bookList, err := queryAllBook()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	// 返回数据
	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 0,
	// 	"data": bookList,
	// })
	c.HTML(http.StatusOK, "book/book_list.tmpl", gin.H{
		"code": 0,
		"data": bookList,
	})
}

// 添加图书
func newBookhandler(c *gin.Context) {
	// 给用返回一个添加书籍的页面的处理函数
	c.HTML(http.StatusOK, "book/newbook.html", nil)
}

//  创建书籍的处理函数
func createnewBookhandler(c *gin.Context) {

	//  创建书籍的处理函数
	// 从form表单提取数据
	var msg string
	titleVal := c.PostForm("title")
	priceVal := c.PostForm("price")
	price, err := strconv.ParseFloat(priceVal, 64)
	if err != nil {
		msg = "无效价格参数"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	fmt.Printf("%T %T\n", titleVal, price)
	// 数据插入数据库
	err = insertBook(titleVal, price)
	if err != nil {
		msg = "插入数据失败 请重试"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

//  删除具体某本书
func deleteBookhandler(c *gin.Context) {
	//  提取query string参数
	idStr := c.Query("id")
	idVal, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	// 数据是个正经数字
	// 去数据库删除具体的记录
	fmt.Println(idStr)
	err = deleteBook(idVal)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}
