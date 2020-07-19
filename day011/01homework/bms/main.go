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
	// //  查看所有书籍数据
	// r.GET("/book/list", bookListHandler)
	// // 返回一个页面给用户填写新增的书籍信息
	// r.GET("/book/new", newBookhandler)
	// r.POST("/book/new", createnewBookhandler)
	// r.GET("/book/detele", deleteBookhandler)
	// // day11
	// r.Any("/book/edit", editBookHandler)

	BookGroup := r.Group("/book")
	{
		BookGroup.GET("/list", bookListHandler)
		BookGroup.GET("/new", newBookhandler)
		BookGroup.POST("/new", createnewBookhandler)
		BookGroup.GET("/detele", deleteBookhandler)
		// day11
		BookGroup.Any("/edit", editBookHandler)

		//  组合查询
		BookGroup.GET("/detail/:id", getBookDetailHandler)
	}
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

// 编辑更新书籍
func editBookHandler(c *gin.Context) {
	// 1  获取到用户编辑的是书的那一本， 从querystring获取到需要编辑数据id值
	idStr := c.Query("id")
	if len(idStr) == 0 {
		// 请求中没有携带我要用的数据， 该请求是无效
		c.String(http.StatusBadRequest, "无效的请求", nil)
		return
	}
	//  HTTP请求传过来的参数 通常都是string类型， 根据需要自己需求转换成相应的数据类型
	bookID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// 请求中没有携带我要用的数据， 该请求是无效
		c.String(http.StatusBadRequest, "无效的请求", nil)
		return
	}
	if c.Request.Method == "POST" {
		// 1. 获取用户提交的数据
		titleVal := c.PostForm("title")
		priceStr := c.PostForm("price")
		priceVal, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "无效价格信息", nil)
			return
		}
		// 2. 去数据库更新对应的书籍数据
		// ? id去哪儿了？
		err = editBook(titleVal, priceVal, bookID)
		if err != nil {
			c.String(http.StatusInternalServerError, "更新数据失败")
			return
		}
		// 3. 跳转回/book/list页面查看是否修改成功
		// 相同网站跳转：可以写相对路径
		// 不一样的网站跳转：需要写绝对路径
		c.Redirect(http.StatusMovedPermanently, "/book/list")
	} else {
		// 需要通过给模板渲染上原来的旧数据
		// 1  获取到用户编辑的是书的那一本， 从querystring获取到需要编辑数据id值
		// idStr := c.Query("id")
		// if len(idStr) == 0 {
		// 	// 请求中没有携带我要用的数据， 该请求是无效
		// 	c.String(http.StatusBadRequest, "无效的请求", nil)
		// 	return
		// }
		// //  HTTP请求传过来的参数 通常都是string类型， 根据需要自己需求转换成相应的数据类型
		// bookID, err := strconv.ParseInt(idStr, 10, 64)
		// if err != nil {
		// 	// 请求中没有携带我要用的数据， 该请求是无效
		// 	c.String(http.StatusBadRequest, "无效的请求", nil)
		// 	return
		// }
		// 2 根据id取到书籍信息
		bookObj, err := queryBookByID(bookID)
		if err != nil {
			c.String(http.StatusBadRequest, "无效的书籍id", nil)
			return
		}

		// 3 把书籍数据渲染到页面上
		c.HTML(http.StatusOK, "book/book_edit.html", bookObj)
	}
}

// 组合查询
func getBookDetailHandler(c *gin.Context) {
	//  1 获取书籍id
	tmpbookID := c.Param("id") // string类型参数
	if len(tmpbookID) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"err": "请求数据参数有误 Param",
		})
		return
	}
	bookID, err := strconv.ParseInt(tmpbookID, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "请求数据参数有误 Param",
		})
		return
	}
	//  2 去数据库查询获得具体的书籍信息
	bookObj, err := queryBookByID(bookID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	//  3 返回JSON格式数据
	c.JSON(http.StatusOK, bookObj)

}
