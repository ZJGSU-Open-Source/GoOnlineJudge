package controllers

import (
	"log"
	"net/http"
)

type NotFoundController struct {
}

func (this *NotFoundController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Not Found")
}
