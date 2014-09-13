package controller

import (
	"GoOnlineJudge/class"
	"net/http"
)

type RecruitController struct {
	class.Controller
}

func (this RecruitController) Route(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Recruit Page")
	this.Init(w, r)

	this.Data["Title"] = "Recruit"
	this.Data["IsRecruit"] = true
	err := this.Execute(w, "view/layout.tpl", "view/recruit.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
	}
}
