package session

import (
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
func (s *Session) Set(Key, Value string) {
    s.Lock()
    defer s.Unlock()
    s.Update()
    s.data[Key] = Value
}

// Get 获得key指定的value
func (s *Session) Get(Key string) string {
    s.Lock()
    defer s.Unlock()
    s.Update()
    return s.data[Key]
}

// Update 更新session时间
func (s *Session) Update() {
    s.timeAccessed = time.Now()
}

// NewSession 生成一个新的session
func NewSession(sid string) (session *Session) {
    session = &Session{Sid: sid}
    session.data = make(map[string]string)
    session.Update()
    return
}
