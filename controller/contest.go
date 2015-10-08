package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/model"
	"restweb"
)

type ContestController struct {
	class.Controller
	Type string
} //@Controller

//@URL: /contests @method: GET
func (c *ContestController) Index() {
	restweb.Logger.Debug("Contest List")

	var contestList []*model.Contest

	req, _ := apiClient.NewRequest("GET", "/contests", "", nil)
	apiClient.Do(req, &contestList)

	c.Output["Contest"] = contestList
	c.Output["Time"] = restweb.GetTime()
	c.Output["Title"] = "Contest List"
	c.Output["IsContest"] = true
	c.Output["Privilege"] = c.Privilege
	c.RenderTemplate("view/layout.tpl", "view/contest_list.tpl")
}
