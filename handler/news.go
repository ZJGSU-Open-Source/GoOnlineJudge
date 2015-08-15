package handler

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"github.com/zenazn/goji/web"

	"encoding/json"
	"net/http"
	"strconv"
)

//列出所有新闻
//@URL: /api/news @method: GET
func ListNews(c web.C, w http.ResponseWriter, r *http.Request) {

	newsModel := model.NewsModel{}
	newsList, err := newsModel.List(-1, -1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(newsList)
}

//@URL: /api/news/:nid @method: GET
func GetNews(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		Nid = c.URLParams["nid"]
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

	if one.Status == config.StatusReverse {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(one)
}
