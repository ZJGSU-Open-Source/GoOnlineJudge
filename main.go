package main

import (
	"log"
	"restweb"

	_ "GoOnlineJudge/schedule"

	"GoOnlineJudge/controller"
	"GoOnlineJudge/controller/admin"
	"GoOnlineJudge/controller/contest"
	"GoOnlineJudge/controller/linkshare"
)

func main() {

	restweb.RegisterController(&controller.FAQController{})
	restweb.RegisterController(&controller.HomeController{})
	restweb.RegisterController(&controller.OSCController{})
	restweb.RegisterController(&controller.ContestController{})
	restweb.RegisterController(&controller.NewsController{})
	restweb.RegisterController(&controller.ProblemController{})
	restweb.RegisterController(&controller.RanklistController{})
	restweb.RegisterController(&controller.SessController{})
	restweb.RegisterController(&controller.StatusController{})
	restweb.RegisterController(&controller.UserController{})
	restweb.RegisterController(&admin.AdminRejudge{})
	restweb.RegisterController(&admin.AdminContest{})
	restweb.RegisterController(&admin.AdminImage{})
	restweb.RegisterController(&admin.AdminProblem{})
	restweb.RegisterController(&admin.AdminTestdata{})
	restweb.RegisterController(&admin.AdminUser{})
	restweb.RegisterController(&admin.AdminHome{})
	restweb.RegisterController(&admin.AdminNews{})
	restweb.RegisterController(&admin.AdminNotice{})
	restweb.RegisterController(&contest.ContestRanklist{})
	restweb.RegisterController(&contest.ContestStatus{})
	restweb.RegisterController(&contest.Contest{})
	restweb.RegisterController(&contest.ContestProblem{})
	restweb.RegisterController(&linkshare.LinkShareController{})

	restweb.AddFile("/static/", ".")
	log.Fatal(restweb.Run())
}
