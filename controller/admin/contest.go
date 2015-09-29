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
} //@Controller

//列出所有的比赛
//@URL: /admin/contests/ @method:GET
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

// 添加比赛页面
//@URL: /admin/contests/new @method: GET
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

// 插入比赛
//@URL: /contests/ @method:POST
func (cc *AdminContest) Insert() {
	restweb.Logger.Debug("Admin Contest Insert")

	if cc.Uid == "" {
		cc.Redirect("/sess", http.StatusFound) //重定向到竞赛列表页
		return
	}

	one := cc.contest()

	if cc.Privilege == config.PrivilegePU {
		one.Status = config.StatusAvailable
	} else {
		one.Status = config.StatusReverse
	}

	contestModel := model.ContestModel{}
	err := contestModel.Insert(one)
	if err != nil {
		cc.Error(err.Error(), 500)
		return
	}

	if cc.Privilege == config.PrivilegePU {
		cc.Redirect("/contests", http.StatusFound) //重定向到竞赛列表页
		return
	}
	cc.Redirect("/admin/contests", http.StatusFound) //重定向到竞赛列表页
}

//更改contest状态
//@URL:/admin/contests/(\d+)/status/ @method:POST
func (cc *AdminContest) Status(Cid string) {
	restweb.Logger.Debug("Admin Contest Status")

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		cc.Error("args error", 400)
		return
	}
	contestModel := model.ContestModel{}
	one, _ := contestModel.Detail(cid)
	if one.Creator != cc.Uid {
		cc.Error("privilege error", 400)
		return
	}

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

//删除竞赛
//@URL: /admin/contests/(\d+)/delete/ @method:POST
func (cc *AdminContest) Delete(Cid string) {
	restweb.Logger.Debug("Admin Contest Delete")

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		cc.Error("args error", 400)
		return
	}
	contestModel := model.ContestModel{}
	old, _ := contestModel.Detail(cid)
	if old.Creator != cc.Uid {
		cc.Error("privilege error", 400)
		return
	}

	err = contestModel.Delete(cid)
	if err != nil {
		cc.Error(err.Error(), 400)
		return
	}
	cc.W.WriteHeader(200)
}

// 竞赛编辑页面，
//@URL:/admin/contests/(\d+)/ @method:GET
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

// 更新竞赛
//@URL:/admin/contests/(\d+)/ @method:POST
func (cc *AdminContest) Update(Cid string) {
	restweb.Logger.Debug("Admin Contest Update")

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		cc.Error("args error", 400)
		return
	}
	contestModel := model.ContestModel{}
	old, _ := contestModel.Detail(cid)
	if old.Creator != cc.Uid {
		cc.Error("privilege error", 400)
		return
	}

	one := cc.contest()

	err = contestModel.Update(cid, one)
	if err != nil {
		cc.Error(err.Error(), 400)
		return
	}
	cc.Redirect("/admin/contests", http.StatusFound)
}

func (cc *AdminContest) contest() (one model.Contest) {

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
		argument = strings.Trim(argument, "\r\n")
		one.Argument = argument
		restweb.Logger.Debug(one.Argument)
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
	one.Creator = cc.Uid
	restweb.Logger.Debug(one.Creator)
	return one
}
