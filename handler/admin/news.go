package admin

import (
	"github.com/zenazn/goji/web"
	"ojapi/config"
	"ojapi/middleware"
	"ojapi/model"

	"encoding/json"
	"html/template"
	"net/http"
)

//@URL: /news @method:POST
func PostNews(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	in := struct {
		Title   string
		Content string
	}{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	one := model.News{}
	one.Title = in.Title
	one.Content = template.HTML(in.Content)

	newsModel := model.NewsModel{}
	err := newsModel.Insert(one)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(one)
}

//@URL: /news/:nid/status @method: PUT
func NewsStatus(c web.C, w http.ResponseWriter, r *http.Request) {

	var (
		news = middleware.ToNews(c)
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var status int
	switch news.Status {
	case config.StatusAvailable:
		status = config.StatusReverse
	default:
		status = config.StatusAvailable
	}

	newsModle := model.NewsModel{}
	err := newsModle.Status(news.Nid, status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

// 删除指定新闻
//@URL: /news/:nid @method: DELETE
func DeleteNews(c web.C, w http.ResponseWriter, r *http.Request) {

	var (
		news = middleware.ToNews(c)
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	newsModel := model.NewsModel{}
	err := newsModel.Delete(news.Nid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

//@URL: /admin/news/(\d+)/ @method: PUT
func PutNews(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		news = middleware.ToNews(c)
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	in := struct {
		Title   string
		Content string
	}{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	one := model.News{}
	newsModel := model.NewsModel{}
	one.Title = in.Title
	one.Content = template.HTML(in.Content)

	err := newsModel.Update(news.Nid, one)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
