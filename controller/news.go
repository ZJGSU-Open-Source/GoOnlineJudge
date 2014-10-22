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

func (nc NewsController) Route(w http.ResponseWriter, r *http.Request) {
	nc.Init(w, r)
	action := nc.GetAction(r.URL.Path, 1)
	class.Logger.Debug(action)
	defer func() {
		if e := recover(); e != nil {
			http.Error(w, "no such page", 404)
		}
	}()
	rv := class.GetReflectValue(w, r)
	class.CallMethod(&nc, strings.Title(action), rv)
}

//列出所有新闻
func (nc *NewsController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("News List")

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
}

//列出指定新闻的详细信息
func (nc *NewsController) Detail(w http.ResponseWriter, r *http.Request) {
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
	nc.Data["Detail"] = one

	if one.Status == config.StatusReverse && nc.Privilege != config.PrivilegeAD { //如果news的状态为普通用户不可见
		nc.Err400(w, r, "No such news", "No such news")
		return
	}

	nc.Data["Title"] = "News Detail"
	nc.Execute(w, "view/layout.tpl", "view/news_detail.tpl")
}
