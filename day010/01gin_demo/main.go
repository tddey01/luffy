package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "北京",
	})
}

func main() {
	// 启动一个默认的路由
	router := gin.Default()
	// 给/hello配置一个处理函数
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello 沙河！",
		})
	})
	router.GET("/index", indexHandler)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
