package admin

import (
	"GoOnlineJudge/class"
	"html/template"
	"net/http"
)

type HomeController struct {
	class.Controller
}

func (this *HomeController) Home(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug(r.RemoteAddr + "visit Admin Home")
	this.Init(w, r)

	var err error
	t := template.New("layout.tpl")
	t, err = t.ParseFiles("view/admin/layout.tpl", "view/admin/home.tpl")
	if err != nil {
		class.Logger.Debug(err)
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Admin - Home"

	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
