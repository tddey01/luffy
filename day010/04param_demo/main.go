package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//  参数相关实例

func queryStringHandlers(c *gin.Context) {
	// 获取query string参数
	// nameVal := c.Query("name") // 查不到 默认就是空字符串
	nameVal := c.DefaultQuery("name", "list") // 查不到 就用指定默认值 （第二个参数）
	cityVal := c.Query("ctiy")
	c.JSON(http.StatusOK, gin.H{
		"name": nameVal,
		"ctiy": cityVal,
	})
}

func formhandlers(c *gin.Context) {
	//  提取form 表单数据
	nameVal := c.PostForm("name")
	cityVal := c.DefaultPostForm("ctiy", "上海")
	c.JSON(http.StatusOK, gin.H{
		"name": nameVal,
		"city": cityVal,
	})
}

func paramshandlers(c *gin.Context) {
	//  提取路径参数
	actionVal := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"action": actionVal,
	})
}

func main() {
	router := gin.Default()
	// query_string https://www.baidu.com/search?name=list&citry=beijing
	router.GET("/query_string", queryStringHandlers)
	//  form params html 页面上form表单提交数据
	router.POST("/form", formhandlers)
	//  URL参数 /book/list /book/new/  book/delete
	// /post/2019/10  /post/2019/06
	
	router.GET("/book/:action", paramshandlers)
	router.Run()
}
