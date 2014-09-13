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

func (this ContestController) Route(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest List")
	this.Init(w, r)

	Type := r.URL.Query().Get("type")
	qry := make(map[string]string)
	qry["type"] = Type

	CModel := model.ContestModel{}
	conetestList, err := CModel.List(qry)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	this.Data["Contest"] = conetestList
	this.Data["Time"] = this.GetTime()
	this.Data["Type"] = Type
	this.Data["Title"] = strings.Title(Type) + " List"
	this.Data["Is"+strings.Title(Type)] = true
	this.Data["Privilege"] = this.Privilege
	this.Execute(w, "view/layout.tpl", "view/contest_list.tpl")
}
