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
	nc.Data["News"] = newsList

	nc.Data["Title"] = "Welcome to ZJGSU Online Judge"
	nc.Data["IsNews"] = true
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
		http.Error(nc.Response, err.Error(), 500)
	}
	nc.Data["Detail"] = one

	if one.Status == config.StatusReverse {
		nc.Err400("No such news", "No such news")
		return
	}

	nc.Data["Title"] = "News Detail"
	nc.RenderTemplate("view/layout.tpl", "view/news_detail.tpl")
}
