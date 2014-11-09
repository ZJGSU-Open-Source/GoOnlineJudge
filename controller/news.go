package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"net/http"
	"strconv"
	// "strings"
)

//新闻控件
type NewsController struct {
	class.Controller
}

//列出所有新闻
func (nc NewsController) Get(w http.ResponseWriter, r *http.Request) {

	restweb.Logger.Debug("News List")
	nc.Init(w, r)

	action := nc.GetAction(r.URL.Path, 1)
	if action == "" {

		newsModel := model.NewsModel{}
		newsList, err := newsModel.List(-1, -1)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		nc.Data["News"] = newsList

		nc.Data["Title"] = "Welcome to ZJGSU Online Judge"
		nc.Data["IsNews"] = true
		nc.Execute(w, "view/layout.tpl", "view/news_list.tpl")
	} else {
		nid, err := strconv.Atoi(action) //获取nid
		if err != nil {
			http.Error(w, "args error", 400)
			return
		}

		newsModel := model.NewsModel{}
		one, err := newsModel.Detail(nid)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		nc.Data["Detail"] = one

		if one.Status == config.StatusReverse && nc.Privilege != config.PrivilegeAD { //如果news的状态为普通用户不可见
			nc.Err400(w, r, "No such news", "No such news")
			return
		}

		nc.Data["Title"] = "News Detail"
		nc.Execute(w, "view/layout.tpl", "view/news_detail.tpl")
	}
}
