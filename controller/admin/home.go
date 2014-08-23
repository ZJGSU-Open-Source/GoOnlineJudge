package admin

import (
	"GoOnlineJudge/class"
	"net/http"
)

type HomeController struct {
	class.Controller
}

func (this *HomeController) Home(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug(r.RemoteAddr + "visit Admin Home")
	this.Init(w, r)

	this.Data["Title"] = "Admin - Home"

	err := this.Execute(w, "view/admin/layout.tpl", "view/admin/home.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
