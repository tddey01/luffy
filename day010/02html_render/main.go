package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loghandlers(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "北京",
	})
}

func indexhandlers(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "hello index",
	})
}
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	// 设置静态文件的目录
	// 第一个参数是代码里使用的路径，第二个参数是实际保存静态文件的路径
	r.Static("/dsb", "./statics")
	r.GET("/login", loghandlers)
	r.GET("/index", indexhandlers)
	r.Run()
}
