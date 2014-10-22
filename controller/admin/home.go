package admin

import (
	"GoOnlineJudge/class"

	"net/http"
)

type HomeController struct {
	class.Controller
}

func (hc HomeController) Home(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug(r.RemoteAddr + "visit Admin Home")
	hc.Init(w, r)

	hc.Data["Title"] = "Admin - Home"
	hc.Execute(w, "view/admin/layout.tpl", "view/admin/home.tpl")
}
