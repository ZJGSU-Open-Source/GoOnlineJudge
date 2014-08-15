package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"log"
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
	log.Println("Ranklist")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	url := ""
	this.Data["URL"] = "/ranklist"

	// Page
	if _, ok := args["page"]; !ok {
		args["page"] = "1"
	}

	response, err := http.Post(config.PostHost+"/user/list", "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	var count int
	ret := make(map[string][]rank)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &ret)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		count = 1
		var len = len(ret["list"])
		for i := 0; i < len; i++ {
			if ret["list"][i].Status == config.StatusAvailable {
				ret["list"][i].Index = count
				count += 1
			}
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
	url += "/offset/" + strconv.Itoa((page-1)*config.UserPerPage) + "/limit/" + strconv.Itoa(config.UserPerPage)
	pageData := this.GetPage(page, pageCount)
	for k, v := range pageData {
		this.Data[k] = v
	}

	//
	response, err = http.Post(config.PostHost+"/user/list"+url, "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	one := make(map[string][]rank)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		var count = 1
		var len = len(one["list"])
		for i := 0; i < len; i++ {
			if one["list"][i].Status == config.StatusAvailable {
				one["list"][i].Index = count
				count += 1
			}
		}
		this.Data["User"] = one["list"]
	}

	funcMap := map[string]interface{}{
		"ShowRatio":  class.ShowRatio,
		"ShowStatus": class.ShowStatus,
		"NumEqual":   class.NumEqual,
		"NumAdd":     class.NumAdd,
		"NumSub":     class.NumSub,
	}
	t := template.New("layout.tpl").Funcs(funcMap)
	t, err = t.ParseFiles("view/layout.tpl", "view/ranklist.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Ranklist"
	this.Data["IsRanklist"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
