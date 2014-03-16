package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"log"
	"net/http"
)

type RanklistController struct {
	class.Controller
}

func (this *RanklistController) Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Ranklist")
	this.Init(w, r)

	response, err := http.Post(config.PostHost+"/user/list", "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	one := make(map[string][]*user)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["User"] = one["list"]
	}

	t := template.New("layout.tpl").Funcs(template.FuncMap{"ShowRatio": class.ShowRatio, "ShowStatus": class.ShowStatus})
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
