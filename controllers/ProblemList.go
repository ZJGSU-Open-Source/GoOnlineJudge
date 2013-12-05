package controllers

import (
	"GoOnlineJudge/classes"
	"GoOnlineJudge/config"
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

	t, _ := template.ParseFiles("views/problemlist.tpl", "views/head.tpl", "views/foot.tpl")

	this.Data["Title"] = "Problem List"
	t.Execute(w, this.Data)
}
