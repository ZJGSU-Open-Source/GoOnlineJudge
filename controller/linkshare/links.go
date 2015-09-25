package linkshare

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/model"

	"restweb"

	"net/http"
	"strconv"
)

//新闻控件

type LinkShareController struct {
	class.Controller
} //@Controller

//列出所有新闻
//@URL: /links @method: GET
func (nc *LinkShareController) List() {
	restweb.Logger.Debug("Links List")

	qry := make(map[string]string)

	linkModel := model.LinkModel{}
	linkList, err := linkModel.List(qry)
	if err != nil {
		// http.Error(w, err.Error(), 500)
		return
	}
	nc.Output["Links"] = linkList

	nc.Output["Title"] = "Links Share"
	nc.Output["IsLink"] = true
	nc.RenderTemplate("view/layout.tpl", "view/link_list.tpl")
}

//@URL: /links/(\d+) @method: GET
func (nc *LinkShareController) Detail(Lid string) {
	restweb.Logger.Debug("Links Detail")

	lid, err := strconv.Atoi(Lid) //获取 lid
	if err != nil {
		// http.Error(w, "args error", 400)
		return
	}

	linkModel := model.LinkModel{}
	one, err := linkModel.Detail(lid)
	if err != nil {
		http.Error(nc.W, err.Error(), 500)
	}

	nc.Output["Detail"] = one
	nc.Output["Title"] = "News Detail"
	nc.RenderTemplate("view/layout.tpl", "view/links_detail.tpl")
}

//@URL: /links @method: POST
func (l *LinkShareController) Add() {
	restweb.Logger.Debug("Links Add")

	one := model.Link{}

	one.Link = l.Input.Get("link")
	one.Title = l.Input.Get("title")
	one.Uid = l.Uid

	linkModel := &model.LinkModel{}
	_, err := linkModel.Insert(one)
	if err != nil {
		http.Error(l.W, err.Error(), 500)
	}

	// l.Output["Title"] = "News Detail"
	// l.Output["Detail"] = one
	// l.RenderTemplate("view/layout.tpl", "view/links_detail.tpl")
	l.Redirect("/links", http.StatusFound)
}
