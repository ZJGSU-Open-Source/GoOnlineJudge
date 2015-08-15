package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"sync"
	"time"
)

// Manager 是session控制器
type Manager struct {
	cookieName    string
	sessions      map[string]*Session
	lock          sync.Mutex
	cookieExpires int64
}

// sessionId() 随机生成一个session id
func (m *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// StartSession 建立session会话
func (m *Manager) StartSession() (session *Session) {
	m.lock.Lock()
	defer m.lock.Unlock()

	sid := m.sessionId()
	session = NewSession(sid)
	m.sessions[sid] = session

	return
}

// StartSession 建立session会话
func (m *Manager) GetSession(sid string) (session *Session) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if session, ok := m.sessions[sid]; ok {
		return session
	}

	return nil
}

// DeleteSession 删除session并设置客户端cookie
func (m *Manager) DeleteSession(sid string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	delete(m.sessions, sid)
}

// DeleteSess 安全删除session
func (m *Manager) DeleteSess(sid string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	delete(m.sessions, sid)
}

// GC 垃圾回收，释放过期session
func (m *Manager) GC() {
	for {
		time.Sleep(time.Duration(m.cookieExpires) * time.Second)
		for sid, sess := range m.sessions {
			if sess.timeAccessed.Unix()+m.cookieExpires < time.Now().Unix() {
				m.DeleteSess(sid)
			}
		}
	}
}

// NewManager 生成一个新的Manager
func NewManager() *Manager {
	return &Manager{cookieName: "sessionID",
		cookieExpires: 1800, sessions: make(map[string]*Session)}
}
