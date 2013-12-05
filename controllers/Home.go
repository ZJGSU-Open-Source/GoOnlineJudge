package controllers

import (
	"GoOnlineJudge/classes"
	"html/template"
	"log"
	"net/http"
)

type HomeController struct {
	classes.Controller
}

func (this *HomeController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")
	this.Init(w, r)

	t, _ := template.ParseFiles("views/home.tpl", "views/head.tpl", "views/foot.tpl")

	this.Data["Title"] = "Home"
	t.Execute(w, this.Data)
}
