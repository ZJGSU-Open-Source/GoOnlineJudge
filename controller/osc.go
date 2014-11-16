package controller

import (
	"GoOnlineJudge/class"
	"restweb"
)

type OSCController struct {
	class.Controller
}

func (oc OSCController) Index() {
	restweb.Logger.Debug("OSC Page")

	oc.Data["Title"] = "ZJGSU OSC"
	oc.Data["IsOSC"] = true
	oc.RenderTemplate("view/layout.tpl", "view/osc.tpl")
}
