package controllers

import (
	"log"
	"net/http"
)

type ProblemDetailController struct {
}

func (this *ProblemDetailController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem Detail")
}
