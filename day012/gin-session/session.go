package ginsession

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
	SessionCookieName  = "session_id" // session id 在cookie中对应的key
	SessionContextName = "session"    //session data在gin上下文中对应的key
)

// session  服务
var (
	MgrObj *SessionMgr // 定义一个全局的 Session 管理对象 大仓库
)

// SessionData 回话标签
type SessionData struct {
	ID     string
	Data   map[string]interface{}
	rwLock sync.RWMutex // 读写锁，锁的是上面的Data
	// 过期时间
}

// NewSessionData 构造函数 工厂模式
func NewSessionData(id string) *SessionData {
	return &SessionData{
		ID:   id,
		Data: make(map[string]interface{}, 8),
	}
}

// SessionMgr 是一个全局的Session 管理
type SessionMgr struct {
	Session map[string]*SessionData
	rwLock  sync.RWMutex
}

func IntMgr() {
	MgrObj = &SessionMgr{
		Session: make(map[string]*SessionData, 1024), // 初始化1024红色小框 用来存储用户的seesion data
	}
}

//func IntMgr(name string, addr string, options...string) {
//	MgrObj =  &SessionMgr{
//		Session:make(map[string]*SessionData, 1024), // 初始化1024红色的小框用来存取用户的session data
//	}
//}
// GetSessionData  获取用户sessionid
func (m *SessionMgr) GetSessionData(sessionID string) (sd *SessionData, err error) {
	// 取之前枷锁
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	sd, ok := m.Session[sessionID]
	if !ok {
		err = fmt.Errorf("invalid seesion id")
		return
	}
	return
}

// CreateSession 创建一条Session记录
func (m *SessionMgr) CreateSession() (sd *SessionData) {
	// 1. 造一个sessionID
	uuidObj := uuid.NewV4()
	//  2  造一个和其他对应的SeessionData
	sd = NewSessionData(uuidObj.String())
	m.Session[sd.ID] = sd // 把新创建的session data保存收到大仓库中
	//  3 返回SessionData
	return
}

//  实现一个gin框架的中间件
// 所有流经过我和中间件的请求， 他的上下文中肯定会有一个Session --> session data
func SessionMiddleware(mgrObj *SessionMgr) gin.HandlerFunc {
	if mgrObj == nil {
		panic("must call InitMgr before use it 请初始化")
	}
	return func(c *gin.Context) {
		//  1 从请求中cookie获取session_id
		var sd *SessionData // session data
		sessionID, err := c.Cookie(SessionCookieName)
		fmt.Println(sessionID)
		if err != nil {
			//  1.1 取不到session_id --> 给这个新的用户创建新的session_id 同事分配一个session_id
			sd = mgrObj.CreateSession()
			sessionID = sd.ID
			fmt.Println("取不到session_id，创建一个新的", sessionID)
		} else {
			//	1.2 取到session_id
			//	2 根据session_id去session大仓库中取到对应的session_data
			sd, err = mgrObj.GetSessionData(sessionID)
			if err != nil {
				//	2.1 根据用户传过来的session_id在仓中取不到session data
				sd = mgrObj.CreateSession()
				//	 2.2 更新哦用户cookie中保存的那个session_id
				sessionID = sd.ID
				fmt.Println("session_id取不到session data,分配一个新的", sessionID)
			}
		}
		//	3 如果实现了让后续所有的处理方法都能拿到session data
		//  4 利用gin框架c.set("session","session data)
		c.Set(SessionContextName, sd)
		// 在gin框架中 要会写cookie必须处理请求的函数返回之前
		c.SetCookie(SessionCookieName, sessionID, 3600, "/", "127.0.0.1", false, true)
		c.Next() // 执行后续的请求处理方法 c.HTML() 时已经把相应头写好了
	}
}
