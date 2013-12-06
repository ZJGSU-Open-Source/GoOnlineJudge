package admin

import (
	"GoOnlineJudge/classes"
	"html/template"
	"log"
	"net/http"
)

type MenuController struct {
	classes.Controller
}

func (this *MenuController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Menu")
	this.Init(w, r)

	t, _ := template.ParseFiles("views/admin/menu.tpl", "views/admin/home.tpl")

	this.Data["Title"] = "Admin Menu"
	t.Execute(w, this.Data)
}
