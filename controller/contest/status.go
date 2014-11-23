package contest

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"strconv"
)

type ContestStatus struct {
	Contest
}

//TODO : list by arguments like :contest/status/list?cid=1&uid=vsake&solved=3
func (sc *ContestStatus) List(Cid string) {
	restweb.Logger.Debug("Contest Status List")

	sc.InitContest(Cid)
	solutionModel := model.SolutionModel{}
	qry := make(map[string]string)

	qry["module"] = strconv.Itoa(config.ModuleC)
	qry["mid"] = Cid
	solutionList, err := solutionModel.List(qry)

	if err != nil {
		sc.Error("load error", 400)
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

	sc.RenderTemplate("view/layout.tpl", "view/contest/status_list.tpl")
}

func (sc *ContestStatus) Code(Cid string, Sid string) {
	restweb.Logger.Debug("Status Code")

	sc.InitContest(Cid)
	sid, err := strconv.Atoi(Sid)
	if err != nil {
		sc.Error("args error", 400)
		return
	}

	solutionModel := model.SolutionModel{}
	one, err := solutionModel.Detail(sid)
	if err != nil {
		sc.Error(err.Error(), 500)
		return
	}

	if one.Error != "" {
		one.Code += "\n/*\n" + one.Error + "\n*/"
	}
	if one.Uid == sc.Uid || sc.Privilege >= config.PrivilegeTC {
		sc.Data["Solution"] = one
		sc.Data["Privilege"] = sc.Privilege
		sc.Data["Title"] = "View Code"
		sc.Data["IsCode"] = true
		sc.RenderTemplate("view/layout.tpl", "view/contest/status_code.tpl")
	}
}
