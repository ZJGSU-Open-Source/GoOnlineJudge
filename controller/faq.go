package controller

import (
	"GoOnlineJudge/class"
	"net/http"
	"restweb"
)

type FAQController struct {
	class.Controller
}

//faq 页面
func (fc FAQController) Get(w http.ResponseWriter, r *http.Request) {
	restweb.Logger.Debug("FAQ Page")
	fc.Init(w, r)

	fc.Data["Title"] = "FAQ"
	fc.Data["IsFAQ"] = true
	fc.Execute(w, "view/layout.tpl", "view/faq.tpl")
}
