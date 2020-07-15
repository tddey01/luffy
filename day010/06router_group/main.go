package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//  路由分组
func main() {
	r := gin.Default()
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndex)
		shoppingGroup.GET("/home", shopHome)
	}

	liveGroup := r.Group("/live")
	{
		liveGroup.GET("/index", liveIndex)
		liveGroup.GET("/home", liveHome)
	}

	r.Run()
}

func shopIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "shopping/index",
	})
}

func shopHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "shopping/homex",
	})
}

func liveIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "live/Index",
	})
}
func liveHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "live/Home",
	})
}
