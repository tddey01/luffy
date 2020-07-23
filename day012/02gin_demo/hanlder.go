package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ginsession "github.com/tddey01/luffy/day012/gin-session"
)

// UserInfo 结构体 
type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// 编写一个校验用户是否登录的中间件
// 其实就是从上下文中取到session data,从session data取到isLogin

// AuthMiddleware 单独认正
func AuthMiddleware(c *gin.Context) {
	// 1. 从上下文中取到session data
	// 1. 先从上下文中获取session data
	fmt.Println("in Auth")
	tmpSD, _ := c.Get(ginsession.SessionContextName)
	sd := tmpSD.(ginsession.SessionData)
	// 2. 从session data取到isLogin
	fmt.Printf("%#v\n", sd)
	value, err := sd.Get("isLogin")
	if err != nil {
		fmt.Println(err)
		// 取不到就是没有登录
		c.Redirect(http.StatusFound, "/login")
		return
	}
	fmt.Println(value)
	isLogin, ok := value.(bool)
	if !ok {
		fmt.Println("!ok")
		c.Redirect(http.StatusFound, "/login")
		return
	}
	fmt.Println(isLogin)
	if !isLogin {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.Next()
}

func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		toPath := c.DefaultQuery("next", "/index")
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "用户名或密码不能为空",
			})
			return
		}
		if u.Username == "kn" && u.Password == "123" {
			// 登陆成功，在当前这个用户的session data 保存一个键值对：isLogin=true
			// 1. 先从上下文中获取session data
			tmpSD, ok := c.Get(ginsession.SessionContextName)
			if !ok {
				panic("session middleware")
			}
			sd := tmpSD.(ginsession.SessionData)
			// 2. 给session data设置isLogin = true
			sd.Set("isLogin", true)
			sd.Save()
			// 跳转到index页面
			c.Redirect(http.StatusMovedPermanently, toPath)
		} else {
			// 密码错误
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "用户名或密码错误",
			})
			return
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}

}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func vipHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "vip.html", nil)
}
