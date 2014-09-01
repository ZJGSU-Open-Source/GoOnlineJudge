package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"html/template"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

type ProblemController struct {
	class.Controller
}

func (this *ProblemController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug(r.RemoteAddr + "visit Problem List")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	url := "/problem?list"

	// Search
	if v, ok := args["pid"]; ok {
		url += "/pid?" + v
		this.Data["SearchPid"] = true
		this.Data["SearchValue"] = v
	}
	if v, ok := args["title"]; ok {
		url += "/title?" + v
		this.Data["SearchTitle"] = true
		this.Data["SearchValue"] = v
	}
	if v, ok := args["source"]; ok {
		v = strings.Replace(v, "%20", " ", -1)
		args["source"] = v
		url += "/source?" + v
		this.Data["SearchSource"] = true
		this.Data["SearchValue"] = v
	}
	this.Data["URL"] = url

	// Page
	if _, ok := args["page"]; !ok {
		args["page"] = "1"
	}

	problemModel := model.ProblemModel{}
	count, err := problemModel.Count(args)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var pageCount = (count-1)/config.ProblemPerPage + 1
	page, err := strconv.Atoi(args["page"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	if page > pageCount {
		http.Error(w, "args error", 400)
		return
	}

	args["offset"] = strconv.Itoa((page - 1) * config.ProblemPerPage)
	args["limit"] = strconv.Itoa(config.ProblemPerPage)
	pageData := this.GetPage(page, pageCount)
	for k, v := range pageData {
		this.Data[k] = v
	}

	problemList, err := problemModel.List(args)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

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

func (this *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Problem Detail")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	problemModel := model.ProblemModel{}
	one, err := problemModel.Detail(pid)
	if err != nil {
		//http.Error(w, err.Error(), 500)
		t := template.New("layout.tpl")
		t, err = t.ParseFiles("view/layout.tpl", "view/400.tpl")
		if err != nil {
			class.Logger.Debug(err)
			http.Error(w, "tpl error", 500)
			return
		}

		this.Data["Info"] = "No such problem"
		this.Data["Title"] = "No such problem"
		err = t.Execute(w, this.Data)
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
		return
	}
	this.Data["Detail"] = one

	if this.Privilege <= config.PrivilegePU && one.Status == config.StatusReverse {
		t := template.New("layout.tpl")
		t, err = t.ParseFiles("view/layout.tpl", "view/400.tpl")
		if err != nil {
			class.Logger.Debug(err)
			http.Error(w, "tpl error", 500)
			return
		}

		this.Data["Info"] = "No such problem"
		this.Data["Title"] = "No such problem"
		err = t.Execute(w, this.Data)
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
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

// URL /problem?submit/pid?<pid>
func (this *ProblemController) Submit(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Problem Submit")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	uid := this.Uid
	if uid == "" {
		http.Error(w, "need sign in", 401)
		return
	}

	var one model.Solution
	one.Pid = pid
	one.Uid = uid
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
	one.Length = this.GetCodeLen(len(r.FormValue("code")))
	one.Language, _ = strconv.Atoi(r.FormValue("compiler_id"))

	if code == "" || pro.Pid == 0 || (pro.Status == config.StatusReverse && this.Privilege <= config.PrivilegePU) {
		switch {
		case pro.Pid == 0 || (pro.Status == config.StatusReverse && this.Privilege <= config.PrivilegePU):
			this.Data["Info"] = "No such problem"
		case code == "":
			this.Data["Info"] = "Your source code is too short"
		}
		this.Data["Title"] = "Problem — " + strconv.Itoa(pid)
		err = this.Execute(w, "view/layout.tpl", "view/400.tpl")
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
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

	go func() {
		cmd := exec.Command("./RunServer", "-sid", strconv.Itoa(sid), "-time", strconv.Itoa(pro.Time), "-memory", strconv.Itoa(pro.Memory), "-rejudge", strconv.Itoa(0)) //Run Judge
		err = cmd.Run()
		if err != nil {
			class.Logger.Debug(err)
		}
	}()
}
