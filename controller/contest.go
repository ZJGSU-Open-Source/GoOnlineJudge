package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/model"
	"restweb"

	"net/http"
)

type ContestController struct {
	class.Controller
	Type string
}

func (c ContestController) Get(w http.ResponseWriter, r *http.Request) {
	restweb.Logger.Debug("Contest List")
	c.Init(w, r)

	Type := r.URL.Query().Get("type")
	qry := make(map[string]string)
	qry["type"] = Type

	CModel := model.ContestModel{}
	conetestList, err := CModel.List(qry)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	c.Data["Contest"] = conetestList
	c.Data["Time"] = restweb.GetTime()
	c.Data["Title"] = "Contest List"
	c.Data["IsContest"] = true
	c.Data["Privilege"] = c.Privilege
	c.Execute(w, "view/layout.tpl", "view/contest_list.tpl")
}
