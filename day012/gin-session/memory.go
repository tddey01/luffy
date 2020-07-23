package ginsession

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
)

// 内存版Session服务
// 仅供参考使用


type MemSD struct {
	ID string
	Data map[string]interface{}
	rwLock sync.RWMutex // 读写锁，锁的是上面的Data
	// 过期时间
}

// NewMemorySessionData MemSD 的构造函数
func NewMemorySessionData(id string)SessionData{
	return &MemSD{
		ID: id,
		Data: make(map[string]interface{}, 8),
	}
}

func (m *MemSD)GetID()string{
	return m.ID
}
// Get 根据key获取值
func (m *MemSD)Get(key string)(value interface{}, err error){
	// 获取读锁
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	value, ok := m.Data[key]
	if !ok{
		err = fmt.Errorf("invalid Key")
		return
	}
	return
}

// Set 根据key获取值
func (m *MemSD)Set(key string, value interface{}){
	// 获取写锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	m.Data[key] = value
}

// Del 删除Key对应的键值对
func (m *MemSD)Del(key string){
	// 删除key对应的键值对
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	delete(m.Data, key)
}

// Save 保存session data
func (m *MemSD)Save(){
	return
}

func (m *MemSD)SetExpire(expired int){
	return
}
// Mgr 是一个全局的Session 管理
type MemoryMgr struct {
	Session map[string]SessionData
	rwLock sync.RWMutex
}

// NewMemoryMgr 内存版session大仓库的构造函数
func NewMemoryMgr()(Mgr){
	return &MemoryMgr{
		Session:make(map[string]SessionData, 1024), // 初始化1024红色的小框用来存取用户的session data
	}
}


func (m *MemoryMgr)Init(addr string, options ...string){
	return
}

// GetSessionData 根据传进来的SessionID找到对应的SessionData
func (m *MemoryMgr)GetSessionData(sessionID string)(sd SessionData, err error){
	// 取之前加锁
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	sd, ok := m.Session[sessionID]
	if !ok {
		err = fmt.Errorf("invalid session id")
		return
	}
	return
}

// CreateSession 创建一条Session记录
func (m *MemoryMgr)CreateSession()(sd SessionData){
	// 1. 造一个sessionID
	uuidObj := uuid.NewV4()
	// 2. 造一个和它对应的SessionData
	sd = NewMemorySessionData(uuidObj.String())
	m.Session[sd.GetID()] = sd // 把新创建的session data保存收到大仓库中
	// 3. 返回SessionData
	return
}

