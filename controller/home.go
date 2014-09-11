package controller

import (
	"GoOnlineJudge/class"
	"net/http"
)

type HomeController struct {
	class.Controller
}

func (this *HomeController) Route(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/news/list", http.StatusFound)
}
