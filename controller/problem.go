package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"encoding/json"
	"net/http"
	"restweb"
	"strconv"
	// "strings"
)

// 问题控件
type ProblemController struct {
	class.Controller
} //@Controller

// 列出特定数量的问题?pid=<pid>&titile=<titile>&source=<source>&page=<page>
//@URL:/api/problems @method:GET
func (pc *ProblemController) List() {

	restweb.Logger.Debug(pc.R.RemoteAddr + "visit Problem List")

	qry := make(map[string]string)
	in := struct {
		pid    string `json:"pid"`
		title  string `json:"title"`
		source string `json:"source"`
		offset int    `json:"offset"`
		limit  int    `json:"limit"`
		status string `json:"status"`
	}{}

	if err := json.NewDecoder(pc.R.Body).Decode(&in); err != nil {
		pc.Error(err.Error(), http.StatusBadRequest)
		return
	}
	//TODO in->qry
	problemModel := &model.ProblemModel{}

	problemList, err := problemModel.List(qry)
	if err != nil {
		pc.Error("post error", 500)
		return
	}

	pc.Output["Problem"] = problemList
	pc.RenderJson()
}

//列出某问题的详细信息
//@URL: /api/problems/(\d+) @method: GET
func (pc *ProblemController) Detail(Pid string) {
	restweb.Logger.Debug("Problem Detail")

	pid, err := strconv.Atoi(Pid)
	if err != nil {
		pc.Error("args error", 400)
		return
	}

	problemModel := model.ProblemModel{}
	one, err := problemModel.Detail(pid)
	if err != nil {
		pc.Err400("Problem "+Pid, "No such problem")
		return
	}
	pc.Output["Detail"] = one
	pc.RenderJson()
}

//提交某一问题的solution
//@URL: /api/problems/(\d+) @method: POST
func (pc *ProblemController) Submit(Pid string) {

	restweb.Logger.Debug("Problem Submit")

	in := struct {
		Uid        string
		Code       string
		CompilerID int
	}{}

	if err := json.NewDecoder(pc.R.Body).Decode(&in); err != nil {
		pc.Error(err.Error(), http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(Pid)
	if err != nil {
		pc.Error("args error", 400)
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
		pc.Error(err.Error(), 500)
		return
	}
	code := pc.Input.Get("code")

	one.Code = code
	one.Length = pc.GetCodeLen(len(pc.Input.Get("code")))
	one.Language, _ = strconv.Atoi(pc.Input.Get("compiler_id"))
	pc.SetSession("Compiler_id", pc.Input.Get("compiler_id")) //or set cookie?
	userModel := model.UserModel{}
	user, _ := userModel.Detail(pc.Uid)
	one.Share = user.ShareCode

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
		pc.W.WriteHeader(400)
		pc.W.Write(b)
		return
	}

	one.Status = config.StatusAvailable
	one.Judge = config.JudgePD

	solutionModel := model.SolutionModel{}
	sid, err := solutionModel.Insert(one)
	if err != nil {
		pc.Error(err.Error(), 500)
		return
	}

	pc.W.WriteHeader(201)
	go func() { //编译运行solution
		one := make(map[string]interface{})
		one["Sid"] = sid
		one["Pid"] = pro.RPid
		one["OJ"] = pro.ROJ
		one["Rejudge"] = false
		reader, _ := pc.JsonReader(&one)
		restweb.Logger.Debug(reader)
		_, err := http.Post(config.JudgeHost, "application/json", reader)
		if err != nil {
			restweb.Logger.Debug("sid[", sid, "] submit post error")
		}
	}()
}
