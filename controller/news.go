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
} //@Controller

//列出所有新闻
//@URL: /api/news @method: GET
func (nc *NewsController) List() {
	restweb.Logger.Debug("News List")

	newsModel := model.NewsModel{}
	newsList, err := newsModel.List(-1, -1)
	if err != nil {
		// http.Error(w, err.Error(), 500)
		return
	}
	nc.Output["News"] = newsList
	nc.RenderJson()
}

//@URL: /api/news/(\d+) @method: GET
func (nc *NewsController) Detail(Nid string) {
	nid, err := strconv.Atoi(Nid) //获取nid
	if err != nil {
		// http.Error(w, "args error", 400)
		return
	}

	newsModel := model.NewsModel{}
	one, err := newsModel.Detail(nid)
	if err != nil {
		http.Error(nc.W, err.Error(), 404)
	}
	nc.Output["Detail"] = one

	if one.Status == config.StatusReverse {
		nc.Error("No such resource", 404)
		return
	}
	nc.RenderJson()
}
