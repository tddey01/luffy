package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// middievare
func main() {
	r := gin.Default()
	r.Use(castTime)
	// 根据URL分执行的函数路由
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndex)
		shoppingGroup.GET("/home", shopHome)
	}
	r.Run()
}

func castTime(c *gin.Context) {
	start := time.Now()
	c.Set("key", "小微那个还在")
	c.Next() // 运行下一个Handler函数
	//  统计耗时

	cast := time.Since(start)
	fmt.Println(cast)
}
