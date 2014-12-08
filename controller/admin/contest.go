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

	cc.Output["Contest"] = contestList
	cc.Output["Title"] = "Admin - Contest List"
	cc.Output["IsContest"] = true
	cc.Output["IsList"] = true
	cc.RenderTemplate("view/admin/layout.tpl", "view/admin/contest_list.tpl")
}

// 添加比赛页面 url:/admin/contest/add?type=<contest,exercise>
func (cc *AdminContest) Add() {
	restweb.Logger.Debug("Admin Contest Add")

	now := time.Now()
	cc.Output["StartYear"] = now.Year()
	cc.Output["StartMonth"] = int(now.Month())
	cc.Output["StartDay"] = int(now.Day())
	cc.Output["StartHour"] = int(now.Hour())

	end := now.Add(5 * time.Hour)
	cc.Output["EndYear"] = end.Year()
	cc.Output["EndMonth"] = int(end.Month())
	cc.Output["EndDay"] = int(end.Day())
	cc.Output["EndHour"] = int(end.Hour())

	cc.Output["Title"] = "Admin - Contest Add"
	cc.Output["IsContest"] = true
	cc.Output["IsAdd"] = true

	cc.RenderTemplate("view/admin/layout.tpl", "view/admin/contest_add.tpl")
}

// 插入比赛 url:/admin/contest/insert?type=<contest,exercise>
func (cc *AdminContest) Insert() {
	restweb.Logger.Debug("Admin Contest Insert")

	one := model.Contest{}

	one.Title = cc.Input.Get("title")
	year, err := strconv.Atoi(cc.Input.Get("startTimeYear"))
	month, err := strconv.Atoi(cc.Input.Get("startTimeMonth"))
	day, err := strconv.Atoi(cc.Input.Get("startTimeDay"))
	hour, err := strconv.Atoi(cc.Input.Get("startTimeHour"))
	min, err := strconv.Atoi(cc.Input.Get("startTimeMinute"))
	start := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one.Start = start.Unix()

	year, err = strconv.Atoi(cc.Input.Get("endTimeYear"))
	month, err = strconv.Atoi(cc.Input.Get("endTimeMonth"))
	day, err = strconv.Atoi(cc.Input.Get("endTimeDay"))
	hour, err = strconv.Atoi(cc.Input.Get("endTimeHour"))
	min, err = strconv.Atoi(cc.Input.Get("endTimeMinute"))
	end := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one.End = end.Unix()

	if start.After(end) {
		cc.Error("args error", 400)
		return
	}

	switch cc.Input.Get("type") {
	case "public":
		one.Encrypt = config.EncryptPB
	case "private": //TODO 设置argument为一个string数组
		one.Encrypt = config.EncryptPT
		argument := cc.Input.Get("userlist")
		var cr rune = 13
		crStr := string(cr)
		argument = strings.Trim(argument, crStr)
		argument = strings.Trim(argument, "/r/n")
		argument = strings.Replace(argument, "/r/n", "", -1)
		argument = strings.Replace(argument, crStr, "/n", -1)
		one.Argument = argument
	case "password":
		one.Encrypt = config.EncryptPW
		one.Argument = cc.Input.Get("password")
	default:
		cc.Error("args error", 400)
		return
	}

	problemString := cc.Input.Get("problemList")
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

	cc.Output["Detail"] = one
	cc.Output["Title"] = "Admin - " + "Contest" + " Edit"
	cc.Output["IsContest"] = true
	cc.Output["IsEdit"] = true

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

	one := model.Contest{}
	one.Title = cc.Input.Get("title")
	year, _ := strconv.Atoi(cc.Input.Get("startTimeYear"))
	month, _ := strconv.Atoi(cc.Input.Get("startTimeMonth"))
	day, _ := strconv.Atoi(cc.Input.Get("startTimeDay"))
	hour, _ := strconv.Atoi(cc.Input.Get("startTimeHour"))
	min, _ := strconv.Atoi(cc.Input.Get("startTimeMinute"))
	start := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one.Start = start.Unix()

	year, _ = strconv.Atoi(cc.Input.Get("endTimeYear"))
	month, _ = strconv.Atoi(cc.Input.Get("endTimeMonth"))
	day, _ = strconv.Atoi(cc.Input.Get("endTimeDay"))
	hour, _ = strconv.Atoi(cc.Input.Get("endTimeHour"))
	min, _ = strconv.Atoi(cc.Input.Get("endTimeMinute"))
	end := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one.End = end.Unix()

	if start.After(end) {
		cc.Error("cc.Query error", 400)
		return
	}

	switch cc.Input.Get("type") {
	case "public":
		one.Encrypt = config.EncryptPB
		one.Argument = ""
	case "private":
		one.Encrypt = config.EncryptPT
		argument := cc.Input.Get("userlist")
		var cr rune = 13
		crStr := string(cr)
		argument = strings.Trim(argument, crStr)
		argument = strings.Trim(argument, "\r\n")
		argument = strings.Replace(argument, "\r\n", "\n", -1)
		argument = strings.Replace(argument, crStr, "\n", -1)
		one.Argument = argument
	case "password":
		one.Encrypt = config.EncryptPW
		one.Argument = cc.Input.Get("password")
	default:
		cc.Error("args error", 400)
		return
	}
	restweb.Logger.Debug(one.Argument)
	problemString := cc.Input.Get("problemList")
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
