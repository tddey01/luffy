package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin demo
func main() {
	router := gin.Default()    // 得到一个默认的处理引擎
	router.LoadHTMLGlob("./*") // 加载HTML文件
	router.GET("/index", indexHandler)
	//  所有请求UTL以V1开头的交给
	v1Group := router.Group("/v1")
	{
		v1Group.GET("/index", indexHandler)
	}
	//  登录页面
	// router.GET("/login", loginGetHandler)
	// router.POST("/login", loginPosthandler)
	router.Any("/login", loginAnyhandler)

	// path 参数
	router.GET("/posts/:year/:month/:day", postsHandler)

	//  query striing
	router.GET("/seacrh", seacrhhandler)
	router.Run()
}

func indexHandler(c *gin.Context) {
	// 具体的处理请求的业务逻辑
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "hello world",
	})
}

/*
func loginAnyhandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		// 处理用户提交过来的请求数据
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	} else {
		// 返回一个登录页面
		c.HTML(http.StatusOK, "login.html", nil)
	}
}
*/

func postsHandler(c *gin.Context) {
	//取到path参数
	year := c.Param("year")
	month := c.Param("month")
	day := c.Param("day")
	c.JSON(http.StatusOK, gin.H{
		"year":  year,
		"month": month,
		"day":   day,
	})
}

func seacrhhandler(c *gin.Context) {
	// query string 参数
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

type Userinfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func loginAnyhandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		// 处理用户提交过来的请求数据
		var u Userinfo          // 声明一个Userinfo类型变量
		err := c.ShouldBind(&u) // ShouldBind:根据请求中的Content-Type
		if err != nil {
			// 解析数据出问题了
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"username": u.Username,
			"password": u.Password,
		})
	} else {
		// 返回一个登录页面
		c.HTML(http.StatusOK, "login.html", nil)
	}
}
