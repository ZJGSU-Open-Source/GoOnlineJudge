package admin

import (
	"GoOnlineJudge/classes"
	"html/template"
	"log"
	"net/http"
)

type AdminMenuController struct {
	classes.Controller
}

func (this *AdminMenuController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Menu")
	this.Init()

	t, _ := template.ParseFiles("views/admin/menu.tpl", "views/admin/home.tpl")

	this.Data["Title"] = "Admin Menu"
	t.Execute(w, this.Data)
}
