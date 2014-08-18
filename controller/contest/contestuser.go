package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
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
