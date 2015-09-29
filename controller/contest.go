package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/model"
	"restweb"

	"time"
)

type ContestController struct {
	class.Controller
	Type string
} //@Controller

//@URL: /contests @method: GET
func (c *ContestController) Index() {
	restweb.Logger.Debug("Contest List")

	CModel := model.ContestModel{}
	conetestList, err := CModel.List(nil)
	if err != nil {
		c.Error(err.Error(), 500)
		return
	}

	c.Output["Contest"] = conetestList
	c.Output["Time"] = restweb.GetTime()
	c.Output["Title"] = "Contest List"
	c.Output["IsContest"] = true
	c.Output["Privilege"] = c.Privilege

	now := time.Now()
	c.Output["StartYear"] = now.Year()
	c.Output["StartMonth"] = int(now.Month())
	c.Output["StartDay"] = int(now.Day())
	c.Output["StartHour"] = int(now.Hour())

	end := now.Add(5 * time.Hour)
	c.Output["EndYear"] = end.Year()
	c.Output["EndMonth"] = int(end.Month())
	c.Output["EndDay"] = int(end.Day())
	c.Output["EndHour"] = int(end.Hour())

	c.RenderTemplate("view/layout.tpl", "view/contest_list.tpl")
}
