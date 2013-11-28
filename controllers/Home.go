package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Title string
}

type HomeController struct {
}

func (this *HomeController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")

	t, err := template.ParseFiles("views/home.tpl", "views/head.tpl", "views/foot.tpl")
	if err != nil {
		log.Println(err)
	}

	data := &Data{}
	data.Title = "Home"
	t.Execute(w, data)
}
