package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"log"
	"net/http"
	// "strconv"
)

type contest struct {
	Cid      int         `json:"cid"bson:"cid"`
	Title    string      `json:"title"bson:"title"`
	Encrypt  int         `json:"encrypt"bson:"encrypt"`
	Argument interface{} `json:"argument"bson:"argument"`

	Start string `json:"start"bson:"start"`
	End   string `json:"end"bson:"end"`

	Status int    `json:"status"bson:"status"`
	Create string `'json:"create"bson:"create"`

	List []int `json:"list"bson:"list"`
}

type ContestController struct {
	class.Controller
}

func (this *ContestController) List(w http.ResponseWriter, r *http.Request) {
	log.Println("Contest List")
	this.Init(w, r)

	response, err := http.Post(config.PostHost+"/contest/list", "application", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	one := make(map[string][]*contest)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["Contest"] = one["list"]
	}
	t := template.New("layout.tpl").Funcs(template.FuncMap{
		"ShowStatus":  class.ShowStatus,
		"ShowExpire":  class.ShowExpire,
		"ShowEncrypt": class.ShowEncrypt,
		"LargePU":     class.LargePU})
	t, err = t.ParseFiles("view/layout.tpl", "view/contest_list.tpl")
	if err != nil {
		log.Println(err)
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Time"] = this.GetTime()
	this.Data["Title"] = "Contest List"
	this.Data["IsContest"] = true
	this.Data["Privilege"] = this.Privilege
	err = t.Execute(w, this.Data)
	if err != nil {
		log.Println(err)
		http.Error(w, "tpl error", 500)
		return
	}

}
