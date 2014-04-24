package class

import (
	"GoOnlineJudge/config"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Controller struct {
	Data      map[string]interface{}
	Uid       string
	Privilege int
}

func (this *Controller) Init(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	this.Data = make(map[string]interface{})

	this.Uid = this.GetSession(w, r, "CurrentUser")
	if this.Uid != "" {
		this.Data["IsCurrentUser"] = true
		this.Data["CurrentUser"] = this.Uid

		var err error
		this.Privilege, err = strconv.Atoi(this.GetSession(w, r, "CurrentPrivilege"))
		if err != nil {
			http.Error(w, "args error", 400)
			return
		}

		if this.Privilege > config.PrivilegeSB {
			this.Data["IsShowAdmin"] = true
		}
	}
}

func (this *Controller) ParseURL(url string) (args map[string]string) {
	args = make(map[string]string)
	path := strings.Trim(url, "/")
	list := strings.Split(path, "/")

	for i := 1; i < len(list); i += 2 {
		args[list[i-1]] = list[i]
	}
	return
}

func (this *Controller) GetTime() (ft string) {
	t := time.Now().Unix()
	ft = time.Unix(t, 0).Format("2006-01-02 15:04:05")
	return
}

func (this *Controller) LoadJson(r io.Reader, i interface{}) (err error) {
	err = json.NewDecoder(r).Decode(i)
	return
}

func (this *Controller) PostReader(i interface{}) (r io.Reader, err error) {
	b, err := json.Marshal(i)
	if err != nil {
		return
	}
	r = strings.NewReader(string(b))
	return
}

func (this *Controller) SetSession(w http.ResponseWriter, r *http.Request, name string, value string) {
	s := Session{
		Name:  name,
		Value: value,
	}
	s.Set(w, r)
}

func (this *Controller) GetSession(w http.ResponseWriter, r *http.Request, name string) (value string) {
	s := Session{
		Name: name,
	}
	value = s.Get(w, r)
	return
}

func (this *Controller) DeleteSession(w http.ResponseWriter, r *http.Request, name string) {
	s := Session{
		Name: name,
	}
	s.Delete(w, r)
}
