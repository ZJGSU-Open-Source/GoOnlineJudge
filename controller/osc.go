package controller

import (
	"GoOnlineJudge/class"
	"net/http"
)

type OSCController struct {
	class.Controller
}

func (this OSCController) Route(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("OSC Page")
	this.Init(w, r)

	this.Data["Title"] = "ZJGSU OSC"
	this.Data["IsOSC"] = true
	this.Execute(w, "view/layout.tpl", "view/osc.tpl")
}
