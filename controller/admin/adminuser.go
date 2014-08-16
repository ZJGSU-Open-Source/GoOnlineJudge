package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
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
		this.Data["Title"] = "Warning"
		this.Data["Info"] = "You are not admin!"
		t := template.New("layout.tpl")
		t, err := t.ParseFiles("view/layout.tpl", "view/400.tpl")
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
		err = t.Execute(w, this.Data)
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
		return
	}
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
