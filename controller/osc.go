package controller

import (
	"GoOnlineJudge/class"

	"net/http"
)

type OSCController struct {
	class.Controller
}

func (oc OSCController) Route(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("OSC Page")
	oc.Init(w, r)

	oc.Data["Title"] = "ZJGSU OSC"
	oc.Data["IsOSC"] = true
	oc.Execute(w, "view/layout.tpl", "view/osc.tpl")
}
