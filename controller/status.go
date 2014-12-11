package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"net/http"
	"restweb"
	"strconv"
)

type StatusController struct {
	class.Controller
}

func (sc StatusController) List() {

	restweb.Logger.Debug("Status List")

	searchUrl := ""
	qry := make(map[string]string)
	// Search
	if v, ok := sc.Input["uid"]; ok {
		searchUrl += "uid=" + v[0] + "&"
		sc.Output["SearchUid"] = v[0]
		qry["uid"] = v[0]
	}
	if v, ok := sc.Input["pid"]; ok {
		searchUrl += "pid=" + v[0] + "&"
		sc.Output["SearchPid"] = v[0]
		qry["pid"] = v[0]
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
	sc.Output["URL"] = "/status?" + searchUrl

	// Page
	qry["page"] = "1"

	if v, ok := sc.Input["page"]; ok {
		qry["page"] = v[0]
	}

	solutionModel := model.SolutionModel{}
	qry["module"] = strconv.Itoa(config.ModuleP)
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

	list, err := solutionModel.List(qry)
	if err != nil {
		sc.Error(err.Error(), 500)
		return
	}

	sc.Output["Solution"] = list
	sc.Output["Title"] = "Status List"
	sc.Output["IsStatus"] = true
	sc.Output["Privilege"] = sc.Privilege
	sc.Output["Uid"] = sc.Uid

	sc.RenderTemplate("view/layout.tpl", "view/status_list.tpl")
}

func (sc *StatusController) Code() {
	restweb.Logger.Debug("Status Code")

	sid, err := strconv.Atoi(sc.Input.Get("sid"))
	if err != nil {
		http.Error(sc.W, "args error", 400)
		return
	}

	solutionModel := model.SolutionModel{}
	one, err := solutionModel.Detail(sid)
	if err != nil {
		http.Error(sc.W, err.Error(), 400)
		return
	}
	if one.Error != "" {
		one.Code = one.Code + "\n/*\n" + one.Error + "*/\n"
	}

	if one.Uid == sc.Uid || sc.Privilege > config.PrivilegePU {
		sc.Output["Solution"] = one
		sc.Output["Title"] = "View Code"
		sc.Output["IsCode"] = true
		sc.RenderTemplate("view/layout.tpl", "view/status_code.tpl")
	} else {
		sc.Err400("Warning", "You can't see it!")
	}
}
