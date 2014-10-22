package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"net/http"
	"strconv"
	"strings"
)

type StatusController struct {
	Contest
}

func (sc StatusController) Route(w http.ResponseWriter, r *http.Request) {
	sc.InitContest(w, r)
	action := sc.GetAction(r.URL.Path, 2)
	defer func() {
		if e := recover(); e != nil {
			http.Error(w, "no such page", 404)
		}
	}()
	rv := class.GetReflectValue(w, r)
	class.CallMethod(&sc, strings.Title(action), rv)
}

//TODO : list by arguments like :contest/status/list?cid=1&uid=vsake&solved=3
func (sc *StatusController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest Status List")

	solutionModel := model.SolutionModel{}
	qry := make(map[string]string)

	qry["module"] = strconv.Itoa(config.ModuleC)
	qry["mid"] = strconv.Itoa(sc.Cid)
	solutionList, err := solutionModel.List(qry)

	if err != nil {
		http.Error(w, "load error", 400)
		return
	}
	for i, v := range solutionList {
		solutionList[i].Pid = sc.Index[v.Pid]
	}
	sc.Data["Solution"] = solutionList
	sc.Data["Privilege"] = sc.Privilege
	sc.Data["IsContestStatus"] = true
	sc.Data["Privilege"] = sc.Privilege
	sc.Data["Uid"] = sc.Uid

	sc.Execute(w, "view/layout.tpl", "view/contest/status_list.tpl")
}

func (sc *StatusController) Code(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Status Code")

	args := r.URL.Query()
	sid, err := strconv.Atoi(args.Get("sid"))
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

	if one.Uid == sc.Uid || sc.Privilege >= config.PrivilegeTC {
		sc.Data["Solution"] = one
		sc.Data["Privilege"] = sc.Privilege
		sc.Data["Title"] = "View Code"
		sc.Data["IsCode"] = true
		sc.Execute(w, "view/layout.tpl", "view/contest/status_code.tpl")
	}
}
