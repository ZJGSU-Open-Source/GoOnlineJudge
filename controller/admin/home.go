package admin

import (
	"GoOnlineJudge/class"

	"restweb"
)

type AdminHome struct {
	class.Controller
}

func (hc AdminHome) Home() {
	restweb.Logger.Debug("Admin Home")

	hc.Data["IsHome"] = true
	hc.Data["Title"] = "Admin - Home"
	hc.RenderTemplate("view/admin/layout.tpl", "view/admin/home.tpl")
}
