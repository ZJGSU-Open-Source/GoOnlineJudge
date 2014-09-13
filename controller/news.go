package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"net/http"
	"strconv"
	"strings"
)

//新闻控件
type NewsController struct {
	class.Controller
}

func (this NewsController) Route(w http.ResponseWriter, r *http.Request) {
	this.Init(w, r)
	action := this.GetAction(r.URL.Path, 1)
	class.Logger.Debug(action)
	defer func() {
		if e := recover(); e != nil {
			http.Error(w, "no such page", 404)
		}
	}()
	rv := class.GetReflectValue(w, r)
	class.CallMethod(&this, strings.Title(action), rv)
}

//列出所有新闻
func (this *NewsController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("News List")

	newsModel := model.NewsModel{}
	newsList, err := newsModel.List(-1, -1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	this.Data["News"] = newsList

	this.Data["Title"] = "Welcome to ZJGSU Online Judge"
	this.Data["IsNews"] = true
	this.Execute(w, "view/layout.tpl", "view/news_list.tpl")
}

//列出指定新闻的详细信息
func (this *NewsController) Detail(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("News Detail")

	args := r.URL.Query()
	nid, err := strconv.Atoi(args.Get("nid")) //获取nid
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	newsModel := model.NewsModel{}
	one, err := newsModel.Detail(nid)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	this.Data["Detail"] = one

	if one.Status == config.StatusReverse && this.Privilege != config.PrivilegeAD { //如果news的状态为普通用户不可见
		this.Err400(w, r, "No such news", "No such news")
		return
	}

	this.Data["Title"] = "News Detail"
	this.Execute(w, "view/layout.tpl", "view/news_detail.tpl")
}
