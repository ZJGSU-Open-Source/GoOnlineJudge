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
	switch this.Privilege {
	case config.PrivilegeTC:
		this.Teacher(w, r)
	case config.PrivilegeAD:
		this.Admin(w, r)
	default:

		if this.Privilege <= config.PrivilegePU {
			class.Logger.Info(r.RemoteAddr + " " + this.Uid + " try to visit Admin page")
			this.Data["Title"] = "Warning"
			this.Data["Info"] = "You are not admin!"
			err := this.Execute(w, "view/layout.tpl", "view/400.tpl")
			if err != nil {
				http.Error(w, "tpl error", 500)
				return
			}
			return
		}
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
	} else if args["user"] != "" {
		c = &UserController{}
		m = args["user"]
	} else if args["testdata"] != "" {
		c = &TestdataController{}
		m = args["testdata"]
	} else if args["image"] != "" {
		c = &ImageController{}
		m = args["image"]
	} else {
		class.Logger.Debug("args err")
		return
	}
	m = strings.Title(m)
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

func (this *AdminUserController) Teacher(w http.ResponseWriter, r *http.Request) {
	var c interface{}
	var m string
	args := this.ParseURL(r.URL.String())
	if len(args) == 1 {
		c = &HomeController{}
		m = "Home"
	} else if args["problem"] != "" {
		c = &ProblemController{}
		m = args["problem"]
	} else if args["contest"] != "" {
		c = &ContestController{}
		m = args["contest"]
	} else if args["user"] != "" {
		c = &UserController{}
		m = args["user"]
	} else if args["testdata"] != "" {
		c = &TestdataController{}
		m = args["testdata"]
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
