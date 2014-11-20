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
}

func (pc *AdminRejudge) Index() {
	restweb.Logger.Debug("Rejudge Page")

	if pc.Privilege < config.PrivilegeTC {
		pc.Err400("Warning", "Error Privilege to Rejudge problem")
		return
	}

	pc.Data["Title"] = "Problem Rejudge"
	pc.Data["RejudgePrivilege"] = true
	pc.Data["IsProblem"] = true
	pc.Data["IsRejudge"] = true

	pc.RenderTemplate("view/admin/layout.tpl", "view/admin/rejudge.tpl")
}

func (pc *AdminRejudge) Rejudge() {
	restweb.Logger.Debug("Problem Rejudge")

	if pc.Privilege < config.PrivilegeTC {
		pc.Err400("Warning", "Error Privilege to Rejudge problem")
		return
	}

	args := pc.Requset.URL.Query()
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
			pc.Response.WriteHeader(400)
			pc.Response.Write(b)

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
			one["Time"] = pro.Time
			one["Memory"] = pro.Memory
			one["Rejudge"] = true
			reader, _ := pc.PostReader(&one)
			response, err := http.Post(config.JudgeHost, "application/json", reader)
			if err != nil {
				// http.Error(w, "post error", 500)
				restweb.Logger.Debug(err)
			} else {
				response.Body.Close()
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
			pc.Response.WriteHeader(400)
			pc.Response.Write(b)
			return
		}

		problemModel := model.ProblemModel{}
		pro, err := problemModel.Detail(sol.Pid)
		if err != nil {
			pc.Error(err.Error(), 500)
			return
		}
		one["Sid"] = sid
		one["Time"] = pro.Time
		one["Memory"] = pro.Memory
		one["Rejudge"] = true
		reader, _ := pc.PostReader(&one)
		restweb.Logger.Debug(reader)
		response, err := http.Post(config.JudgeHost, "application/json", reader)
		if err != nil {
			pc.Error("post error", 500)
			return
		}
		defer response.Body.Close()
	}
	pc.Response.WriteHeader(200)
}
