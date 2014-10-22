package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/model"

	"net/http"
	"strings"
)

type ContestController struct {
	class.Controller
	Type string
}

func (c ContestController) Route(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest List")
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
	c.Data["Time"] = c.GetTime()
	c.Data["Type"] = Type
	c.Data["Title"] = strings.Title(Type) + " List"
	c.Data["Is"+strings.Title(Type)] = true
	c.Data["Privilege"] = c.Privilege
	c.Execute(w, "view/layout.tpl", "view/contest_list.tpl")
}
