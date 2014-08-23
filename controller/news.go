package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"net/http"
	"strconv"
)

type NewsController struct {
	class.Controller
}

func (this *NewsController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("News List")
	this.Init(w, r)

	newsModel := model.NewsModel{}
	newsList, err := newsModel.List(-1, -1)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	this.Data["News"] = newsList

	this.Data["Title"] = "Welcome to ZJGSU Online Judge"
	this.Data["IsNews"] = true
	err = this.Execute(w, "view/layout.tpl", "view/news_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *NewsController) Detail(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("News Detail")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	nid, err := strconv.Atoi(args["nid"])
	if err != nil {
		class.Logger.Debug(args["nid"])
		http.Error(w, "args error", 400)
		return
	}

	newsModel := model.NewsModel{}
	one, err := newsModel.Detail(nid)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	this.Data["Detail"] = one

	if one.Status == config.StatusReverse && this.Privilege != config.PrivilegeAD {
		this.Data["Title"] = "No such news"
		this.Data["Info"] = "No such news"
		err = this.Execute(w, "view/layout.tpl", "view/400.tpl")
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
		return
	}

	this.Data["Title"] = "News Detail"
	err = this.Execute(w, "view/layout.tpl", "view/news_detail.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
