package handler

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/middleware"
	"GoOnlineJudge/model"

	"github.com/zenazn/goji/web"

	"encoding/json"
	"net/http"
	"strconv"
)

//列出所有新闻
//@URL: /api/news @method: GET
func ListNews(c web.C, w http.ResponseWriter, r *http.Request) {
	var user = middleware.ToUser(c)

	newsModel := model.NewsModel{}
	newsList, err := newsModel.List(-1, -1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	ok := false
	if user != nil && user.Privilege == config.PrivilegeAD {
		ok = true
	}

	var _newsList []*model.News
	for _, n := range newsList {
		if ok || n.Status == config.StatusAvailable {
			_newsList = append(_newsList, n)
		}
	}

	json.NewEncoder(w).Encode(_newsList)
}

//@URL: /api/news/:nid @method: GET
func GetNews(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		Nid  = c.URLParams["nid"]
		user = middleware.ToUser(c)
	)

	nid, err := strconv.Atoi(Nid) //获取nid
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	newsModel := model.NewsModel{}
	one, err := newsModel.Detail(nid)
	if err != nil {
		http.Error(w, err.Error(), 404)
	}

	ok := false
	if user != nil && user.Privilege == config.PrivilegeAD {
		ok = true
	}

	if one.Status == config.StatusReverse && !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(one)
}
