package main

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/controller"
	"GoOnlineJudge/controller/admin"
	"GoOnlineJudge/controller/contest"
	"net/http"
	"strings"
)

var RouterMap = map[string]class.Router{
	"/":            controller.HomeController{},
	"/news/":       controller.NewsController{},
	"/problem/":    controller.ProblemController{},
	"/status/":     controller.StatusController{},
	"/ranklist/":   controller.RanklistController{},
	"/contestlist": controller.ContestController{},
	"/user/":       controller.UserController{},
	"/contest/":    contest.ContestUserContorller{},
	"/admin/":      admin.AdminUserController{},
	"/FAQ/":        controller.FAQController{},
}

type Server struct {
}

func (this *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path + "/"

	class.Logger.Debug(path)
	if strings.HasPrefix(path, "/static/") {
		http.FileServer(http.Dir(".")).ServeHTTP(w, r)
		return
	}
	maxlenth := 0
	var realRouter class.Router
	for pattern, router := range RouterMap {
		if len(pattern) > maxlenth && strings.HasPrefix(path, pattern) {
			maxlenth = len(pattern)
			realRouter = router
		}
	}
	if maxlenth > 0 {
		realRouter.Route(w, r)
	} else {
		http.Error(w, "no such page", 404)
	}
}
