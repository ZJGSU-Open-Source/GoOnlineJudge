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

func (sc *ContestStatus) List(Cid string) {
	restweb.Logger.Debug("Contest Status List")

	sc.InitContest(Cid)
	solutionModel := model.SolutionModel{}
	qry := make(map[string]string)

	qry["module"] = strconv.Itoa(config.ModuleC)
	qry["mid"] = Cid
	searchUrl := ""
	// Search
	if v, ok := sc.Input["uid"]; ok {
		searchUrl += "uid=" + v[0] + "&"
		sc.Output["SearchUid"] = v[0]
		qry["uid"] = v[0]
	}
	if v, ok := sc.Input["pid"]; ok {
		searchUrl += "pid=" + v[0] + "&"
		sc.Output["SearchPid"] = v[0]
		idx, _ := strconv.Atoi(v[0])
		qry["pid"] = strconv.Itoa(sc.ContestDetail.List[idx])
		restweb.Logger.Debug(qry["pid"], idx)
	}
	if v, ok := sc.Input["judge"]; ok {
		searchUrl += "judge=" + v[0] + "&"
		sc.Output["SearchJudge"+v[0]] = v[0]
		qry["judge"] = v[0]
	}
	if v, ok := sc.Input["language"]; ok {
		searchUrl += "language=" + v[0] + "&"
		sc.Output["SearchLanguage"+v[0]] = v[0]
		qry["language"] = v[0]
	}

	qry["page"] = "1"
	if v, ok := sc.Input["page"]; ok {
		qry["page"] = v[0]
	}

	sc.Output["URL"] = "/contests/" + Cid + "/status?" + searchUrl

	qry["action"] = "submit"
	count, err := solutionModel.Count(qry)
	if err != nil {
		sc.Error(err.Error(), 400)
		return
	}
	var pageCount = (count-1)/config.SolutionPerPage + 1

	page, err := strconv.Atoi(qry["page"])
	if err != nil {
		sc.Error("args error", 400)
		return
	}
	if page > pageCount {
		sc.Error("args error", 400)
		return
	}
	qry["offset"] = strconv.Itoa((page - 1) * config.SolutionPerPage)
	qry["limit"] = strconv.Itoa(config.SolutionPerPage)

	pageData := sc.GetPage(page, pageCount)
	for k, v := range pageData {
		sc.Output[k] = v
	}

	solutionList, err := solutionModel.List(qry)

	if err != nil {
		sc.Error("load error", 400)
		return
	}
	for i, v := range solutionList {
		solutionList[i].Pid = sc.Index[v.Pid]
	}
	sc.Output["Solution"] = solutionList
	sc.Output["Privilege"] = sc.Privilege
	sc.Output["IsContestStatus"] = true
	sc.Output["Privilege"] = sc.Privilege
	sc.Output["Uid"] = sc.Uid

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
		sc.Output["Solution"] = one
		sc.Output["Privilege"] = sc.Privilege
		sc.Output["Title"] = "View Code"
		sc.Output["IsCode"] = true
		sc.RenderTemplate("view/layout.tpl", "view/contest/status_code.tpl")
	}
}
