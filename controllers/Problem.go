package controllers

import (
	"GoOnlineJudge/classes"
	"html/template"
	"log"
	"net/http"
)

type ProblemController struct {
	classes.Controller
}

func (this *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem Detail")
	this.Init(w, r)
}

func (this *ProblemController) List(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem List")
	this.Init(w, r)

	t, _ := template.ParseFiles("views/layout.tpl", "views/problemlist.tpl")

	this.Data["Title"] = "Problem List"
	t.Execute(w, this.Data)
}
