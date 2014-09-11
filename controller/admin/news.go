package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"html/template"
	"net/http"
	"strconv"
)

//news新闻控件
type NewsController struct {
	class.Controller
}

func (this *NewsController) Route(w http.ResponseWriter, r *http.Request) {
	this.Init(w, r)
	action := this.GetAction(r.URL.Path, 2)
	switch action {
	case "detail":
		this.Detail(w, r)
	case "list":
		this.List(w, r)
	case "edit":
		this.Edit(w, r)
	case "update":
		this.Update(w, r)
	case "delete":
		this.Delete(w, r)
	case "status":
		this.Status(w, r)
	case "add":
		this.Add(w, r)
	case "insert":
		this.Insert(w, r)
	default:
		http.Error(w, "no such page", 404)
	}

}

//新闻详细信息
func (this *NewsController) Detail(w http.ResponseWriter, r *http.Request) {
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
	this.Data["Detail"] = one

	this.Data["Title"] = "Admin - News Detail"
	this.Data["IsNews"] = true
	this.Data["IsList"] = false

	err = this.Execute(w, "view/admin/layout.tpl", "view/news_detail.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

// 列出所有新闻
func (this *NewsController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News List")
	this.Init(w, r)

	newsModel := model.NewsModel{}
	newlist, err := newsModel.List(-1, -1)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	this.Data["News"] = newlist
	this.Data["Title"] = "Admin - News List"
	this.Data["IsNews"] = true
	this.Data["IsList"] = true
	err = this.Execute(w, "view/admin/layout.tpl", "view/admin/news_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *NewsController) Add(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Add")
	this.Init(w, r)

	this.Data["Title"] = "Admin - News Add"
	this.Data["IsNews"] = true
	this.Data["IsAdd"] = true
	this.Data["IsEdit"] = true

	err := this.Execute(w, "view/admin/layout.tpl", "view/admin/news_add.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *NewsController) Insert(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Insert")

	if r.Method != "POST" {
		this.Err400(w, r, "Error", "Error Method to Insert news")
		return
	}

	this.Init(w, r)

	if this.Privilege != config.PrivilegeAD {
		this.Err400(w, r, "Warning", "Error Privilege to Insert news")
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

func (this *NewsController) Status(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Status")
	if r.Method != "POST" {
		this.Err400(w, r, "Error", "Error Method to change news status")
		return
	}

	this.Init(w, r)

	if this.Privilege != config.PrivilegeAD {
		this.Err400(w, r, "Warning", "Error Privilege to change news status")
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
func (this *NewsController) Delete(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Delete")
	if r.Method != "POST" {
		this.Err400(w, r, "Error", "Error Method to Delete news")
		return
	}

	this.Init(w, r)

	if this.Privilege != config.PrivilegeAD {
		this.Err400(w, r, "Warning", "Error Privilege to Delete news")
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

func (this *NewsController) Edit(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Edit")
	this.Init(w, r)

	nid, err := strconv.Atoi(r.URL.Query().Get("nid"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	newsModel := model.NewsModel{}
	one, err := newsModel.Detail(nid)
	this.Data["Detail"] = one
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	this.Data["Title"] = "Admin - News Edit"
	this.Data["IsNews"] = true
	this.Data["IsList"] = false
	this.Data["IsEdit"] = true

	err = this.Execute(w, "view/admin/layout.tpl", "view/admin/news_edit.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *NewsController) Update(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Update")
	if r.Method != "POST" {
		this.Err400(w, r, "Error", "Error Method to Update news")
		return
	}

	this.Init(w, r)

	if this.Privilege != config.PrivilegeAD {
		this.Err400(w, r, "Warning", "Error Privilege to Update news")
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
