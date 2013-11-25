package controllers

import (
	"log"
	"net/http"
)

type ProblemListController struct {
}

func (this *ProblemListController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem List")
}
