package main

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/controller"
	"GoOnlineJudge/controller/admin"
	"GoOnlineJudge/controller/contest"
	"net/http"
)

func main() {
	class.AddRouter("/", controller.HomeController{})
	class.AddRouter("/news/", controller.NewsController{})
	class.AddRouter("/problem/", controller.ProblemController{})
	class.AddRouter("/status/", controller.StatusController{})
	class.AddRouter("/ranklist/", controller.RanklistController{})
	class.AddRouter("/contestlist", controller.ContestController{})
	class.AddRouter("/user/", controller.UserController{})
	class.AddRouter("/contest/", contest.ContestUserContorller{})
	class.AddRouter("/admin/", admin.AdminUserController{})
	class.AddRouter("/FAQ/", controller.FAQController{})
	class.AddRouter("/recruit/", controller.RecruitController{})

	class.AddFile("/static/", http.FileServer(http.Dir(".")))
	class.Run()
}
