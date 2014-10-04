package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ProblemController struct {
	Contest
}

func (this ProblemController) Route(w http.ResponseWriter, r *http.Request) {
	this.InitContest(w, r)

	action := this.GetAction(r.URL.Path, 2)
	defer func() {
		if e := recover(); e != nil {
			http.Error(w, "no such page", 404)
		}
	}()
	rv := class.GetReflectValue(w, r)
	class.CallMethod(&this, strings.Title(action), rv)

}

func (this *ProblemController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest Problem List")

	list := make([]*model.Problem, len(this.ContestDetail.List))
	for k, v := range this.ContestDetail.List {
		problemModel := model.ProblemModel{}
		one, err := problemModel.Detail(v)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		one.Pid = k
		qry := make(map[string]string)
		qry["pid"] = strconv.Itoa(v)
		qry["action"] = "accept"
		one.Solve, err = this.GetCount(qry)
		if err != nil {
			http.Error(w, "count error", 500)
			return
		}
		qry["action"] = "submit"
		one.Submit, err = this.GetCount(qry)
		if err != nil {
			http.Error(w, "count error", 500)
			return
		}

		list[k] = one
	}

	this.Data["Problem"] = list
	this.Data["IsContestProblem"] = true
	this.Data["Start"] = this.ContestDetail.Start
	this.Data["End"] = this.ContestDetail.End

	this.Execute(w, "view/layout.tpl", "view/contest/problem_list.tpl")
}

func (this *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest Problem Detail")

	args := r.URL.Query()
	pid, err := strconv.Atoi(args.Get("pid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	problemModel := model.ProblemModel{}
	one, err := problemModel.Detail(this.ContestDetail.List[pid])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	this.Data["Detail"] = one
	this.Data["Pid"] = pid
	this.Data["Status"] = this.ContestDetail.Status

	this.Execute(w, "view/layout.tpl", "view/contest/problem_detail.tpl")
}

func (this *ProblemController) Submit(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest Problem Submit")

	args := r.URL.Query()

	pid, err := strconv.Atoi(args.Get("pid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	pid = this.ContestDetail.List[pid] //get real pid
	uid := this.Uid
	if uid == "" {
		http.Error(w, "user login required", 401)
	}

	one := model.Solution{}
	one.Pid = pid
	one.Uid = uid
	one.Mid = this.ContestDetail.Cid
	one.Module = config.ModuleC

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

	hint := make(map[string]string)
	errflag := true
	switch {
	case pro.Pid == 0 || (pro.Status == config.StatusReverse && this.Privilege <= config.PrivilegePU):
		hint["info"] = "No such problem"
	case code == "":
		hint["info"] = "Your source code is too short"
	case time.Now().Unix() > this.ContestDetail.End:
		hint["info"] = "The contest has ended"
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

	solutionModle := model.SolutionModel{}
	sid, err := solutionModle.Insert(one)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)

	go func() {
		one := make(map[string]string)
		one["Sid"] = strconv.Itoa(sid)
		one["Time"] = strconv.Itoa(pro.Time)
		one["Memory"] = strconv.Itoa(pro.Memory)
		one["Rejudge"] = "false"
		reader, _ := this.PostReader(&one)
		response, err := http.Post(config.JudgeHost, "application/json", reader)
		if err != nil {
			http.Error(w, "post error", 500)
		}
		response.Body.Close()
	}()
}
