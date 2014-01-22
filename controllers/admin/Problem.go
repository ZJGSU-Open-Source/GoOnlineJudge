package admin

import (
	"GoOnlineJudge/classes"
	"html/template"
	"log"
	"net/http"
)

type ProblemController struct {
	classes.Controller
}

func (this *ProblemController) Add(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Problem")
	this.Init(w, r)

	t, _ := template.ParseFiles("views/admin/menu.tpl", "views/admin/problemadd.tpl")

	this.Data["Title"] = "Add Problem"
	t.Execute(w, this.Data)
}

func (this *ProblemController) Edit(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Problem")
	this.Init(w, r)

	t, _ := template.ParseFiles("views/admin/menu.tpl", "views/admin/problemedit.tpl")

	this.Data["Title"] = "Edit Problem"
	t.Execute(w, this.Data)
}
