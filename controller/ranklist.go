package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"net/http"
	"strconv"
)

// 排名
type rank struct {
	model.User
	Index int `json:"index"bson:"index"`
}

// 排名控件
type RanklistController struct {
	class.Controller
}

func (this *RanklistController) Route(w http.ResponseWriter, r *http.Request) {
	this.Init(w, r)
	this.Index(w, r)
}

func (this *RanklistController) Index(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Ranklist")

	args := r.URL.Query()

	// Page

	if v := args.Get("page"); v == "" {
		args.Set("page", "1")
	}

	userModel := model.UserModel{}
	userList, err := userModel.List(nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var count int
	count = 1
	for _, one := range userList {
		if one.Status == config.StatusAvailable {
			count += 1
		}
	}

	var pageCount = (count-1)/config.UserPerPage + 1
	page, err := strconv.Atoi(args.Get("page"))
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	if page > pageCount {
		http.Error(w, "args error", 400)
		return
	}

	pageData := this.GetPage(page, pageCount)
	for k, v := range pageData {
		this.Data[k] = v
	}

	qry := make(map[string]string)
	qry["offset"] = strconv.Itoa((page - 1) * config.UserPerPage)
	qry["limit"] = strconv.Itoa(config.UserPerPage)
	userList, err = userModel.List(qry)
	if err != nil {

	}

	list := make([]rank, len(userList), len(userList))
	count = 1
	for i, one := range userList {
		list[i].User = *one
		if one.Status == config.StatusAvailable {
			list[count-1].Index = count
			count += 1
		}
	}
	this.Data["URL"] = "/ranklist?"
	this.Data["User"] = list
	this.Data["Title"] = "Ranklist"
	this.Data["IsRanklist"] = true
	err = this.Execute(w, "view/layout.tpl", "view/ranklist.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
	}
}
