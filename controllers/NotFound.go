package controllers

import (
	"GoOnlineJudge/classes"
	"log"
	"net/http"
)

type NotFoundController struct {
	classes.Controller
}

func (this *NotFoundController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Not Found")
	this.Init(w, r)
}
