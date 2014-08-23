package class

import (
	"net/http"
	"sync"
	"time"
)

type Session struct {
	Sid          string
	data         map[string]string
	timeAccessed time.Time
	sync.Mutex
}

func (this *Session) Set(Key, Value string) {
	this.Lock()
	defer this.Unlock()
	this.Update()
	this.data[Key] = Value
}

func (this *Session) Get(Key string) string {
	this.Lock()
	defer this.Unlock()
	this.Update()
	return this.data[Key]
}

func (this *Session) Update() {
	this.timeAccessed = time.Now()
}

func NewSession(sid string, w http.ResponseWriter, r *http.Request) (session *Session) {
	session = &Session{Sid: sid}
	session.data = make(map[string]string)
	session.Update()
	return
}
