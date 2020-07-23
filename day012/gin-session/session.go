package ginsession

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	SessionCookieName = "session_id" // sesion_id在Cookie中对应的key
	SessionContextName = "session" // session data在gin上下文中对应的key
)

var (
	// MgrObj 全局的Session管理对象（大仓库）
	MgrObj Mgr
)

// 自己实现的gin框架的session中间件

// Session服务

// SessionData 表示一个具体的用户Session数据

// 课上版本存在问题：
// 1. 调用save的时候不管秀没有修改都会保存数据库 ， sessionData定义一个标志位：r.modifyFlag
// 2. 从redis中加载数据，会存在一个问题 sync.Once 用来加载一次！
// 3. 过期时间

type Option struct {
	MaxAge int
	Path string
	Domain string
	Secure bool
	HttpOnly bool
}


type SessionData interface {
	GetID()string // 返回自己的ID
	Get(key string)(value interface{}, err error)
	Set(key string, value interface{})
	Del(key string)
	Save() // 保存
	SetExpire(int) // 设置过期时间
}
// SessionData支持的操作


// Mgr 所有类型的大仓库都应该遵循的接口
type Mgr interface {
	Init(addr string, options ...string) // 所有支持的后端都必须实现Init()来执行具体的连接
	GetSessionData(sessionID string)(sd SessionData, err error)
	CreateSession()(sd SessionData)
}

func InitMgr(name string, addr string, options...string) {
	switch name {
	case "memory":
		MgrObj = NewMemoryMgr()
	case "redis":
		MgrObj = NewRedisMgr()
	}
	MgrObj.Init(addr, options...) // 初始化Mgr
}


// 实现一个gin框架的中间件
// 所有流经我这个中间件的请求，它的上下文中肯定会有一个session -> session data
func SessionMiddleware(mgrObj Mgr, option *Option)gin.HandlerFunc {
	if mgrObj == nil {
		panic("must call InitMgr before use it.")
	}
	return func(c *gin.Context){
		// 1. 从请求的Cookie中获取session_id
		var sd SessionData // session data
		sessionID, err := c.Cookie(SessionCookieName)
		fmt.Println(sessionID)
		if err != nil {
			// 1.1 取不到session_id -> 给这个新用户创建一个新的session data，同时分配一个session_id
			sd = mgrObj.CreateSession()
			sessionID = sd.GetID()
			fmt.Println("取不到session_id，创建一个新的", sessionID)
		}else {
			// 1.2 取到session_id
			// 2. 根据session_id去Session大仓库中取到对应的session data
			sd, err = mgrObj.GetSessionData(sessionID)
			if err != nil {
				// 2.1 根据用户传过来的session_id在大仓库中根本取不到session data
				sd = mgrObj.CreateSession() // sd
				// 2.2 更新用户Cookie中保存的那个session_id
				sessionID = sd.GetID()
				fmt.Println("session_id取不到session data,分配一个新的", sessionID)
			}
			fmt.Println("session_id未过期", sessionID)
		}
		sd.SetExpire(option.MaxAge) // 设置session data过期时间
		// 3. 如何实现让后续所有的处理请求的方法都能拿到session data
		// 3. 利用gin的c.Set("session", session data)
		c.Set(SessionContextName, sd) // 保存到上下文
		// 在gin框架中，要回写Cookie必须在处理请求的函数返回之前
		c.SetCookie(SessionCookieName, sessionID, option.MaxAge, option.Path, option.Domain, option.Secure, option.HttpOnly)
		c.Next() // 执行后续的请求处理方法 c.HTML()时已经把响应头写好了
	}
}
