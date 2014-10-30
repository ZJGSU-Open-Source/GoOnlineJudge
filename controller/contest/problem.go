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

func (pc ProblemController) Route(w http.ResponseWriter, r *http.Request) {
	pc.InitContest(w, r)

	action := pc.GetAction(r.URL.Path, 2)
	defer func() {
		if e := recover(); e != nil {
			http.Error(w, "no such page", 404)
		}
	}()
	rv := class.GetReflectValue(w, r)
	class.CallMethod(&pc, strings.Title(action), rv)

}

func (pc *ProblemController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest Problem List")

	list := make([]*model.Problem, len(pc.ContestDetail.List))

	idx := 0
	for _, v := range pc.ContestDetail.List {
		problemModel := model.ProblemModel{}
		one, err := problemModel.Detail(v)
		if err != nil {
			class.Logger.Debug(err)
			continue
		}
		one.Pid = idx
		qry := make(map[string]string)
		qry["pid"] = strconv.Itoa(v)
		qry["module"] = strconv.Itoa(config.ModuleC)
		qry["action"] = "accept"
		one.Solve, err = pc.GetCount(qry)
		if err != nil {
			class.Logger.Debug(err)
			continue
		}
		qry["action"] = "submit"
		one.Submit, err = pc.GetCount(qry)
		if err != nil {
			class.Logger.Debug(err)
			continue
		}

		list[idx] = one
		idx++
	}

	pc.Data["Problem"] = list
	pc.Data["IsContestProblem"] = true
	pc.Data["Start"] = pc.ContestDetail.Start
	pc.Data["End"] = pc.ContestDetail.End

	pc.Execute(w, "view/layout.tpl", "view/contest/problem_list.tpl")
}

func (pc *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest Problem Detail")

	args := r.URL.Query()
	pid, err := strconv.Atoi(args.Get("pid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	problemModel := model.ProblemModel{}
	one, err := problemModel.Detail(pc.ContestDetail.List[pid])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	qry := make(map[string]string)
	qry["pid"] = strconv.Itoa(pid)
	qry["action"] = "accept"
	one.Solve, err = pc.GetCount(qry)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	qry["action"] = "submit"
	one.Submit, err = pc.GetCount(qry)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	pc.Data["Detail"] = one
	pc.Data["Pid"] = pid
	pc.Data["Status"] = pc.ContestDetail.Status

	pc.Execute(w, "view/layout.tpl", "view/contest/problem_detail.tpl")
}

func (pc *ProblemController) Submit(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest Problem Submit")

	uid := pc.Uid
	if uid == "" {
		http.Error(w, "user login required", 401)
		return
	}

	args := r.URL.Query()

	pid, err := strconv.Atoi(args.Get("pid"))
	if err != nil || pid >= len(pc.ContestDetail.List) {
		http.Error(w, "args error", 400)
		return
	}

	pid = pc.ContestDetail.List[pid] //get real pid

	one := model.Solution{}
	one.Pid = pid
	one.Uid = uid
	one.Mid = pc.ContestDetail.Cid
	one.Module = config.ModuleC

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
		hint["info"] = "No such problem"
	case code == "":
		hint["info"] = "Your source code is too short"
	case time.Now().Unix() > pc.ContestDetail.End:
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
