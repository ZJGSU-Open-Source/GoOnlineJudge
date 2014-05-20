package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"log"
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
	if r.URL.Path[6:] == "/" {
		c = &HomeController{}
		m := "Home"
		rv := getReflectValue(w, r)
		callMethod(c, m, rv)
		return
	}
	p := strings.Trim(r.URL.Path, "/")
	s := strings.Split(p, "/")
	if len(s) < 3 {
		log.Println("args err")
		return
	}

	switch s[1] {
	case "news":
		c = &NewsController{}
	case "problem":
		c = &ProblemController{}
	case "contest":
		c = &ContestController{}
	case "user":
		c = &UserController{}
	case "testdata":
		c = &TestdataController{}
	default:
		log.Println("args err")
		return
	}
	m := strings.Title(s[2])
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
