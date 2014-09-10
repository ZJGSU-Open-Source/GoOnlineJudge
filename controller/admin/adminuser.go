package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"net/http"
	"reflect"
	"strings"
)

type AdminUserController struct {
	class.Controller
}

func (this *AdminUserController) Register(w http.ResponseWriter, r *http.Request) {
	this.Init(w, r)

	if this.Privilege <= config.PrivilegePU {
		this.Err400(w, r, "Warning", "You are not admin!")
		return
	} else {
		this.Admin(w, r)
	}
}

func (this *AdminUserController) Admin(w http.ResponseWriter, r *http.Request) {
	var c interface{}
	var m string
	args := this.ParseURL(r.URL.String())
	if len(args) == 1 {
		c = &HomeController{}
		m = "Home"
	} else if args["news"] != "" {
		c = &NewsController{}
		m = args["news"]
	} else if args["problem"] != "" {
		c = &ProblemController{}
		m = args["problem"]
	} else if args["contest"] != "" {
		c = &ContestController{}
		m = args["contest"]
	} else if args["testdata"] != "" {
		c = &TestdataController{}
		m = args["testdata"]
	} else if this.Privilege >= config.PrivilegeAD {
		if args["user"] != "" {
			c = &UserController{}
			m = args["user"]
		} else if args["image"] != "" {
			c = &ImageController{}
			m = args["image"]
		}
	} else {
		class.Logger.Debug("args err")
		return
	}
	m = strings.Title(m)
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
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
