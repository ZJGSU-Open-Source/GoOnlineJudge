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
	args := sc.Requset.URL.Query()
	searchUrl := ""
	qry := make(map[string]string)
	// Search
	if v := args.Get("uid"); v != "" {
		searchUrl += "uid=" + v + "&"
		sc.Data["SearchUid"] = v
		qry["uid"] = v
	}
	if v := args.Get("pid"); v != "" {
		searchUrl += "pid=" + v + "&"
		sc.Data["SearchPid"] = v
		qry["pid"] = v
	}
	if v := args.Get("judge"); v != "" {
		searchUrl += "judge=" + v + "&"
		sc.Data["SearchJudge"+v] = v
		qry["judge"] = v
	}
	if v := args.Get("language"); v != "" {
		searchUrl += "language=" + v + "&"
		sc.Data["SearchLanguage"+v] = v
		qry["language"] = v
	}
	sc.Data["URL"] = "/status/list?" + searchUrl

	// Page
	qry["page"] = args.Get("page")
	if qry["page"] == "" {
		qry["page"] = "1"
	}

	solutionModel := model.SolutionModel{}
	qry["module"] = strconv.Itoa(config.ModuleP)
	qry["action"] = "submit"
	count, err := solutionModel.Count(qry)
	if err != nil {
		http.Error(sc.Response, err.Error(), 400)
		return
	}
	var pageCount = (count-1)/config.SolutionPerPage + 1

	page, err := strconv.Atoi(qry["page"])
	if err != nil {
		http.Error(sc.Response, "args error", 400)
		return
	}
	if page > pageCount {
		http.Error(sc.Response, "args error", 400)
		return
	}
	qry["offset"] = strconv.Itoa((page - 1) * config.SolutionPerPage)
	qry["limit"] = strconv.Itoa(config.SolutionPerPage)

	pageData := sc.GetPage(page, pageCount)
	for k, v := range pageData {
		sc.Data[k] = v
	}

	list, err := solutionModel.List(qry)
	if err != nil {
		http.Error(sc.Response, err.Error(), 500)
		return
	}

	sc.Data["Solution"] = list
	sc.Data["Title"] = "Status List"
	sc.Data["IsStatus"] = true
	sc.Data["Privilege"] = sc.Privilege
	sc.Data["Uid"] = sc.Uid

	sc.RenderTemplate("view/layout.tpl", "view/status_list.tpl")
}

func (sc *StatusController) Code() {
	restweb.Logger.Debug("Status Code")

	args := sc.Requset.URL.Query()
	restweb.Logger.Debug(args.Get("sid"))
	sid, err := strconv.Atoi(args.Get("sid"))
	if err != nil {
		http.Error(sc.Response, "args error", 400)
		return
	}

	solutionModel := model.SolutionModel{}
	one, err := solutionModel.Detail(sid)
	if err != nil {
		http.Error(sc.Response, err.Error(), 400)
		return
	}
	if one.Error != "" {
		one.Code = one.Code + "\n/*\n" + one.Error + "*/\n"
	}

	if one.Uid == sc.Uid || sc.Privilege > config.PrivilegePU {
		sc.Data["Solution"] = one
		sc.Data["Title"] = "View Code"
		sc.Data["IsCode"] = true
		sc.RenderTemplate("view/layout.tpl", "view/status_code.tpl")
	} else {
		sc.Err400("Warning", "You can't see it!")
	}
}
