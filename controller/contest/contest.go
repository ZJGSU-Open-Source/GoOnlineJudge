package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"net/http"
	"strconv"
)

type Contest struct {
	Cid           int
	ContestDetail *model.Contest
	Index         map[int]int
	class.Controller
}

func (this *Contest) InitContest(w http.ResponseWriter, r *http.Request) {
	this.Init(w, r)

	args := r.URL.Query()

	cid, err := strconv.Atoi(args.Get("cid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	this.Cid = cid

	contestModel := model.ContestModel{}
	this.ContestDetail, err = contestModel.Detail(cid)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	this.Index = make(map[int]int)
	for k, v := range this.ContestDetail.List {
		this.Index[v] = k
	}
	this.Data["Cid"] = strconv.Itoa(this.Cid)
	this.Data["Title"] = "Contest Detail " + strconv.Itoa(this.Cid)
	this.Data["Contest"] = this.ContestDetail.Title
	this.Data["IsContestDetail"] = true
}

func (this *Contest) GetCount(qry map[string]string) (int, error) {
	if qry == nil {
		qry = make(map[string]string)
	}
	qry["module"] = strconv.Itoa(config.ModuleC)
	qry["mid"] = strconv.Itoa(this.Cid)
	solutionModel := model.SolutionModel{}
	count, err := solutionModel.Count(qry)
	if err != nil {
		return 0, err
	}
	return count, nil
}
