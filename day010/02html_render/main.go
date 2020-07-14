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
	r.GET("/login", loghandlers)
	r.GET("/index", indexhandlers)
	r.Run()
}
