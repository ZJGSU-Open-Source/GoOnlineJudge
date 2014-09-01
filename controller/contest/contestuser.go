package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type ContestUserContorller struct {
	Contest
}

func (this *ContestUserContorller) Register(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest User")
	this.InitContest(w, r)

	if this.ContestDetail.Encrypt == config.EncryptPW && this.Privilege <= config.PrivilegePU {
		if this.Uid == "" {
			http.Redirect(w, r, "/user?signin", http.StatusFound)
			return
		} else if this.GetSession(w, r, strconv.Itoa(this.Cid)) != this.ContestDetail.Argument.(string) {
			this.Password(w, r)
			return
		}
	} else if this.ContestDetail.Encrypt == config.EncryptPT && this.Privilege <= config.PrivilegePU {
		if this.Uid == "" {
			http.Redirect(w, r, "/user?signin", http.StatusFound)
		} else {
			userlist := strings.Split(this.ContestDetail.Argument.(string), "\n")
			flag := false
			for _, user := range userlist {
				class.Logger.Debug(user)
				if user == this.Uid {
					flag = true
					break
				}
			}
			if flag == false {
				this.Data["Title"] = this.ContestDetail.Title
				this.Data["Info"] = "Sorry, the contest is private and you are not granted to participate in the contest."
				err := this.Execute(w, "view/layout.tpl", "view/400.tpl")
				if err != nil {
					http.Error(w, "tpl error", 500)
					return
				}
				return
			}
		}
	}
	var c interface{}
	var m string

	args := this.ParseURL(r.URL.String())
	class.Logger.Debug(args)
	if args["problem"] != "" {
		c = &ProblemController{}
		m = strings.Title(args["problem"])
	} else if args["status"] != "" {
		c = &StatusController{}
		m = strings.Title(args["status"])
	} else if args["ranklist"] != "" {
		c = &RanklistController{}
		m = "List"
	} else {
		class.Logger.Debug("args err")
		return
	}
	class.Logger.Debug(m)
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

func (this *ContestUserContorller) Password(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := this.Execute(w, "view/layout.tpl", "view/contest/passwd.tpl")
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
		return
	} else if r.Method == "POST" {
		passwd := r.FormValue("password")
		if passwd == this.ContestDetail.Argument.(string) {
			this.SetSession(w, r, strconv.Itoa(this.Cid), passwd)
			w.WriteHeader(200)
		} else {
			http.Error(w, "incorrect password", 400)
		}
	}
}
func callMethod(c interface{}, m string, rv []reflect.Value) {
	rc := reflect.ValueOf(c)
	rm := rc.MethodByName(m)
	rm.Call(rv)
}

func getReflectValue(w http.ResponseWriter, r *http.Request) (rv []reflect.Value) {
	rw := reflect.ValueOf(w)
	rr := reflect.ValueOf(r)
	rv = []reflect.Value{rw, rr}
	return
}
