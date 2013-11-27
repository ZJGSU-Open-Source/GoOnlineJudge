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

	cookie, _ := r.Cookie("uid")
	user := "Sign In"
	if cookie != nil {
		user = "Hi, " + cookie.Value
	}

	t, err := template.ParseFiles("views/problemlist.tpl", "views/head.tpl", "views/foot.tpl")
	if err != nil {
		log.Println(err)
	}

	data := &Data{}
	data.Title = "Home"
	data.User = user
	t.Execute(w, data)
}
