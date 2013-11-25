package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Title string
	User  string
}

type HomeController struct {
}

func (this *HomeController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")

	c, _ := r.Cookie("uid")
	uid := "Sign In"
	if c != nil {
		uid = c.Value
	}

	t, err := template.ParseFiles("views/home.tpl", "views/head.tpl", "views/foot.tpl")
	if err != nil {
		log.Println(err)
	}

	data := &Data{}
	data.Title = "Home"
	data.User = "Hi, " + uid
	t.Execute(w, data)
}
