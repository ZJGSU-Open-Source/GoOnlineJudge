package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// 问题控件
type ProblemController struct {
	class.Controller
}

func (pc ProblemController) Route(w http.ResponseWriter, r *http.Request) {
	pc.Init(w, r)
	action := pc.GetAction(r.URL.Path, 1)
	defer func() {
		if e := recover(); e != nil {
			http.Error(w, "no such page", 404)
		}
	}()
	rv := class.GetReflectValue(w, r)
	class.CallMethod(&pc, strings.Title(action), rv)
}

// 列出特定数量的问题,URL，/problem/list?pid=<pid>&titile=<titile>&source=<source>&page=<page>
func (pc *ProblemController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug(r.RemoteAddr + "visit Problem List")

	args := r.URL.Query()
	qry := make(map[string]string)
	url := "/problem/list?"

	class.Logger.Debug(r.URL.RequestURI())
	// Search
	if v := args.Get("pid"); v != "" { //按pid查找
		qry["pid"] = v
		url += "pid=" + v + "&"
		pc.Data["SearchPid"] = true
		pc.Data["SearchValue"] = v
	} else if v := args.Get("title"); v != "" { //按问题标题查找
		class.Logger.Debug(v)
		url += "title=" + v + "&"
		pc.Data["SearchTitle"] = true
		pc.Data["SearchValue"] = v
		for _, ep := range "+.?$|*^ " {
			v = strings.Replace(v, string(ep), "\\"+string(ep), -1)
		}
		qry["title"] = v
	} else if v := args.Get("source"); v != "" { //按问题来源查找
		class.Logger.Debug(v)
		url += "source=" + v + "&"
		pc.Data["SearchSource"] = true
		pc.Data["SearchValue"] = v
		for _, ep := range "+.?$|*^ " {
			v = strings.Replace(v, string(ep), "\\"+string(ep), -1)
		}
		qry["source"] = v
	}
	pc.Data["URL"] = url

	// Page
	qry["page"] = args.Get("page")
	if v := qry["page"]; v == "" { //指定页码
		qry["page"] = "1"
	}

	if pc.Privilege <= config.PrivilegePU {
		qry["status"] = "2" //strconv.Itoa(config.StatusAvailable)
	}

	problemModel := model.ProblemModel{}
	count, err := problemModel.Count(qry)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	class.Logger.Debug(count)
	var pageCount = (count-1)/config.ProblemPerPage + 1
	page, err := strconv.Atoi(qry["page"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	if page > pageCount {
		http.Error(w, "args error", 400)
		return
	}

	qry["offset"] = strconv.Itoa((page - 1) * config.ProblemPerPage) //偏移位置
	qry["limit"] = strconv.Itoa(config.ProblemPerPage)               //每页问题数量
	pageData := pc.GetPage(page, pageCount)
	for k, v := range pageData {
		pc.Data[k] = v
	}

	problemList, err := problemModel.List(qry)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	class.Logger.Debug(len(problemList))

	pc.Data["Problem"] = problemList
	pc.Data["Privilege"] = pc.Privilege
	pc.Data["Time"] = pc.GetTime()
	pc.Data["Title"] = "Problem List"
	pc.Data["IsProblem"] = true
	pc.Execute(w, "view/layout.tpl", "view/problem_list.tpl")
}

//列出某问题的详细信息，URL，/probliem/detail?pid=<pid>
func (pc *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Problem Detail")

	args := r.URL.Query()
	pid, err := strconv.Atoi(args.Get("pid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	problemModel := model.ProblemModel{}
	one, err := problemModel.Detail(pid)
	if err != nil {
		pc.Err400(w, r, "Problem "+args.Get("pid"), "No such problem")
		return
	}
	pc.Data["Detail"] = one

	if pc.Privilege <= config.PrivilegePU && one.Status == config.StatusReverse { // 如果问题状态为普通用户不可见
		pc.Err400(w, r, "Problem "+args.Get("pid"), "No such problem")
		return
	}

	pc.Data["Privilege"] = pc.Privilege
	pc.Data["Title"] = "Problem — " + strconv.Itoa(pid)
	pc.Execute(w, "view/layout.tpl", "view/problem_detail.tpl")
}

//提交某一问题的solution， URL /problem/submit?pid=<pid>，method POST
func (pc *ProblemController) Submit(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Problem Submit")

	if r.Method != "POST" { // 要求请求方法为post
		http.Error(w, "method error", 400)
		return
	}

	if pc.Uid == "" { //要求用户登入
		http.Error(w, "user login required", 401)
		return
	}

	args := r.URL.Query()
	pid, err := strconv.Atoi(args.Get("pid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	var one model.Solution
	one.Pid = pid
	one.Uid = pc.Uid
	one.Module = config.ModuleP
	one.Mid = config.ModuleP

	problemModel := model.ProblemModel{}
	pro, err := problemModel.Detail(pid)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	code := r.FormValue("code")

	one.Code = code
	one.Length = pc.GetCodeLen(len(r.FormValue("code")))
	one.Language, _ = strconv.Atoi(r.FormValue("compiler_id"))

	hint := make(map[string]string)
	errflag := true
	switch {
	case pro.Pid == 0:
		hint["info"] = "No such problem."
	case code == "":
		hint["info"] = "Your source code is too short."
	default:
		errflag = false
	}
	if errflag {
		b, _ := json.Marshal(&hint)
		w.WriteHeader(400)
		w.Write(b)
		return
	}

	one.Status = config.StatusAvailable
	one.Judge = config.JudgePD

	solutionModel := model.SolutionModel{}
	sid, err := solutionModel.Insert(one)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)

	go func() { //编译运行solution
		one := make(map[string]interface{})
		one["Sid"] = sid
		one["Time"] = pro.Time
		one["Memory"] = pro.Memory
		one["Rejudge"] = false
		reader, _ := pc.PostReader(&one)
		class.Logger.Debug(reader)
		response, err := http.Post(config.JudgeHost, "application/json", reader)
		if err != nil {
			http.Error(w, "post error", 500)
		}
		response.Body.Close()
	}()
}
