package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"net/http"
	"restweb"
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
} //@Controller

//@URL: /ranklist @method: GET
func (rc *RanklistController) Index() {
	restweb.Logger.Debug("Ranklist")

	// Page

	if _, ok := rc.Input["page"]; !ok {
		rc.Input.Set("page", "1")
	}

	userModel := model.UserModel{}
	userList, err := userModel.List(nil)
	if err != nil {
		http.Error(rc.W, err.Error(), 400)
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
	page, err := strconv.Atoi(rc.Input.Get("page"))
	if err != nil {
		rc.Error("args error", 400)
		return
	}
	if page > pageCount {
		rc.Error("args error", 400)
		return
	}

	pageData := rc.GetPage(page, page >= pageCount)
	for k, v := range pageData {
		rc.Output[k] = v
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
			list[count-1].Index = count + (page-1)*config.UserPerPage
			count += 1
		}
	}
	rc.Output["URL"] = "/ranklist?"
	rc.Output["User"] = list
	rc.Output["Title"] = "Ranklist"
	rc.Output["IsRanklist"] = true
	rc.RenderTemplate("view/layout.tpl", "view/ranklist.tpl")
}
