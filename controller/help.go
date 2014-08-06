package controller

import (
	"GoOnlineJudge/class"
	"html/template"
	"log"
	"net/http"
)

type HelpController struct {
	class.Controller
}

func (this *HelpController) Help(w http.ResponseWriter, r *http.Request) {
	log.Println("Help Page")
	this.Init(w, r)

	t, err := template.ParseFiles("view/layout.tpl", "view/help.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Help"
	this.Data["IsHelp"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
