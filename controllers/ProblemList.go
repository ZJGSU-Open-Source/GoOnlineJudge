package controllers

import (
	"GoOnlineJudge/classes"
	"GoOnlineJudge/config"
	"GoOnlineJudge/models"
	"html/template"
	"log"
	"net/http"
)

type ProblemListController struct {
	classes.Controller
}

func (this *ProblemListController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem List")
	this.Init(w, r)

	if !config.Problem {
		http.Redirect(w, r, "/close", http.StatusFound)
	}

	m := &models.ProblemModel{}
	list := m.List()

	t, _ := template.ParseFiles("views/problemlist.tpl", "views/head.tpl", "views/foot.tpl")

	this.Data["Title"] = "Problem List"
	this.Data["Problem"] = list
	t.Execute(w, this.Data)
}
