package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
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

	response, err := http.Post(config.PostHost+"/help", "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	//t := template.New("layout.tpl")
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
