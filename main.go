package main

import (
	"log"
	"restweb"

	"GoOnlineJudge/controller"
	"GoOnlineJudge/controller/admin"
	"GoOnlineJudge/controller/contest"
	_ "GoOnlineJudge/schedule"
)

func main() {

	restweb.RegisterController(&controller.StatusController{})
	restweb.RegisterController(&controller.UserController{})
	restweb.RegisterController(&controller.ContestController{})
	restweb.RegisterController(&controller.HomeController{})
	restweb.RegisterController(&controller.NewsController{})
	restweb.RegisterController(&controller.OSCController{})
	restweb.RegisterController(&controller.ProblemController{})
	restweb.RegisterController(&controller.SessController{})
	restweb.RegisterController(&controller.FAQController{})
	restweb.RegisterController(&controller.RanklistController{})
	restweb.RegisterController(&admin.AdminUser{})
	restweb.RegisterController(&admin.AdminContest{})
	restweb.RegisterController(&admin.AdminHome{})
	restweb.RegisterController(&admin.AdminImage{})
	restweb.RegisterController(&admin.AdminNews{})
	restweb.RegisterController(&admin.AdminNotice{})
	restweb.RegisterController(&admin.AdminProblem{})
	restweb.RegisterController(&admin.AdminRejudge{})
	restweb.RegisterController(&admin.AdminTestdata{})
	restweb.RegisterController(&contest.Contest{})
	restweb.RegisterController(&contest.ContestProblem{})
	restweb.RegisterController(&contest.ContestRanklist{})
	restweb.RegisterController(&contest.ContestStatus{})

	restweb.AddFile("/static/", ".")
	log.Fatal(restweb.Run())
}
