package admin

import (
	"GoOnlineJudge/classes"
	"html/template"
	"log"
	"net/http"
)

type NewsController struct {
	classes.Controller
}

func (this *NewsController) Add(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Menu")
	this.Init(w, r)

	t, _ := template.ParseFiles("views/admin/menu.tpl", "views/admin/newsadd.tpl")

	this.Data["Title"] = "Admin Notice Add"
	t.Execute(w, this.Data)
}

func (this *NewsController) Edit(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Menu")
	this.Init(w, r)

	t, _ := template.ParseFiles("views/admin/menu.tpl", "views/admin/newsedit.tpl")

	this.Data["Title"] = "Admin Notice Edit"
	t.Execute(w, this.Data)
}
