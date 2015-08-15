package handler

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/middleware"
	"GoOnlineJudge/model"

	"github.com/zenazn/goji/web"

	"encoding/json"
	"log"
	"net/http"
	"strconv"
	// "strings"
)

// 列出特定数量的问题?pid=<pid>&titile=<titile>&source=<source>&page=<page>
//@URL:/api/problems @method:GET
func ListProblems(c web.C, w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	qry := make(map[string]string)

	if v := r.FormValue("pid"); len(v) > 0 {
		qry["pid"] = v
	}
	if v := r.FormValue("title"); len(v) > 0 {
		qry["title"] = v
	}
	if v := r.FormValue("source"); len(v) > 0 {
		qry["source"] = v
	}
	if v := r.FormValue("offset"); len(v) > 0 {
		qry["offset"] = v
	}
	if v := r.FormValue("limit"); len(v) > 0 {
		qry["limit"] = v
	}
	if v := r.FormValue("status"); len(v) > 0 {
		qry["status"] = v
	}

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
//@URL: /api/problems/:id @method: GET
func GetProblem(c web.C, w http.ResponseWriter, r *http.Request) {
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
//@URL: /api/problems/:id @method: POST
func PostSolution(c web.C, w http.ResponseWriter, r *http.Request) {

	var (
		Pid  = c.URLParams["pid"]
		user = middleware.ToUser(c)
	)

	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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
	one.Uid = user.Uid
	one.Module = config.ModuleP
	one.Mid = config.ModuleP

	problemModel := model.ProblemModel{}
	pro, err := problemModel.Detail(pid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	one.Code = in.Code
	one.Length = len(in.Code)
	one.Language, _ = strconv.Atoi(r.FormValue("compiler_id"))
	one.Share = user.ShareCode

	hint := make(map[string]string)
	errflag := true
	switch {
	case pro.Pid == 0:
		hint["info"] = "No such problem."
	case in.Code == "":
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
