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

func (this *ContestController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest List")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())

	Type := args["type"]
	CModel := model.ContestModel{}
	conetestList, err := CModel.List(args)
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
	err = this.Execute(w, "view/layout.tpl", "view/contest_list.tpl")
	if err != nil {
		class.Logger.Debug(err)
		http.Error(w, "tpl error", 500)
		return
	}
}
