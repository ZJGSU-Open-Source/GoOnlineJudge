package controller

import (
	"GoOnlineJudge/class"
	"net/http"
)

type FAQController struct {
	class.Controller
}

//faq 页面
func (this FAQController) Route(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("FAQ Page")
	this.Init(w, r)

	this.Data["Title"] = "FAQ"
	this.Data["IsFAQ"] = true
	err := this.Execute(w, "view/layout.tpl", "view/faq.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
	}
}
