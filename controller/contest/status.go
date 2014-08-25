package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"net/http"
	"strconv"
)

type StatusController struct {
	Contest
}

func (this *StatusController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest Status List")
	this.InitContest(w, r)

	solutionModel := model.SolutionModel{}
	qry := make(map[string]string)

	qry["module"] = strconv.Itoa(config.ModuleC)
	qry["mid"] = strconv.Itoa(this.Cid)
	solutionList, err := solutionModel.List(qry)

	if err != nil {
		http.Error(w, "load error", 400)
		return
	}
	for i, v := range solutionList {
		solutionList[i].Pid = this.Index[v.Pid]
	}
	this.Data["Solution"] = solutionList
	this.Data["Privilege"] = this.Privilege
	this.Data["IsContestStatus"] = true
	err = this.Execute(w, "view/layout.tpl", "view/contest/status_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *StatusController) Code(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Status Code")
	this.Init(w, r)
	this.InitContest(w, r)

	args := this.ParseURL(r.URL.String())
	sid, err := strconv.Atoi(args["sid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	solutionModel := model.SolutionModel{}
	one, err := solutionModel.Detail(sid)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	this.Data["Solution"] = one
	this.Data["Privilege"] = this.Privilege
	this.Data["Title"] = "View Code"
	this.Data["IsCode"] = true
	err = this.Execute(w, "view/layout.tpl", "view/contest/status_code.tpl")
	if err != nil {
		class.Logger.Debug(err)
		http.Error(w, "tpl error", 500)
		return
	}
}
