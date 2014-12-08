package controller

import (
	"GoOnlineJudge/class"
	"restweb"
)

type FAQController struct {
	class.Controller
}

//faq 页面
func (fc FAQController) Index() {
	restweb.Logger.Debug("FAQ Page")

	fc.Output["Title"] = "FAQ"
	fc.Output["IsFAQ"] = true
	fc.RenderTemplate("view/layout.tpl", "view/faq.tpl")
}
