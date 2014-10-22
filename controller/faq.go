package controller

import (
	"GoOnlineJudge/class"

	"net/http"
)

type FAQController struct {
	class.Controller
}

//faq 页面
func (fc FAQController) Route(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("FAQ Page")
	fc.Init(w, r)

	fc.Data["Title"] = "FAQ"
	fc.Data["IsFAQ"] = true
	fc.Execute(w, "view/layout.tpl", "view/faq.tpl")
}
