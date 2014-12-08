package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"
	"strconv"
)

type Contest struct {
	Cid           int
	ContestDetail *model.Contest
	Index         map[int]int
	class.Controller
}

func (c *Contest) InitContest(Cid string) {
	c.Init()

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		c.Error(err.Error(), 400)
		return
	}
	c.Cid = cid

	contestModel := model.ContestModel{}
	c.ContestDetail, err = contestModel.Detail(cid)
	if err != nil {
		c.Error(err.Error(), 500)
		return
	}

	c.Index = make(map[int]int)
	for k, v := range c.ContestDetail.List {
		c.Index[v] = k
	}
	c.Output["Cid"] = strconv.Itoa(c.Cid)
	c.Output["Title"] = "Contest Detail " + strconv.Itoa(c.Cid)
	c.Output["Contest"] = c.ContestDetail.Title
	c.Output["IsContestDetail"] = true
	c.Output["IsContest"] = true
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

func (c *Contest) Detail(Cid string) {
	restweb.Logger.Debug("Contest Problem List")

	c.InitContest(Cid)
	list := make([]*model.Problem, len(c.ContestDetail.List))
	idx := 0
	for _, v := range c.ContestDetail.List {
		problemModel := model.ProblemModel{}
		one, err := problemModel.Detail(v)
		if err != nil {
			restweb.Logger.Debug(err)
			continue
		}
		one.Pid = idx
		qry := make(map[string]string)
		qry["pid"] = strconv.Itoa(v)
		qry["module"] = strconv.Itoa(config.ModuleC)
		qry["action"] = "accept"
		one.Solve, err = c.GetCount(qry)
		if err != nil {
			restweb.Logger.Debug(err)
			continue
		}
		qry["action"] = "submit"
		one.Submit, err = c.GetCount(qry)
		if err != nil {
			restweb.Logger.Debug(err)
			continue
		}

		list[idx] = one
		idx++
	}

	c.Output["Problem"] = list
	c.Output["IsContestProblem"] = true
	c.Output["Start"] = c.ContestDetail.Start
	c.Output["End"] = c.ContestDetail.End

	c.RenderTemplate("view/layout.tpl", "view/contest/problem_list.tpl")
}
