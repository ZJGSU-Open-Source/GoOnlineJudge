package controller

import (
	"GoOnlineJudge/class"
	"html/template"
	"log"
	"net/http"
)

type HomeController struct {
	class.Controller
}

func (this *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")
	this.Init(w, r)

	t, err := template.ParseFiles("view/layout.tpl", "view/home.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Home"
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
