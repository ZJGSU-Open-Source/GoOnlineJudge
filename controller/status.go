package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"net/http"
	"strconv"
	"strings"
)

type StatusController struct {
	class.Controller
}

func (this StatusController) Route(w http.ResponseWriter, r *http.Request) {
	this.Init(w, r)
	action := this.GetAction(r.URL.Path, 1)
	defer func() {
		if e := recover(); e != nil {
			http.Error(w, "no such page", 404)
		}
	}()
	rv := class.GetReflectValue(w, r)
	class.CallMethod(&this, strings.Title(action), rv)
}

func (this *StatusController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Status List")
	this.Init(w, r)
	args := r.URL.Query()
	searchUrl := ""
	qry := make(map[string]string)
	// Search
	if v := args.Get("uid"); v != "" {
		searchUrl += "uid=" + v + "&"
		this.Data["SearchUid"] = v
		qry["uid"] = v
	}
	if v := args.Get("pid"); v != "" {
		searchUrl += "pid=" + v + "&"
		this.Data["SearchPid"] = v
		qry["pid"] = v
	}
	if v := args.Get("judge"); v != "" {
		searchUrl += "judge=" + v + "&"
		this.Data["SearchJudge"+v] = v
		qry["judge"] = v
	}
	if v := args.Get("language"); v != "" {
		searchUrl += "language=" + v + "&"
		this.Data["SearchLanguage"+v] = v
		qry["language"] = v
	}
	this.Data["URL"] = "/status/list?" + searchUrl

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
		http.Error(w, err.Error(), 400)
		return
	}
	var pageCount = (count-1)/config.SolutionPerPage + 1

	page, err := strconv.Atoi(qry["page"])
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
	this.Data["Privilege"] = this.Privilege
	this.Data["Uid"] = this.Uid

	this.Execute(w, "view/layout.tpl", "view/status_list.tpl")
}

func (this *StatusController) Code(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Status Code")
	this.Init(w, r)

	args := r.URL.Query()
	class.Logger.Debug(args.Get("sid"))
	sid, err := strconv.Atoi(args.Get("sid"))
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
	this.Execute(w, "view/layout.tpl", "view/status_code.tpl")
}
