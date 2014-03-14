package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"encoding/json"
	"html/template"
	"io/ioutil"
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
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, "read error", 500)
			return
		}

		err = json.Unmarshal(body, &one)
		if err != nil {
			http.Error(w, "json error", 500)
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
