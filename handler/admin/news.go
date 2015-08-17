package admin

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/middleware"
	"GoOnlineJudge/model"
	"github.com/zenazn/goji/web"

	"html/template"
	"net/http"
	"strconv"
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

	one := model.News{}
	one.Title = r.FormValue("title")
	one.Content = template.HTML(r.FormValue("content"))

	newsModel := model.NewsModel{}
	err := newsModel.Insert(one)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
}

//@URL: /news/:nid/status @method: PUT
func NewsStatus(c web.C, w http.ResponseWriter, r *http.Request) {

	var (
		Nid  = c.URLParams["nid"]
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	nid, err := strconv.Atoi(Nid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newsModle := model.NewsModel{}
	one, err := newsModle.Detail(nid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var status int
	switch one.Status {
	case config.StatusAvailable:
		status = config.StatusReverse
	default:
		status = config.StatusAvailable
	}

	err = newsModle.Status(nid, status)
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
		Nid  = c.URLParams["nid"]
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	nid, err := strconv.Atoi(Nid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newsModel := model.NewsModel{}
	err = newsModel.Delete(nid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

//@URL: /admin/news/(\d+)/ @method: PUT
func PutNews(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		Nid  = c.URLParams["nid"]
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	nid, err := strconv.Atoi(Nid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	one := model.News{}
	newsModel := model.NewsModel{}
	one.Title = r.FormValue("title")
	one.Content = template.HTML(r.FormValue("content"))

	err = newsModel.Update(nid, one)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
