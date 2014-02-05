package controllers

import (
	"GoOnlineJudge/classes"
	"GoOnlineJudge/config"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type ProblemController struct {
	classes.Controller
}

func (this *ProblemController) List(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem List")
	this.Init(w, r)

	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		page = 1
	}
	offset := (page - 1) * config.ProblemPerPage

	response, _ := http.Get(config.Host + "/problem/list/offset/" + strconv.Itoa(offset) + "/limit/" + strconv.Itoa(config.ProblemPerPage))
	defer response.Body.Close()

	type problem struct {
		Pid    int
		Title  string
		Source string
		Solve  int
		Submit int
		Status int
	}
	type problemList struct {
		List []*problem
	}

	list := &problemList{}
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(body, list)
		this.Data["Problem"] = list.List
	}

	t, _ := template.ParseFiles("views/layout.tpl", "views/problemlist.tpl")

	this.Data["Title"] = "Problem List"
	t.Execute(w, this.Data)
}

func (this *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem Detail")
	this.Init(w, r)
}
