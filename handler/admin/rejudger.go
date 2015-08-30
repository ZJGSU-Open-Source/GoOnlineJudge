package admin

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"github.com/zenazn/goji/web"

	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//@URL: /api/admin/rejudger @method: POST
func Rejudge(c web.C, w http.ResponseWriter, r *http.Request) {

	args := r.URL.Query()
	types := args.Get("type")
	id, err := strconv.Atoi(args.Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hint := make(map[string]string)
	one := make(map[string]interface{})

	if types == "Pid" {
		pid := id
		proModel := model.ProblemModel{}
		pro, err := proModel.Detail(pid)
		if err != nil {
			hint["info"] = "Problem does not exist!"

			b, _ := json.Marshal(&hint)
			w.WriteHeader(400)
			w.Write(b)

			return
		}
		qry := make(map[string]string)
		qry["pid"] = strconv.Itoa(pro.Pid)

		solutionModel := model.SolutionModel{}
		list, err := solutionModel.List(qry)

		for i := range list {
			sid := list[i].Sid
			time.Sleep(1 * time.Second)
			one["Sid"] = sid
			one["Pid"] = pro.RPid
			one["OJ"] = pro.ROJ
			one["Rejudge"] = true
			reader, _ := JsonReader(&one)
			_, err := http.Post(config.JudgeHost, "application/json", reader)
			if err != nil {
				// restweb.Logger.Debug(err)
			}
		}
	} else if types == "Sid" {
		sid := id

		solutionModel := model.SolutionModel{}
		sol, err := solutionModel.Detail(sid)
		if err != nil {

			hint["info"] = "Solution does not exist!"
			b, _ := json.Marshal(&hint)
			w.WriteHeader(404)
			w.Write(b)
			return
		}

		problemModel := model.ProblemModel{}
		pro, err := problemModel.Detail(sol.Pid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		one["Sid"] = sid
		one["Pid"] = pro.RPid
		one["OJ"] = pro.ROJ
		one["Rejudge"] = true
		reader, _ := JsonReader(&one)
		_, err = http.Post(config.JudgeHost, "application/json", reader)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(200)
}

func JsonReader(i interface{}) (r io.Reader, err error) {
	b, err := json.Marshal(i)
	if err != nil {
		return
	}
	r = strings.NewReader(string(b))
	return
}
