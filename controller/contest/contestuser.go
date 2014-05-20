package contest

import (
	"GoOnlineJudge/config"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type ContestUserContorller struct {
	Contest
}

func (this *ContestUserContorller) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("Contest User")
	this.InitContest(w, r)

	if this.ContestDetail.Encrypt == config.EncryptPW {
		if reflect.ValueOf(this.ContestDetail.Argument) != this.Data[strconv.Itoa(this.ContestDetail.Cid)] {
			this.Data["Title"] = "Warning"
			this.Data["Info"] = "You don't have permission to participate in!"
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
	}
	var c interface{}
	p := strings.Trim(r.URL.Path, "/")
	s := strings.Split(p, "/")
	if len(s) < 3 {
		log.Println("args err")
		return
	}
	switch s[1] {
	case "problem":
		c = &ProblemController{}
	case "status":
		c = &StatusController{}
	case "ranklist":
		c = &RanklistController{}
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
