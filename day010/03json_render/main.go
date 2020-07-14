package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexhandlers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "这是index界面",
	})
}

func hellohandlers(c *gin.Context) {
	type userinfo struct {
		Name     string `json:"name"`
		Passwrod string `json:"pwd"`
	}

	u1 := userinfo{
		Name:     "hello",
		Passwrod: "123",
	}
	c.JSON(http.StatusOK, u1)
}

func xmlhandlers(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{
		"msg": "xml",
	})
}

func main() {
	r := gin.Default()
	r.GET("/index", indexhandlers)
	r.GET("/hello", hellohandlers)
	r.GET("/xml", xmlhandlers)
	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"massge": "ok", "status": http.StatusOK,
		})
	})
	r.Run()
}
