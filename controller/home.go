package controller

import (
	"GoOnlineJudge/class"
	"net/http"
)

type HomeController struct {
	class.Controller
}

func (this HomeController) Route(w http.ResponseWriter, r *http.Request) {
	this.Init(w, r)
	newsController := NewsController{}
	newsController.Data = this.Data
	newsController.List(w, r)
}
