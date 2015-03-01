package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"encoding/json"
	"net/http"
	"restweb"
	"strconv"
	"time"
)

type AdminRejudge struct {
	class.Controller
} //@Controller

//@URL: /admin/rejudger/ @method: GET
func (pc *AdminRejudge) Index() {
	restweb.Logger.Debug("Rejudge Page")

	pc.Output["Title"] = "Problem Rejudge"
	pc.Output["RejudgePrivilege"] = true
	pc.Output["IsProblem"] = true
	pc.Output["IsRejudge"] = true

	pc.RenderTemplate("view/admin/layout.tpl", "view/admin/rejudge.tpl")
}

//@URL: /admin/rejudger/ @method: POST
func (pc *AdminRejudge) Rejudge() {
	restweb.Logger.Debug("Problem Rejudge")

	args := pc.R.URL.Query()
	types := args.Get("type")
	id, err := strconv.Atoi(args.Get("id"))
	if err != nil {
		pc.Error("args error", 400)
		return
	}

	hint := make(map[string]string)
	one := make(map[string]interface{})

	if types == "Pid" {
		pid := id
		proModel := model.ProblemModel{}
		pro, err := proModel.Detail(pid)
		if err != nil {
			restweb.Logger.Debug(err)
			hint["info"] = "Problem does not exist!"

			b, _ := json.Marshal(&hint)
			pc.W.WriteHeader(400)
			pc.W.Write(b)

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
			reader, _ := pc.PostReader(&one)
			_, err := http.Post(config.JudgeHost, "application/json", reader)
			if err != nil {
				// http.Error(w, "post error", 500)
				restweb.Logger.Debug(err)
			}
		}
	} else if types == "Sid" {
		sid := id

		solutionModel := model.SolutionModel{}
		sol, err := solutionModel.Detail(sid)
		if err != nil {
			restweb.Logger.Debug(err)

			hint["info"] = "Solution does not exist!"
			b, _ := json.Marshal(&hint)
			pc.W.WriteHeader(400)
			pc.W.Write(b)
			return
		}

		problemModel := model.ProblemModel{}
		pro, err := problemModel.Detail(sol.Pid)
		if err != nil {
			pc.Error(err.Error(), 500)
			return
		}
		one["Sid"] = sid
		one["Pid"] = pro.RPid
		one["OJ"] = pro.ROJ
		one["Rejudge"] = true
		reader, _ := pc.PostReader(&one)
		_, err = http.Post(config.JudgeHost, "application/json", reader)
		if err != nil {
			restweb.Logger.Debug("Sid[", sid, "] rejudger post error.")
			return
		}
	}
	pc.W.WriteHeader(200)
}
