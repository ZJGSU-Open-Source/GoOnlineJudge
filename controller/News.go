package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"log"
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

func (this *NewsController) List(w http.ResponseWriter, r *http.Request) {
	log.Println("News List")
	this.Init(w, r)

	response, err := http.Post(config.PostHost+"/news/list", "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	one := make(map[string][]news)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["News"] = one["list"]
	}

	t := template.New("layout.tpl").Funcs(template.FuncMap{"ShowStatus": class.ShowStatus})
	t, err = t.ParseFiles("view/layout.tpl", "view/news_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "News List"
	this.Data["IsNews"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *NewsController) Detail(w http.ResponseWriter, r *http.Request) {
	log.Println("News Detail")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path)
	nid, err := strconv.Atoi(args["nid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/news/detail/nid/"+strconv.Itoa(nid), "application/json", nil)
	defer response.Body.Close()

	var one news
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["Detail"] = one
	}

	if one.Status == config.StatusReverse && this.Privilege != config.PrivilegeAD {
		t := template.New("layout.tpl")
		t, err = t.ParseFiles("view/layout.tpl", "view/400.tpl")
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}

		this.Data["Title"] = "No such news"
		this.Data["Info"] = "No such news"
		err = t.Execute(w, this.Data)
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
		return
	}

	t := template.New("layout.tpl")
	t, err = t.ParseFiles("view/layout.tpl", "view/news_detail.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "News Detail"
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
