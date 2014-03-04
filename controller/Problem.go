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

type problem struct {
	Pid int `json:"pid"bson:"pid"`

	Time   int `json:"time"bson:"time"`
	Memory int `json:"memory"bson:"memory"`

	Title       string `json:"title"bson:"title"`
	Description string `json:"description"bson:"description"`
	Input       string `json:"input"bson:"input"`
	Output      string `json:"output"bson:"output"`
	Source      string `json:"source"bson:"source"`
	Hint        string `json:"hint"bson:"hint"`

	In  string `json:"in"bson:"in"`
	Out string `json:"out"bson:"out"`

	Solve  int `json:"solve"bson:"solve"`
	Submit int `json:"submit"bson:"submit"`

	Status int    `json:"status"bson:"status"`
	Create string `json:"create"bson:"create"`
}

type ProblemController struct {
	class.Controller
}

func (this *ProblemController) List(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem List")
	this.Init(w, r)

	response, err := http.Post(config.PostHost+"/problem/list", "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	one := make(map[string][]*problem)
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
		this.Data["Problem"] = one["list"]
	}

	t := template.New("layout.tpl").Funcs(template.FuncMap{"ShowRatio": class.ShowRatio, "ShowStatus": class.ShowStatus})
	t, err = t.ParseFiles("view/layout.tpl", "view/problem_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Problem List"
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (this *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem Detail")
	this.Init(w, r)
}
