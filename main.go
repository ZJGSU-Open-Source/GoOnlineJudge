package main

import (
	"log"
	"restweb"

	"../controller"
	"../controller/admin"
	"../controller/contest"
)

func main() {

	restweb.RegisterController(&controller.ContestController{})
	restweb.RegisterController(&controller.OSCController{})
	restweb.RegisterController(&controller.SessController{})
	restweb.RegisterController(&controller.UserController{})
	restweb.RegisterController(&controller.HomeController{})
	restweb.RegisterController(&controller.NewsController{})
	restweb.RegisterController(&controller.ProblemController{})
	restweb.RegisterController(&controller.RanklistController{})
	restweb.RegisterController(&controller.StatusController{})
	restweb.RegisterController(&controller.FAQController{})
	restweb.RegisterController(&admin.AdminImage{})
	restweb.RegisterController(&admin.AdminRejudge{})
	restweb.RegisterController(&admin.AdminTestdata{})
	restweb.RegisterController(&admin.AdminContest{})
	restweb.RegisterController(&admin.AdminNews{})
	restweb.RegisterController(&admin.AdminNotice{})
	restweb.RegisterController(&admin.AdminProblem{})
	restweb.RegisterController(&admin.AdminUser{})
	restweb.RegisterController(&admin.AdminHome{})
	restweb.RegisterController(&contest.Contest{})
	restweb.RegisterController(&contest.ContestProblem{})
	restweb.RegisterController(&contest.ContestRanklist{})
	restweb.RegisterController(&contest.ContestStatus{})

	restweb.AddFile("/static/", ".")
	log.Fatal(restweb.Run())
}
