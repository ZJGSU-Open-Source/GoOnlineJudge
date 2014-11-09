package controller

import (
	"GoOnlineJudge/class"

	"net/http"
)

type HomeController struct {
	class.Controller
}

func (hc HomeController) Get(w http.ResponseWriter, r *http.Request) {
	newsController := NewsController{}
	newsController.Get(w, r)
}
