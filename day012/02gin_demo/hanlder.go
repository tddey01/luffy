package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginsession "github.com/tddey01/luffy/day012/gin-session"
	"net/http"
)

// UserInfo 用户结构
type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// 校验用户中间件
// AuthMiddleware 其实就是上线文中取出session data 从session data取到islogin
func AuthMiddleware(c *gin.Context) {
	// 1. 从上下文中取到session data
	// 1. 先从上下文中获取session data
	fmt.Println("in Auth")
	tmpSD, _ := c.Get(ginsession.SessionContextName)
	sd := tmpSD.(*ginsession.SessionData)
	// 从session data取到islogin
	fmt.Printf("%#v\n", sd)
	value, err := sd.Get("isLogin")
	if err != nil {
		fmt.Println(err)
		//  取不到没有登录
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
				"err": "用户名或者密码不能为空",
			})
			return
		}
		if u.Username == "kn" && u.Password == "123" {
			//  登录成功 在当前页面这个用户的session data 保存一个简直对， isLogin=true
			// 1 先要从上下文中获取seesion data
			tmpSD, ok := c.Get(ginsession.SessionContextName)
			if !ok {
				panic("session middleware 中间件")
			}
			sd := tmpSD.(*ginsession.SessionData)
			// 2 给session data设置islogin=true
			sd.Set("isLogin", true)
			// 跳转到index界面
			c.Redirect(http.StatusMovedPermanently, toPath)
		} else {
			// 密码错误
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "用户名或者密码错误",
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

func viphandlers(c *gin.Context) {
	c.HTML(http.StatusOK, "vip.html", nil)
}


