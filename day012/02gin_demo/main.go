package main

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/tddey01/luffy/day012/gin-session"
	"net/http"
)

// 测试 session服务 gin demo
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	// session 中间件服务器 为一个全局的中间件
	// 初始化全部MgrObj对象
	ginsession.IntMgr()
	r.Use(ginsession.SessionMiddleware(ginsession.MgrObj))
	r.Any("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.GET("/home", homeHandler)
	r.GET("/vip",AuthMiddleware, viphandlers)

	// 没有匹配的路由都走这个
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
	r.Run()
}
