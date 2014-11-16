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

	fc.Data["Title"] = "FAQ"
	fc.Data["IsFAQ"] = true
	fc.RenderTemplate("view/layout.tpl", "view/faq.tpl")
}
