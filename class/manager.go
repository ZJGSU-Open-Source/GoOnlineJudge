package class

import (
	"GoOnlineJudge/config"
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"sync"
	"time"
)

type Manager struct {
	cookieName    string
	sessions      map[string]*Session
	lock          sync.Mutex
	cookieExpires int64
}

func (m *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (m *Manager) StartSession(w http.ResponseWriter, r *http.Request) (session *Session) {
	m.lock.Lock()
	defer m.lock.Unlock()

	cookie, err := r.Cookie(m.cookieName)

	if err != nil || cookie.Value == "" || m.sessions[cookie.Value] == nil {
		sid := m.sessionId()
		session = NewSession(sid, w, r)
		m.sessions[sid] = session

		cookie := http.Cookie{
			Name:     m.cookieName,
			Value:    sid,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   config.CookieExpires,
		}
		http.SetCookie(w, &cookie)
	} else {
		sid := cookie.Value
		session = m.sessions[sid]
		session.Update()
	}
	return
}

func (m *Manager) DeleteSession(w http.ResponseWriter, r *http.Request) {
	m.lock.Lock()
	defer m.lock.Unlock()

	cookie, err := r.Cookie(m.cookieName)
	if err != nil {
		return
	}

	sid := cookie.Value
	delete(m.sessions, sid)

	newcookie := http.Cookie{
		Name:     m.cookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	http.SetCookie(w, &newcookie)
}

func (m *Manager) DeleteSess(sid string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	delete(m.sessions, sid)
}

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

func NewManager() *Manager {
	return &Manager{cookieName: "sessionID",
		cookieExpires: config.CookieExpires, sessions: make(map[string]*Session)}
}
