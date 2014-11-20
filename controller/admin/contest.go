package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"net/http"
	"strconv"
	"strings"
	"time"
)

//竞赛控件
type AdminContest struct {
	Cid           int
	ContestDetail *model.Contest
	Index         map[int]int
	class.Controller
}

//列出所有的比赛 url:/admin/contest/list?type=<contest,exercise>
func (cc *AdminContest) List() {
	restweb.Logger.Debug("Contest List")

	qry := make(map[string]string)
	contestModel := model.ContestModel{}
	contestList, err := contestModel.List(qry)
	if err != nil {
		cc.Error(err.Error(), 400)
	}

	cc.Data["Contest"] = contestList
	cc.Data["Title"] = "Admin - Contest List"
	cc.Data["IsContest"] = true
	cc.Data["IsList"] = true
	cc.RenderTemplate("view/admin/layout.tpl", "view/admin/contest_list.tpl")
}

// 添加比赛页面 url:/admin/contest/add?type=<contest,exercise>
func (cc *AdminContest) Add() {
	restweb.Logger.Debug("Admin Contest Add")

	now := time.Now()
	cc.Data["StartYear"] = now.Year()
	cc.Data["StartMonth"] = int(now.Month())
	cc.Data["StartDay"] = int(now.Day())
	cc.Data["StartHour"] = int(now.Hour())

	end := now.Add(5 * time.Hour)
	cc.Data["EndYear"] = end.Year()
	cc.Data["EndMonth"] = int(end.Month())
	cc.Data["EndDay"] = int(end.Day())
	cc.Data["EndHour"] = int(end.Hour())

	cc.Data["Title"] = "Admin - Contest Add"
	cc.Data["IsContest"] = true
	cc.Data["IsAdd"] = true

	cc.RenderTemplate("view/admin/layout.tpl", "view/admin/contest_add.tpl")
}

// 插入比赛 url:/admin/contest/insert?type=<contest,exercise>
func (cc *AdminContest) Insert() {
	restweb.Logger.Debug("Admin Contest Insert")

	one := model.Contest{}
	r := cc.Requset

	one.Title = r.FormValue("title")
	year, err := strconv.Atoi(r.FormValue("startTimeYear"))
	month, err := strconv.Atoi(r.FormValue("startTimeMonth"))
	day, err := strconv.Atoi(r.FormValue("startTimeDay"))
	hour, err := strconv.Atoi(r.FormValue("startTimeHour"))
	min, err := strconv.Atoi(r.FormValue("startTimeMinute"))
	start := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one.Start = start.Unix()

	year, err = strconv.Atoi(r.FormValue("endTimeYear"))
	month, err = strconv.Atoi(r.FormValue("endTimeMonth"))
	day, err = strconv.Atoi(r.FormValue("endTimeDay"))
	hour, err = strconv.Atoi(r.FormValue("endTimeHour"))
	min, err = strconv.Atoi(r.FormValue("endTimeMinute"))
	end := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one.End = end.Unix()

	if start.After(end) {
		cc.Error("args error", 400)
		return
	}

	switch r.FormValue("type") {
	case "public":
		one.Encrypt = config.EncryptPB
	case "private":
		one.Encrypt = config.EncryptPT
		argument := r.FormValue("userlist")
		var cr rune = 13
		crStr := string(cr)
		argument = strings.Trim(argument, crStr)
		argument = strings.Trim(argument, "/r/n")
		argument = strings.Replace(argument, "/r/n", "", -1)
		argument = strings.Replace(argument, crStr, "/n", -1)
		one.Argument = argument
	case "password":
		one.Encrypt = config.EncryptPW
		one.Argument = r.FormValue("password")
	default:
		cc.Error("args error", 400)
		return
	}

	problemString := r.FormValue("problemList")
	problemString = strings.Trim(problemString, " ")
	problemString = strings.Trim(problemString, ";")
	problemList := strings.Split(problemString, ";")
	var list []int
	for _, v := range problemList {
		pid, err := strconv.Atoi(v)
		if err != nil {
			restweb.Logger.Debug(err)
			continue
		}
		problemModel := model.ProblemModel{}
		_, err = problemModel.Detail(pid) //检查题目是否存在
		if err != nil {
			restweb.Logger.Debug(err)
			continue
		}
		list = append(list, pid)
	}
	one.List = list

	contestModel := model.ContestModel{}
	err = contestModel.Insert(one)
	if err != nil {
		cc.Error(err.Error(), 500)
		return
	}

	cc.Redirect("/admin/contests", http.StatusFound) //重定向到竞赛列表页
}

//更改contest状态 url:/admin/contest/status/
func (cc *AdminContest) Status(Cid string) {
	restweb.Logger.Debug("Admin Contest Status")

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		cc.Error("args error", 400)
		return
	}

	contestModel := model.ContestModel{}
	one, err := contestModel.Detail(cid)

	var status int
	switch one.Status {
	case config.StatusAvailable:
		status = config.StatusReverse
	default:
		status = config.StatusAvailable
	}

	err = contestModel.Status(cid, status)
	if err != nil {
		cc.Error(err.Error(), 500)
		return
	}

	cc.Redirect("/admin/contests", http.StatusFound) //重定向到竞赛列表页
}

//删除竞赛 url:/admin/contest/delete/，method:POST
func (cc *AdminContest) Delete(Cid string) {
	restweb.Logger.Debug("Admin Contest Delete")

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		cc.Error("args error", 400)
		return
	}

	contestModel := model.ContestModel{}
	err = contestModel.Delete(cid)
	if err != nil {
		cc.Error(err.Error(), 400)
		return
	}
	cc.Response.WriteHeader(200)
}

// 竞赛编辑页面，url:/admin/contests/
func (cc *AdminContest) Edit(Cid string) {
	restweb.Logger.Debug("Admin Contest Edit")

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		cc.Error("args error", 400)
		return
	}

	var one struct {
		*model.Contest
		StartTimeYear   int
		StartTimeMonth  int
		StartTimeDay    int
		StartTimeHour   int
		StartTimeMinute int
		EndTimeYear     int
		EndTimeMonth    int
		EndTimeDay      int
		EndTimeHour     int
		EndTimeMinute   int
		ProblemList     string
		IsPublic        bool
		IsPrivate       bool
		IsPassword      bool
	}
	contestModel := model.ContestModel{}
	one.Contest, err = contestModel.Detail(cid)
	if err != nil {
		cc.Error(err.Error(), 400)
		return
	}

	start := time.Unix(one.Start, 0).Local()
	one.StartTimeYear = start.Year()
	one.StartTimeMonth = int(start.Month())
	one.StartTimeDay = start.Day()
	one.StartTimeHour = start.Hour()
	one.StartTimeMinute = start.Minute()

	end := time.Unix(one.End, 0).Local()
	one.EndTimeYear = end.Year()
	one.EndTimeMonth = int(end.Month())
	one.EndTimeDay = end.Day()
	one.EndTimeHour = end.Hour()
	one.EndTimeMinute = end.Minute()
	one.ProblemList = ""
	for _, v := range one.List {
		one.ProblemList += strconv.Itoa(v) + ";"
	}
	one.IsPublic = false
	one.IsPrivate = false
	one.IsPassword = false
	switch one.Encrypt {
	case config.EncryptPB:
		one.IsPublic = true
	case config.EncryptPT:
		one.IsPrivate = true
	case config.EncryptPW:
		one.IsPassword = true
	}

	cc.Data["Detail"] = one
	cc.Data["Title"] = "Admin - " + "Contest" + " Edit"
	cc.Data["IsContest"] = true
	cc.Data["IsEdit"] = true

	cc.RenderTemplate("view/admin/layout.tpl", "view/admin/contest_edit.tpl")
}

// 更新竞赛，url:/admin/contest/update/，method:POST
func (cc *AdminContest) Update(Cid string) {
	restweb.Logger.Debug("Admin Contest Update")

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		cc.Error("args error", 400)
		return
	}

	r := cc.Requset
	one := model.Contest{}
	one.Title = r.FormValue("title")
	year, _ := strconv.Atoi(r.FormValue("startTimeYear"))
	month, _ := strconv.Atoi(r.FormValue("startTimeMonth"))
	day, _ := strconv.Atoi(r.FormValue("startTimeDay"))
	hour, _ := strconv.Atoi(r.FormValue("startTimeHour"))
	min, _ := strconv.Atoi(r.FormValue("startTimeMinute"))
	start := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one.Start = start.Unix()

	year, _ = strconv.Atoi(r.FormValue("endTimeYear"))
	month, _ = strconv.Atoi(r.FormValue("endTimeMonth"))
	day, _ = strconv.Atoi(r.FormValue("endTimeDay"))
	hour, _ = strconv.Atoi(r.FormValue("endTimeHour"))
	min, _ = strconv.Atoi(r.FormValue("endTimeMinute"))
	end := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one.End = end.Unix()

	if start.After(end) {
		cc.Error("cc.Query error", 400)
		return
	}

	switch r.FormValue("type") {
	case "public":
		one.Encrypt = config.EncryptPB
		one.Argument = ""
	case "private":
		one.Encrypt = config.EncryptPT
		argument := r.FormValue("userlist")
		var cr rune = 13
		crStr := string(cr)
		argument = strings.Trim(argument, crStr)
		argument = strings.Trim(argument, "\r\n")
		argument = strings.Replace(argument, "\r\n", "\n", -1)
		argument = strings.Replace(argument, crStr, "\n", -1)
		one.Argument = argument
	case "password":
		one.Encrypt = config.EncryptPW
		one.Argument = r.FormValue("password")
	default:
		cc.Error("args error", 400)
		return
	}
	restweb.Logger.Debug(one.Argument)
	problemString := r.FormValue("problemList")
	problemString = strings.Trim(problemString, " ")
	problemString = strings.Trim(problemString, ";")
	problemList := strings.Split(problemString, ";")
	var list []int
	for _, v := range problemList {
		pid, err := strconv.Atoi(v)
		if err != nil {
			restweb.Logger.Debug(err)
			continue
		}
		problemModel := model.ProblemModel{}
		_, err = problemModel.Detail(pid) //检查题目是否存在
		if err != nil {
			restweb.Logger.Debug(err)
			continue
		}
		list = append(list, pid)
	}
	one.List = list

	contestModel := model.ContestModel{}
	err = contestModel.Update(cid, one)
	if err != nil {
		cc.Error(err.Error(), 400)
		return
	}
	cc.Redirect("/admin/contests", http.StatusFound)
}
