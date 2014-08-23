package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"net/http"
	"strconv"
)

type StatusController struct {
	class.Controller
}

func (this *StatusController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Status List")
	this.Init(w, r)
	args := this.ParseURL(r.URL.String())
	url := "/solution?list"
	searchUrl := ""
	qry := make(map[string]string)
	// Search
	if v, ok := args["uid"]; ok {
		searchUrl += "/uid?" + v
		this.Data["SearchUid"] = v
		qry["uid"] = v
	}
	if v, ok := args["pid"]; ok {
		searchUrl += "/pid?" + v
		this.Data["SearchPid"] = v
		qry["pid"] = v
	}
	if v, ok := args["judge"]; ok {
		searchUrl += "/judge?" + v
		this.Data["SearchJudge"+v] = v
		qry["judge"] = v
	}
	if v, ok := args["language"]; ok {
		searchUrl += "/language?" + v
		this.Data["SearchLanguage"+v] = v
		qry["language"] = v
	}
	url += searchUrl
	this.Data["URL"] = "/status?list" + searchUrl

	// Page
	if _, ok := args["page"]; !ok {
		args["page"] = "1"
	}

	solutionModel := model.SolutionModel{}
	qry["module"] = strconv.Itoa(config.ModuleP)
	qry["action"] = "submit"
	count, err := solutionModel.Count(qry)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var pageCount = (count-1)/config.SolutionPerPage + 1

	page, err := strconv.Atoi(args["page"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	if page > pageCount {
		http.Error(w, "args error", 400)
		return
	}
	qry["offset"] = strconv.Itoa((page - 1) * config.SolutionPerPage)
	qry["limit"] = strconv.Itoa(config.SolutionPerPage)

	pageData := this.GetPage(page, pageCount)
	for k, v := range pageData {
		this.Data[k] = v
	}

	list, err := solutionModel.List(qry)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	this.Data["Solution"] = list
	this.Data["Title"] = "Status List"
	this.Data["IsStatus"] = true
	err = this.Execute(w, "view/layout.tpl", "view/status_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *StatusController) Code(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Status Code")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	class.Logger.Debug(args["sid"])
	sid, err := strconv.Atoi(args["sid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	solutionModel := model.SolutionModel{}
	one, err := solutionModel.Detail(sid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	this.Data["Solution"] = one
	this.Data["Title"] = "View Code"
	this.Data["IsCode"] = true
	err = this.Execute(w, "view/layout.tpl", "view/status_code.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
