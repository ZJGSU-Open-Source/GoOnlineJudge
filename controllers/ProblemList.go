package controllers

import (
	"GoOnlineJudge/config"
	"html/template"
	"log"
	"net/http"
)

type ProblemListController struct {
}

func (this *ProblemListController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem List")

	t, err := this.getTemplate()
	if err != nil {
		log.Println(err)
	}

	data := &Data{}
	data.Title = "Problem List"
	t.Execute(w, data)
}

func (this *ProblemListController) getTemplate() (*template.Template, error) {
	if config.Problem {
		return template.ParseFiles("views/problemlist.tpl", "views/head.tpl", "views/foot.tpl")
	} else {
		return template.ParseFiles("views/close.tpl", "views/head.tpl", "views/foot.tpl")
	}
}
