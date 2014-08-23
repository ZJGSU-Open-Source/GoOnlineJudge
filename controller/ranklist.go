package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"html/template"
	"net/http"
	"strconv"
)

type rank struct {
	user
	Index int `json:"index"bson:"index"`
}

type RanklistController struct {
	class.Controller
}

func (this *RanklistController) Index(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Ranklist")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	url := ""
	this.Data["URL"] = "/ranklist"

	// Page
	if _, ok := args["page"]; !ok {
		args["page"] = "1"
	}

	userModel := model.UserModel{}
	userList, err := userModel.List(nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var count int
	ret := make(map[string][]rank)
	count = 1
	for _, one := range userList {
		if one.Status == config.StatusAvailable {
			count += 1
		}
	}

	var pageCount = (count-1)/config.UserPerPage + 1
	page, err := strconv.Atoi(args["page"])
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
	var len = len(one["list"])
	for _, one := range userList {
		if one.Status == config.StatusAvailable {
			list[count-1].user = *one
			list[count-1].Index = count
			count += 1
		}
	}
	this.Data["User"] = list

	funcMap := map[string]interface{}{
		"ShowRatio":  class.ShowRatio,
		"ShowStatus": class.ShowStatus,
		"NumEqual":   class.NumEqual,
		"ShowTime":   class.ShowTime,
		"NumAdd":     class.NumAdd,
		"NumSub":     class.NumSub,
	}
	t := template.New("layout.tpl").Funcs(funcMap)
	t, err = t.ParseFiles()
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Ranklist"
	this.Data["IsRanklist"] = true
	err = t.Execute(w, this.Data)
	err = this.Excute("view/layout.tpl", "view/ranklist.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
