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
	this.Init()

	if !config.Problem {
		http.Redirect(w, r, "/close/", http.StatusFound)
	}

	c, err := r.Cookie("uid")
	this.CheckError(err)
	if c.Value == "" {
		log.Println("Logout")
	}

	m := &models.ProblemModel{}
	list := m.List()

	t, err := template.ParseFiles("views/problemlist.tpl", "views/head.tpl", "views/foot.tpl")
	this.CheckError(err)

	this.Data["Title"] = "Problem List"
	this.Data["Problem"] = list
	t.Execute(w, this.Data)
}
