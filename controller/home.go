package controller

import (
	"GoOnlineJudge/class"

	"net/http"
)

type HomeController struct {
	class.Controller
}

func (hc HomeController) Route(w http.ResponseWriter, r *http.Request) {
	hc.Init(w, r)
	newsController := NewsController{}
	newsController.Data = hc.Data
	newsController.List(w, r)
}
