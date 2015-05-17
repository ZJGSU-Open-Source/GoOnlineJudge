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

//@URL: /api/contests @method: GET
func (c *ContestController) Index() {
	restweb.Logger.Debug("Contest List")

	CModel := model.ContestModel{}
	conetestList, err := CModel.List(nil)
	if err != nil {
		c.Error(err.Error(), 500)
		return
	}

	c.Output["Contest"] = conetestList
	c.RenderJson()

}
