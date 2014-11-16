package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/model"
	"net/http"
	"restweb"
)

type HomeController struct {
	class.Controller
}

func (hc HomeController) Index() {
	// newsController := NewsController{}
	// newsController.Get()

	restweb.Logger.Debug("Home")

	newsModel := model.NewsModel{}
	newsList, err := newsModel.List(-1, -1)
	if err != nil {
		http.Error(hc.Response, err.Error(), 500)
		return
	}
	hc.Data["News"] = newsList
	hc.Data["Title"] = "Welcome to ZJGSU Online Judge"
	hc.Data["IsNews"] = true

	hc.RenderTemplate("view/layout.tpl", "view/news_list.tpl")
}
