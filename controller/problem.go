package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"encoding/json"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

// 问题控件
type ProblemController struct {
	class.Controller
}

func (this ProblemController) Route(w http.ResponseWriter, r *http.Request) {
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

// 列出特定数量的问题,URL，/problem?list/pid?<pid>/titile?<titile>/source?<source>/page?<page>
func (this *ProblemController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug(r.RemoteAddr + "visit Problem List")

	args := r.URL.Query()
	//args := this.ParseURL(r.URL.String())
	qry := make(map[string]string)
	url := "/problem/list?"

	class.Logger.Debug(r.URL.RequestURI())
	// Search
	if v := args.Get("pid"); v != "" { //按pid查找
		qry["pid"] = v
		url += "pid=" + v + "&"
		this.Data["SearchPid"] = true
		this.Data["SearchValue"] = v
	} else if v := args.Get("title"); v != "" { //按问题标题查找
		class.Logger.Debug(v)
		url += "title=" + v + "&"
		this.Data["SearchTitle"] = true
		this.Data["SearchValue"] = v
		for _, ep := range "+.?$|*^" {
			v = strings.Replace(v, string(ep), "\\"+string(ep), -1)
		}
		qry["title"] = v
	} else if v := args.Get("source"); v != "" { //按问题来源查找
		class.Logger.Debug(v)
		url += "source=" + v + "&"
		this.Data["SearchSource"] = true
		this.Data["SearchValue"] = v
		for _, ep := range "+.?$|*^ " {
			v = strings.Replace(v, string(ep), "\\"+string(ep), -1)
		}
		qry["source"] = v
	}
	this.Data["URL"] = url

	// Page
	qry["page"] = args.Get("page")
	if v := qry["page"]; v == "" { //指定页码
		qry["page"] = "1"
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
	pageData := this.GetPage(page, pageCount)
	for k, v := range pageData {
		this.Data[k] = v
	}

	problemList, err := problemModel.List(qry)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	class.Logger.Debug(len(problemList))

	this.Data["Problem"] = problemList
	this.Data["Privilege"] = this.Privilege
	this.Data["Time"] = this.GetTime()
	this.Data["Title"] = "Problem List"
	this.Data["IsProblem"] = true
	err = this.Execute(w, "view/layout.tpl", "view/problem_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

//列出某问题的详细信息，URL，/probliem/detail?pid=<pid>
func (this *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
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
		this.Err400(w, r, "Problem "+args.Get("pid"), "No such problem")
		return
	}
	this.Data["Detail"] = one

	if this.Privilege <= config.PrivilegePU && one.Status == config.StatusReverse { // 如果问题状态为普通用户不可见
		this.Err400(w, r, "Problem "+args.Get("pid"), "No such problem")
		return
	}

	this.Data["Privilege"] = this.Privilege
	this.Data["Title"] = "Problem — " + strconv.Itoa(pid)
	err = this.Execute(w, "view/layout.tpl", "view/problem_detail.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

//提交某一问题的solution， URL /problem?submit/pid?<pid>，method POST
func (this *ProblemController) Submit(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Problem Submit")

	if r.Method != "POST" { // 要求请求方法为post
		http.Error(w, "method error", 400)
		return
	}

	if this.Uid == "" { //要求用户登入
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
	one.Uid = this.Uid
	one.Module = config.ModuleP
	one.Mid = config.ModuleP

	problemModel := model.ProblemModel{}
	pro, err := problemModel.Detail(pid)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	code := r.FormValue("code")

	class.Logger.Debug(code)

	one.Code = code
	one.Length = this.GetCodeLen(len(r.FormValue("code")))
	one.Language, _ = strconv.Atoi(r.FormValue("compiler_id"))

	hint := make(map[string]string)
	errflag := true
	switch {
	case pro.Pid == 0 || (pro.Status == config.StatusReverse && this.Privilege <= config.PrivilegePU):
		hint["info"] = "No such problem"
	case code == "":
		hint["info"] = "Your source code is too short"
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
		cmd := exec.Command("./RunServer", "-sid", strconv.Itoa(sid), "-time", strconv.Itoa(pro.Time), "-memory", strconv.Itoa(pro.Memory), "-rejudge", strconv.Itoa(0)) //Run Judge
		err = cmd.Run()
		if err != nil {
			class.Logger.Debug(err)
		}
	}()
}
