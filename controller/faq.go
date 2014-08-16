package controller

import (
	"GoOnlineJudge/class"
	"html/template"
	"net/http"
)

type FAQController struct {
	class.Controller
}

func (this *FAQController) FAQ(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("FAQ Page")
	this.Init(w, r)

	t, err := template.ParseFiles("view/layout.tpl", "view/faq.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "FAQ"
	this.Data["IsFAQ"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
