package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"html/template"
	"net/http"
	"strconv"
)

type news struct {
	Nid int `json:"nid"bson:"nid"`

	Title   string        `json:"title"bson:"title"`
	Content template.HTML `json:"content"bson:"content"`

	Status int    `json:"status"bson:"status"`
	Create string `json:"create"bson:'create'`
}

type NewsController struct {
	class.Controller
}

func (this *NewsController) Detail(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Detail")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	nid, err := strconv.Atoi(args["nid"])
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
	this.Init(w, r)

	one := model.News{}
	one.Title = r.FormValue("title")
	one.Content = r.FormValue("content")

	newsModel := model.NewsModel{}
	err := newsModel.Insert(one)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/admin/news?list", http.StatusFound)
}

func (this *NewsController) Status(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Status")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	nid, err := strconv.Atoi(args["nid"])
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
	http.Redirect(w, r, "/admin/news?list", http.StatusFound)
}

func (this *NewsController) Delete(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin News Delete")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	nid, err := strconv.Atoi(args["nid"])
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

	args := this.ParseURL(r.URL.String())
	nid, err := strconv.Atoi(args["nid"])
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
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	nid, err := strconv.Atoi(args["nid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	one := model.News{}
	newsModel := model.NewsModel{}
	one.Title = r.FormValue("title")
	one.Content = r.FormValue("content")

	err = newsModel.Update(nid, one)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	} else {
		http.Redirect(w, r, "/admin/news?detail/nid?"+strconv.Itoa(nid), http.StatusFound)
	}
}
