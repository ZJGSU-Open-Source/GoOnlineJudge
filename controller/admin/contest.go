package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"net/http"
	"strconv"
	"strings"
	"time"
)

//竞赛控件
type ContestController struct {
	Cid           int
	ContestDetail *model.Contest
	Index         map[int]int
	class.Controller
}

func (cc ContestController) Route(w http.ResponseWriter, r *http.Request) {
	cc.Init(w, r)
	action := cc.GetAction(r.URL.Path, 2)
	defer func() {
		if e := recover(); e != nil {
			http.Error(w, "no such page", 404)
		}
	}()
	rv := class.GetReflectValue(w, r)
	class.CallMethod(&cc, strings.Title(action), rv)
}

//列出所有的比赛 url:/admin/contest/list?type=<contest,exercise>
func (cc *ContestController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest List")

	Type := r.URL.Query().Get("type")

	qry := make(map[string]string)
	qry["type"] = Type
	contestModel := model.ContestModel{}
	contestList, err := contestModel.List(qry)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	cc.Data["Contest"] = contestList
	cc.Data["Title"] = "Admin - " + strings.Title(Type) + " List"
	cc.Data["Is"+strings.Title(Type)] = true
	cc.Data["IsList"] = true
	cc.Execute(w, "view/admin/layout.tpl", "view/admin/contest_list.tpl")
}

// 添加比赛页面 url:/admin/contest/add?type=<contest,exercise>
func (cc *ContestController) Add(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Contest Add")

	Type := r.URL.Query().Get("type")
	//class.Logger.Debug(Type)
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

	cc.Data["Title"] = "Admin - " + strings.Title(Type) + " Add"
	cc.Data["Is"+strings.Title(Type)] = true
	cc.Data["IsAdd"] = true
	cc.Data["Type"] = Type

	cc.Execute(w, "view/admin/layout.tpl", "view/admin/contest_add.tpl")
}

// 插入比赛 url:/admin/contest/insert?type=<contest,exercise>
func (cc *ContestController) Insert(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Contest Insert")
	if r.Method != "POST" {
		cc.Err400(w, r, "Error", "Error Method to Insert contest")
		return
	}

	Type := r.URL.Query().Get("type")

	one := model.Contest{}

	one.Title = r.FormValue("title")
	one.Type = Type
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
		http.Error(w, "args error", 400)
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
		http.Error(w, "args error", 400)
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
			class.Logger.Debug(err)
			continue
		}
		problemModel := model.ProblemModel{}
		_, err = problemModel.Detail(pid) //检查题目是否存在
		if err != nil {
			class.Logger.Debug(err)
			continue
		}
		list = append(list, pid)
	}
	one.List = list

	contestModel := model.ContestModel{}
	err = contestModel.Insert(one)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/admin/contest/list?type="+Type, http.StatusFound) //重定向到竞赛列表页
}

//更改contest状态 url:/admin/contest/status/
func (cc *ContestController) Status(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Contest Status")
	if r.Method != "POST" {
		cc.Err400(w, r, "Error", "Error Method to Change contest status")
		return
	}
	cid, err := strconv.Atoi(r.URL.Query().Get("cid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	contestModel := model.ContestModel{}
	one, err := contestModel.Detail(cid)

	Type := one.Type

	var status int
	switch one.Status {
	case config.StatusAvailable:
		status = config.StatusReverse
	default:
		status = config.StatusAvailable
	}

	err = contestModel.Status(cid, status)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/admin/contest/list?type="+Type, http.StatusFound) //重定向到竞赛列表页
}

//删除竞赛 url:/admin/contest/delete/，method:POST
func (cc *ContestController) Delete(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Contest Delete")
	if r.Method != "POST" {
		cc.Err400(w, r, "Error", "Error Method to Delete contest")
		return
	}

	cid, err := strconv.Atoi(r.URL.Query().Get("cid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	contestModel := model.ContestModel{}
	err = contestModel.Delete(cid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(200)
}

// 竞赛编辑页面，url:/admin/contest/edit/
func (cc *ContestController) Edit(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Contest Edit")

	cid, err := strconv.Atoi(r.URL.Query().Get("cid"))
	if err != nil {
		http.Error(w, "args error", 400)
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
		http.Error(w, err.Error(), 400)
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
	Type := one.Type
	cc.Data["Title"] = "Admin - " + strings.Title(Type) + " Edit"
	cc.Data["Is"+strings.Title(Type)] = true
	cc.Data["IsEdit"] = true

	cc.Execute(w, "view/admin/layout.tpl", "view/admin/contest_edit.tpl")
}

// 更新竞赛，url:/admin/contest/update/，method:POST
func (cc *ContestController) Update(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Contest Update")
	if r.Method != "POST" {
		cc.Err400(w, r, "Error", "Error Method to Update contest")
		return
	}

	cid, err := strconv.Atoi(r.URL.Query().Get("cid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	Type := r.URL.Query().Get("type")

	one := model.Contest{}
	one.Title = r.FormValue("title")
	one.Type = Type
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
		http.Error(w, "cc.Query error", 400)
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
		http.Error(w, "args error", 400)
		return
	}
	class.Logger.Debug(one.Argument)
	problemString := r.FormValue("problemList")
	problemString = strings.Trim(problemString, " ")
	problemString = strings.Trim(problemString, ";")
	problemList := strings.Split(problemString, ";")
	var list []int
	for _, v := range problemList {
		pid, err := strconv.Atoi(v)
		if err != nil {
			class.Logger.Debug(err)
			continue
		}
		problemModel := model.ProblemModel{}
		_, err = problemModel.Detail(pid) //检查题目是否存在
		if err != nil {
			class.Logger.Debug(err)
			continue
		}
		list = append(list, pid)
	}
	one.List = list

	contestModel := model.ContestModel{}
	err = contestModel.Update(cid, one)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	http.Redirect(w, r, "/admin/contest/list?type="+Type, http.StatusFound)
}
