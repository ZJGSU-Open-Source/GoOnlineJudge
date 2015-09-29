package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"html/template"
	"net/http"
	"strconv"
)

//news新闻控件

type AdminNews struct {
	class.Controller
} //@Controller

// //新闻详细信息
// func (nc *AdminNews) Detail() {
// 	restweb.Logger.Debug("Admin News Detail")

// 	nid, err := strconv.Atoi(r.URL.Query().Get("nid"))
// 	if err != nil {
// 		nc.Error("args error", 400)
// 		return
// 	}

// 	newsModel := model.NewsModel{}
// 	one, err := newsModel.Detail(nid)
// 	if err != nil {
// 		nc.Error("load error", 400)
// 		return
// 	}
// 	nc.Output["Detail"] = one

// 	nc.Output["Title"] = "Admin - News Detail"
// 	nc.Output["IsNews"] = true
// 	nc.Output["IsList"] = false

// 	nc.RenderTemplate("view/admin/layout.tpl", "view/news_detail.tpl")
// }

// 列出所有新闻
//@URL: /admin/news/ @method: GET
func (nc *AdminNews) List() {
	restweb.Logger.Debug("Admin News List")

	newsModel := model.NewsModel{}
	newlist, err := newsModel.List(nil, -1, -1)
	if err != nil {
		nc.Error(err.Error(), 500)
		return
	}
	nc.Output["News"] = newlist
	nc.Output["Title"] = "Admin - News List"
	nc.Output["IsNews"] = true
	nc.Output["IsList"] = true
	nc.RenderTemplate("view/admin/layout.tpl", "view/admin/news_list.tpl")
}

//@URL: /admin/news/new/ @method: GET
func (nc *AdminNews) Add() {
	restweb.Logger.Debug("Admin News Add")

	nc.Output["Title"] = "Admin - News Add"
	nc.Output["IsNews"] = true
	nc.Output["IsAdd"] = true
	nc.Output["IsEdit"] = true

	nc.RenderTemplate("view/admin/layout.tpl", "view/admin/news_add.tpl")

}

//@URL: /admin/news/ @method:POST
func (nc *AdminNews) Insert() {
	restweb.Logger.Debug("Admin News Insert")

	if nc.Privilege != config.PrivilegeAD {
		nc.Err400("Warning", "Error Privilege to Insert news")
		return
	}

	one := model.News{}
	one.Title = nc.R.FormValue("title")
	one.Content = template.HTML(nc.Input.Get("content"))

	newsModel := model.NewsModel{}
	err := newsModel.Insert(one)
	if err != nil {
		nc.Error(err.Error(), 500)
		return
	}

	nc.Redirect("/admin/news", http.StatusFound)
}

//@URL: /admin/news/(\d+)/status/ @method: POST
func (nc *AdminNews) Status(Nid string) {
	restweb.Logger.Debug("Admin News Status")

	if nc.Privilege != config.PrivilegeAD {
		nc.Err400("Warning", "Error Privilege to change news status")
		return
	}

	nid, err := strconv.Atoi(Nid)
	if err != nil {
		nc.Error("args error", 400)
		return
	}

	newsModle := model.NewsModel{}
	one, err := newsModle.Detail(nid)
	if err != nil {
		nc.Error(err.Error(), 400)
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
		nc.Error(err.Error(), 400)
		return
	}
	nc.Redirect("/admin/news", http.StatusFound)
}

// 删除指定新闻
//@URL: /admin/news/(\d+)/ @method: DELETE
func (nc *AdminNews) Delete(Nid string) {
	restweb.Logger.Debug("Admin News Delete")

	if nc.Privilege != config.PrivilegeAD {
		nc.Err400("Warning", "Error Privilege to Delete news")
		return
	}

	nid, err := strconv.Atoi(Nid)
	if err != nil {
		nc.Error("args error", 400)
		return
	}

	newsModel := model.NewsModel{}
	err = newsModel.Delete(nid)
	if err != nil {
		nc.Error(err.Error(), 400)
		return
	}

	nc.W.WriteHeader(200)
}

//@URL: /admin/news/(\d+)/ @method: GET
func (nc *AdminNews) Edit(Nid string) {
	restweb.Logger.Debug("Admin News Edit")

	nid, err := strconv.Atoi(Nid)
	if err != nil {
		nc.Error("args error", 400)
		return
	}

	newsModel := model.NewsModel{}
	one, err := newsModel.Detail(nid)
	nc.Output["Detail"] = one
	if err != nil {
		nc.Error(err.Error(), 400)
		return
	}

	nc.Output["Title"] = "Admin - News Edit"
	nc.Output["IsNews"] = true
	nc.Output["IsList"] = false
	nc.Output["IsEdit"] = true

	nc.RenderTemplate("view/admin/layout.tpl", "view/admin/news_edit.tpl")
}

//@URL: /admin/news/(\d+)/ @method: POST
func (nc *AdminNews) Update(Nid string) {
	restweb.Logger.Debug("Admin News Update")

	if nc.Privilege != config.PrivilegeAD {
		nc.Err400("Warning", "Error Privilege to Update news")
		return
	}

	nid, err := strconv.Atoi(Nid)
	if err != nil {
		nc.Error("args error", 400)
		return
	}

	one := model.News{}
	newsModel := model.NewsModel{}
	one.Title = nc.Input.Get("title")
	one.Content = template.HTML(nc.Input.Get("content"))

	err = newsModel.Update(nid, one)
	if err != nil {
		nc.Error(err.Error(), 500)
		return
	} else {
		nc.Redirect("/admin/news", http.StatusFound)
	}
}
