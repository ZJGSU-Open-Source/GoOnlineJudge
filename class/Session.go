package class

import (
	"GoOnlineJudge/config"
	"net/http"
	"sync"
)

type Session struct {
	Name  string
	Value string
	sync.Mutex
}

func (this *Session) Set(w http.ResponseWriter, r *http.Request) {
	this.Lock()
	defer this.Unlock()

	cookie := http.Cookie{
		Name:   this.Name,
		Value:  this.Value,
		Path:   "/",
		MaxAge: config.CookieExpires,
	}
	http.SetCookie(w, &cookie)
}

func (this *Session) Get(w http.ResponseWriter, r *http.Request) interface{} {
	this.Lock()
	defer this.Unlock()

	cookie, err := r.Cookie(this.Name)
	if err != nil || cookie.Value == "" {
		return nil
	} else {
		cookie.MaxAge = config.CookieExpires
		http.SetCookie(w, cookie)
		return cookie.Value
	}
}

func (this *Session) Delete(w http.ResponseWriter, r *http.Request) {
	this.Lock()
	defer this.Unlock()

	cookie := http.Cookie{
		Name:   this.Name,
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
}
