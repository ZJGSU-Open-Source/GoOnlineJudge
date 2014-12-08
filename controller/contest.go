package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/model"
	"restweb"
)

type ContestController struct {
	class.Controller
	Type string
}

func (c ContestController) Index() {
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
	c.RenderTemplate("view/layout.tpl", "view/contest_list.tpl")
}
