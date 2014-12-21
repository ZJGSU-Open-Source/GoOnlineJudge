package contest

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type ContestProblem struct {
	Contest
} //@Controller

//@URL: /contests/(\d+)/problems/(\d+) @method:GET
func (pc *ContestProblem) Detail(Cid, Pid string) {
	pc.InitContest(Cid)
	restweb.Logger.Debug("Contest Problem Detail")

	pid, err := strconv.Atoi(Pid)
	if err != nil {
		pc.Error("args error", 400)
		return
	}
	rpid := pc.ContestDetail.List[pid]
	problemModel := model.ProblemModel{}
	one, err := problemModel.Detail(rpid)
	if err != nil {
		pc.Error(err.Error(), 500)
		return
	}
	qry := make(map[string]string)
	qry["pid"] = strconv.Itoa(rpid)
	qry["action"] = "accept"
	one.Solve, err = pc.GetCount(qry)
	if err != nil {
		pc.Error(err.Error(), 500)
		return
	}
	qry["action"] = "submit"
	one.Submit, err = pc.GetCount(qry)
	if err != nil {
		pc.Error(err.Error(), 500)
		return
	}

	pc.Output["Detail"] = one
	pc.Output["Pid"] = pid
	pc.Output["Status"] = pc.ContestDetail.Status

	pc.RenderTemplate("view/layout.tpl", "view/contest/problem_detail.tpl")
}

//@URL: /contests/(\d+)/problems/(\d+) @method: POST
func (pc *ContestProblem) Submit(Cid, Pid string) {
	restweb.Logger.Debug("Contest Problem Submit")
	pc.InitContest(Cid)

	uid := pc.Uid
	if uid == "" {
		pc.Error("user login required", 401)
		return
	}

	pid, err := strconv.Atoi(Pid)
	if err != nil || pid >= len(pc.ContestDetail.List) {
		pc.Error("args error", 400)
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
		pc.Error(err.Error(), 500)
		return
	}

	w := pc.W
	code := pc.Input.Get("code")
	one.Code = code
	one.Length = pc.GetCodeLen(len(pc.Input.Get("code")))
	one.Language, _ = strconv.Atoi(pc.Input.Get("compiler_id"))

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
		pc.Error(err.Error(), 500)
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
		restweb.Logger.Debug(reader)
		W, err := http.Post(config.JudgeHost, "application/json", reader)
		if err != nil {
			pc.Error("post error", 500)
		}
		W.Body.Close()
	}()
}
