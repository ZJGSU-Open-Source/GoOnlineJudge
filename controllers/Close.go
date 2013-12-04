package controllers

import (
	"GoOnlineJudge/classes"
	"html/template"
	"log"
	"net/http"
)

type CloseController struct {
	classes.Controller
}

func (this *CloseController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Close")
	this.Init()

	t, _ := template.ParseFiles("views/close.tpl", "views/head.tpl", "views/foot.tpl")

	this.Data["Title"] = "Feature Closed"
	t.Execute(w, this.Data)
}
