package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func shopIndex(c *gin.Context) {
	fmt.Println(c.MustGet("key"))
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
