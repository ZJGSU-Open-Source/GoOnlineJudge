package controllers

import (
	"GoOnlineJudge/classes"
	"log"
	"net/http"
)

type ProblemDetailController struct {
	classes.Controller
}

func (this *ProblemDetailController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem Detail")
	this.Init()
}
