package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"html/template"
	"net/http"
	"strconv"
	"strings"
)

//news新闻控件
type NewsController struct {
	class.Controller
}

func (nc NewsController) Route(w http.ResponseWriter, r *http.Request) {
	nc.Init(w, r)
	action := nc.GetAction(r.URL.Path, 2)
	defer func() {
		if e := recover(); e != nil {
			http.Error(w, "no such page", 404)
		}
	}()
	rv := class.GetReflectValue(w, r)
	class.CallMethod(&nc, strings.Title(action), rv)
}

//新闻详细信息
func (nc *NewsController) Detail(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Detail")

	nid, err := strconv.Atoi(r.URL.Query().Get("nid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	newsModel := model.NewsModel{}
	one, err := newsModel.Detail(nid)
	if err != nil {
		http.Error(w, "load error", 400)
		return
	}
	nc.Data["Detail"] = one

	nc.Data["Title"] = "Admin - News Detail"
	nc.Data["IsNews"] = true
	nc.Data["IsList"] = false

	nc.Execute(w, "view/admin/layout.tpl", "view/news_detail.tpl")
}

// 列出所有新闻
func (nc *NewsController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News List")
	nc.Init(w, r)

	newsModel := model.NewsModel{}
	newlist, err := newsModel.List(-1, -1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	nc.Data["News"] = newlist
	nc.Data["Title"] = "Admin - News List"
	nc.Data["IsNews"] = true
	nc.Data["IsList"] = true
	nc.Execute(w, "view/admin/layout.tpl", "view/admin/news_list.tpl")
}

func (nc *NewsController) Add(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Add")
	nc.Init(w, r)

	nc.Data["Title"] = "Admin - News Add"
	nc.Data["IsNews"] = true
	nc.Data["IsAdd"] = true
	nc.Data["IsEdit"] = true

	nc.Execute(w, "view/admin/layout.tpl", "view/admin/news_add.tpl")

}

func (nc *NewsController) Insert(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Insert")

	if r.Method != "POST" {
		nc.Err400(w, r, "Error", "Error Method to Insert news")
		return
	}

	nc.Init(w, r)

	if nc.Privilege != config.PrivilegeAD {
		nc.Err400(w, r, "Warning", "Error Privilege to Insert news")
		return
	}

	one := model.News{}
	one.Title = r.FormValue("title")
	one.Content = template.HTML(r.FormValue("content"))

	newsModel := model.NewsModel{}
	err := newsModel.Insert(one)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/admin/news/list", http.StatusFound)
}

func (nc *NewsController) Status(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Status")
	if r.Method != "POST" {
		nc.Err400(w, r, "Error", "Error Method to change news status")
		return
	}

	nc.Init(w, r)

	if nc.Privilege != config.PrivilegeAD {
		nc.Err400(w, r, "Warning", "Error Privilege to change news status")
		return
	}

	nid, err := strconv.Atoi(r.URL.Query().Get("nid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	newsModle := model.NewsModel{}
	one, err := newsModle.Detail(nid)
	if err != nil {
		http.Error(w, err.Error(), 400)
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
		http.Error(w, err.Error(), 400)
		return
	}
	http.Redirect(w, r, "/admin/news/list", http.StatusFound)
}

// 删除指定新闻
func (nc *NewsController) Delete(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Delete")
	if r.Method != "POST" {
		nc.Err400(w, r, "Error", "Error Method to Delete news")
		return
	}

	nc.Init(w, r)

	if nc.Privilege != config.PrivilegeAD {
		nc.Err400(w, r, "Warning", "Error Privilege to Delete news")
		return
	}

	nid, err := strconv.Atoi(r.URL.Query().Get("nid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	newsModel := model.NewsModel{}
	err = newsModel.Delete(nid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func (nc *NewsController) Edit(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Edit")
	nc.Init(w, r)

	nid, err := strconv.Atoi(r.URL.Query().Get("nid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	newsModel := model.NewsModel{}
	one, err := newsModel.Detail(nid)
	nc.Data["Detail"] = one
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	nc.Data["Title"] = "Admin - News Edit"
	nc.Data["IsNews"] = true
	nc.Data["IsList"] = false
	nc.Data["IsEdit"] = true

	nc.Execute(w, "view/admin/layout.tpl", "view/admin/news_edit.tpl")
}

func (nc *NewsController) Update(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Update")
	if r.Method != "POST" {
		nc.Err400(w, r, "Error", "Error Method to Update news")
		return
	}

	nc.Init(w, r)

	if nc.Privilege != config.PrivilegeAD {
		nc.Err400(w, r, "Warning", "Error Privilege to Update news")
		return
	}

	nid, err := strconv.Atoi(r.URL.Query().Get("nid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	one := model.News{}
	newsModel := model.NewsModel{}
	one.Title = r.FormValue("title")
	one.Content = template.HTML(r.FormValue("content"))

	err = newsModel.Update(nid, one)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	} else {
		http.Redirect(w, r, "/admin/news/detail?nid="+strconv.Itoa(nid), http.StatusFound)
	}
}
