package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type ProblemListController struct {
}

func (this *ProblemListController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem List")

	t, err := template.ParseFiles("views/problemlist.tpl", "views/head.tpl", "views/foot.tpl")
	if err != nil {
		log.Println(err)
	}

	data := &Data{}
	data.Title = "Problem List"
	t.Execute(w, data)
}
