package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserInfo 用户结构
type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// cookie 示例
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.Any("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.GET("/home", homeHandler)
	r.GET("/vip", cookieMiddeware, viphandlers)
	r.Run()
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
			//  登录成功
			c.SetCookie("username", u.Username, 20, "/", "127.0.0.1", false, true)
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
	//  再返回页面之前先要校验是否存在username的Cookie
	username, err := c.Cookie("username")
	if err != nil {
		// 直接跳转到登录界面
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"username": username,
	})
}

func viphandlers(c *gin.Context) {
	tmpUsername, ok := c.Get("username")
	if !ok {
		// 如果取不到值，说明前面中间件出问题了
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	username, ok := tmpUsername.(string)
	if !ok {
		// 类型断言失败
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	c.HTML(http.StatusOK, "vip.html", gin.H{
		"username": username,
	})
}

// 中间件 基于用户登录认真验证
func cookieMiddeware(c *gin.Context) {
	//  再返回页面之前先要校验是否存在username的Cookie
	username, err := c.Cookie("username")
	if err != nil {
		// 直接跳转到登录界面
		toPath := fmt.Sprintf("%s?next=%s", "/login", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, toPath)
		return
	}
	//  用户已经登录
	c.Set("username", username)
	c.Next() // 继续处理后续函数
}
