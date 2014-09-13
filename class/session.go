package class

import (
	"net/http"
	"sync"
	"time"
)

// Session 服务端会话
type Session struct {
	Sid          string
	data         map[string]string
	timeAccessed time.Time
	sync.Mutex
}

// Set 设置session，添加key－value
func (this *Session) Set(Key, Value string) {
	this.Lock()
	defer this.Unlock()
	this.Update()
	this.data[Key] = Value
}

// Get 获得key指定的value
func (this *Session) Get(Key string) string {
	this.Lock()
	defer this.Unlock()
	this.Update()
	return this.data[Key]
}

// Update 更新session时间
func (this *Session) Update() {
	this.timeAccessed = time.Now()
}

// NewSession 生成一个新的session
func NewSession(sid string, w http.ResponseWriter, r *http.Request) (session *Session) {
	session = &Session{Sid: sid}
	session.data = make(map[string]string)
	session.Update()
	return
}
