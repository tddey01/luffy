package ginsession

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"sync"
	"github.com/go-redis/redis"
	"time"
)

// redis版Session服务

type RedisSD struct {
	ID string
	Data map[string]interface{}
	rwLock sync.RWMutex // 读写锁，锁的是上面的Data
	expired int // 过期时间
	client *redis.Client // redis连接池
}

// NewRedisSessionData 构造函数
func NewRedisSessionData(id string, client *redis.Client)SessionData{
	return &RedisSD{
		ID: id,
		Data: make(map[string]interface{}, 8),
		client:client,
	}
}


func (r *RedisSD)GetID()string{
	return r.ID
}
func (r *RedisSD)Get(key string)(value interface{}, err error){
	// 获取读锁
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	value, ok := r.Data[key]
	if !ok{
		err = fmt.Errorf("invalid Key")
		return
	}
	return
}
func (r *RedisSD)Set(key string, value interface{}){
	// 获取写锁
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	r.Data[key] = value
}
func (r *RedisSD)Del(key string){
	// 删除key对应的键值对
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	delete(r.Data, key)
}
func (r *RedisSD)Save(){
	// 将最新的session data保存到redis中
	value, err := json.Marshal(r.Data)
	if err != nil {
		// 序列化session data 失败
		fmt.Printf("marshal session data failed, err:%v\n", err)
		return
	}
	// 将数据保存到redis
	r.client.Set(r.ID, value, time.Second*time.Duration(r.expired))
}

// SetExpire 设置过期时间
func (r *RedisSD)SetExpire(expired int){
	r.expired = expired
}



type RedisMgr struct{
	Session map[string]SessionData
	rwLock sync.RWMutex
	client *redis.Client  // redis连接池
}

// NewRedisMgr RedisMgr的构造函数
func NewRedisMgr()(Mgr){
	return &RedisMgr{
		Session: make(map[string]SessionData, 1024),
	}
}
func (r *RedisMgr)Init(addr string, options ...string){
	// 初始化Redis连接
	var (
		password string
		db string
	)
	if len(options) == 1{
		password = options[0]
	}else if len(options) == 2{
		password = options[0]
		db = options[1]
	}
	dbValue, err := strconv.Atoi(db)
	if err != nil {
		dbValue = 0
	}

	r.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       dbValue,  // use default DB
	})

	_, err = r.client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func (r *RedisMgr)loadFromRedis(sessionID string)(err error){
	// 1. 连接redis
	value, err := r.client.Get(sessionID).Result()
	if err != nil {
		// redis中没有该session_id对应的session data
		return
	}
	err = json.Unmarshal([]byte(value), &r.Session)
	if err != nil {
		// 从redis取出来的数据反序列化失败
		return
	}
	// 2. 根据sessionID找到对应的数据
	// 3. 把数据取出来反序列化到r.data
	return
}


// GetSessionData 获取sessionID对应的sessionData
func (r *RedisMgr) GetSessionData(sessionID string) (sd SessionData, err error) {
	// 1. r.Session中必须已经从Redis里面加载出来数据
	if r.Session == nil {
		err := r.loadFromRedis(sessionID)
		if err != nil {
			return nil, err
		}
	}
	// 2. r.Session[sessionID] 拿到对应的session data
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	sd, ok := r.Session[sessionID]
	if !ok {
		err = fmt.Errorf("invalid session id")
		return
	}
	return

}

func (r *RedisMgr) CreateSession() (sd SessionData) {
	// 1. 造一个sessionID
	uuidObj := uuid.NewV4()
	// 2. 造一个和它对应的SessionData
	sd = NewRedisSessionData(uuidObj.String(), r.client)
	r.Session[sd.GetID()] = sd // 把新创建的session data保存收到大仓库中
	// 3. 返回SessionData
	return
}
