package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"net/http"
	"strconv"
)

//新闻控件
type NewsController struct {
	class.Controller
}

//列出所有新闻
func (nc NewsController) List() {
	restweb.Logger.Debug("News List")

	newsModel := model.NewsModel{}
	newsList, err := newsModel.List(-1, -1)
	if err != nil {
		// http.Error(w, err.Error(), 500)
		return
	}
	nc.Output["News"] = newsList

	nc.Output["Title"] = "Welcome to ZJGSU Online Judge"
	nc.Output["IsNews"] = true
	nc.RenderTemplate("view/layout.tpl", "view/news_list.tpl")
}

func (nc NewsController) Detail(Nid string) {
	nid, err := strconv.Atoi(Nid) //获取nid
	if err != nil {
		// http.Error(w, "args error", 400)
		return
	}

	newsModel := model.NewsModel{}
	one, err := newsModel.Detail(nid)
	if err != nil {
		http.Error(nc.W, err.Error(), 500)
	}
	nc.Output["Detail"] = one

	if one.Status == config.StatusReverse {
		nc.Err400("No such news", "No such news")
		return
	}

	nc.Output["Title"] = "News Detail"
	nc.RenderTemplate("view/layout.tpl", "view/news_detail.tpl")
}
