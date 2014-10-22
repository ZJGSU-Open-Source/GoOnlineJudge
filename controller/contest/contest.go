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

func (c *Contest) InitContest(w http.ResponseWriter, r *http.Request) {
	c.Init(w, r)

	args := r.URL.Query()

	cid, err := strconv.Atoi(args.Get("cid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	c.Cid = cid

	contestModel := model.ContestModel{}
	c.ContestDetail, err = contestModel.Detail(cid)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	c.Index = make(map[int]int)
	for k, v := range c.ContestDetail.List {
		c.Index[v] = k
	}
	c.Data["Cid"] = strconv.Itoa(c.Cid)
	c.Data["Title"] = "Contest Detail " + strconv.Itoa(c.Cid)
	c.Data["Contest"] = c.ContestDetail.Title
	c.Data["IsContestDetail"] = true
	c.Data["IsContest"] = true
}

func (c *Contest) GetCount(qry map[string]string) (int, error) {
	if qry == nil {
		qry = make(map[string]string)
	}
	qry["module"] = strconv.Itoa(config.ModuleC)
	qry["mid"] = strconv.Itoa(c.Cid)
	solutionModel := model.SolutionModel{}
	count, err := solutionModel.Count(qry)
	if err != nil {
		return 0, err
	}
	return count, nil
}
