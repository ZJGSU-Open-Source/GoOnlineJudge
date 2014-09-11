package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

type ProblemController struct {
	Contest
}

func (this *ProblemController) Route(w http.ResponseWriter, r *http.Request) {
	this.InitContest(w, r)

	action := this.GetAction(r.URL.Path, 2)
	switch action {
	case "list":
		this.List(w, r)
	case "detail":
		this.Detail(w, r)
	case "submit":
		this.Submit(w, r)
	default:
		http.Error(w, "no such page", 404)
	}

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

	err := this.Execute(w, "view/layout.tpl", "view/contest/problem_list.tpl")
	if err != nil {
		class.Logger.Debug(err)
		http.Error(w, "tpl error", 500)
		return
	}
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

	err = this.Execute(w, "view/layout.tpl", "view/contest/problem_detail.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
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
		w.WriteHeader(401)
		return
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

	info := ""
	errflag := true
	switch {
	case pro.Pid == 0 || (pro.Status == config.StatusReverse && this.Privilege <= config.PrivilegePU):
		info = "No such problem"
	case code == "":
		info = "Your source code is too short"
	case time.Now().Unix() > this.ContestDetail.End:
		info = "The contest has ended"
	default:
		errflag = false
	}
	if errflag {
		this.Err400(w, r, "Problem — "+strconv.Itoa(pid), info)
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
		class.Logger.Debug(sid)
		cmd := exec.Command("./RunServer", "-sid", strconv.Itoa(sid), "-time", strconv.Itoa(pro.Time), "-memory", strconv.Itoa(pro.Memory)) //Run Judge
		err = cmd.Run()
		if err != nil {
			class.Logger.Debug(err)
		}
	}()
}
