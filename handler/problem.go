package handler

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"github.com/zenazn/goji/web"

	"encoding/json"
	"log"
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
func ListProblems(c web.C, w http.ResponseWriter, r *http.Request) {
	restweb.Logger.Debug(r.RemoteAddr + "visit Problem List")

	qry := make(map[string]string)
	// in := struct {
	//     pid    string `json:"pid"`
	//     title  string `json:"title"`
	//     source string `json:"source"`
	//     offset int    `json:"offset"`
	//     limit  int    `json:"limit"`
	//     status string `json:"status"`
	// }{}

	// if err := json.NewDecoder(pc.R.Body).Decode(&in); err != nil {
	//     pc.Error(err.Error(), http.StatusBadRequest)
	//     return
	// }
	//TODO in->qry
	problemModel := &model.ProblemModel{}

	problemList, err := problemModel.List(qry)
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)

		return
	}

	json.NewEncoder(w).Encode(problemList)

}

//列出某问题的详细信息
//@URL: /api/problems/(\d+) @method: GET
func GetProblem(c web.C, w http.ResponseWriter, r *http.Request) {
	restweb.Logger.Debug("Problem Detail")

	var (
		Pid = c.URLParams["pid"]
	)

	pid, err := strconv.Atoi(Pid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	problemModel := model.ProblemModel{}
	one, err := problemModel.Detail(pid)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(one)

}

//提交某一问题的solution
//@URL: /api/problems/(\d+) @method: POST
func Submit(c web.C, w http.ResponseWriter, r *http.Request) {
	restweb.Logger.Debug("Problem Submit")

	var (
		Pid  = c.URLParams["pid"]
		user *model.User
	)

	in := struct {
		Uid        string
		Code       string
		CompilerID int
	}{}

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(Pid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var one model.Solution
	one.Pid = pid
	// one.Uid = pc.Uid
	one.Module = config.ModuleP
	one.Mid = config.ModuleP

	problemModel := model.ProblemModel{}
	pro, err := problemModel.Detail(pid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	code := r.FormValue("code")

	one.Code = code
	one.Length = len(code)
	one.Language, _ = strconv.Atoi(r.FormValue("compiler_id"))
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
		w.WriteHeader(400)
		w.Write(b)
		return
	}

	one.Status = config.StatusAvailable
	one.Judge = config.JudgePD

	solutionModel := model.SolutionModel{}
	sid, err := solutionModel.Insert(one)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	go func() { //编译运行solution
		one := make(map[string]interface{})
		one["Sid"] = sid
		one["Pid"] = pro.RPid
		one["OJ"] = pro.ROJ
		one["Rejudge"] = false

		json.NewEncoder(w).Encode(one)
	}()
}
