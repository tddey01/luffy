package seesion

import (
	"fmt"
	"sync"
	"time"
)

// session  服务

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
	Session map[string]SessionData
	rwLock  sync.RWMutex
}

// GetSessionData  获取用户sessionid
func (m *SessionMgr) GetSessionData(sessionID string) (sd SessionData, err error) {
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
func (m *SessionMgr) CreateSession() (sd *SessionData, err error) {
	// 1. 造一个sessionID
	temptime := time.Now().UnixNano()
	if err != nil {
		return
	}
	//  2  造一个和其他对应的SeessionData
	sd = NewSessionData(string(temptime))
	//  3 返回SessionData
	return
}
