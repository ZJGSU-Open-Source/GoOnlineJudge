// GoOnlineJudge Version 14.10.12 is an online judge system for zjgsu.
// Copyright (C) 2013-2014 -  ZJGSU OSC[https://github.com/ZJGSU-Open-Source/]

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 2 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License along
// with this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

package main

import (
	"GoOnlineJudge/controller"
	"GoOnlineJudge/controller/admin"
	"GoOnlineJudge/controller/contest"

	"restweb"

	"log"
	"net/http"
)

func main() {
	restweb.RegisterController(controller.HomeController{})
	restweb.RegisterController(controller.NewsController{})
	restweb.RegisterController(controller.ProblemController{})
	restweb.RegisterController(controller.StatusController{})
	restweb.RegisterController(controller.RanklistController{})
	restweb.RegisterController(controller.ContestController{})
	restweb.RegisterController(controller.UserController{})
	restweb.RegisterController(controller.FAQController{})
	restweb.RegisterController(controller.OSCController{})
	restweb.RegisterController(controller.SessController{})
	restweb.RegisterController(contest.Contest{})
	restweb.RegisterController(contest.ContestRanklist{})
	restweb.RegisterController(contest.ContestProblem{})
	restweb.RegisterController(contest.ContestStatus{})
	restweb.RegisterController(admin.AdminHome{})
	restweb.RegisterController(admin.AdminNews{})
	restweb.RegisterController(admin.AdminProblem{})
	restweb.RegisterController(admin.AdminContest{})
	restweb.RegisterController(admin.AdminRejudge{})
	restweb.RegisterController(admin.AdminTestdata{})

	restweb.RegisterFilters(restweb.ANY, `^/admin`, restweb.Before, requireAdmin)
	restweb.RegisterFilters(restweb.POST, `^/problems/\d+`, restweb.Before, requireLogin)
	restweb.RegisterFilters(restweb.ANY, `^/account`, restweb.Before, requireLogin)
	restweb.RegisterFilters(restweb.GET, `^/user/(settings|profile)`, restweb.Before, requireLogin)
	restweb.RegisterFilters(restweb.POST, `^/user/\w+`, restweb.Before, requireLogin)
	restweb.RegisterFilters(restweb.ANY, `^/contests/\d+`, restweb.Before, requireContest)

	restweb.AddFile("/static/", http.FileServer(http.Dir(".")))
	log.Fatal(restweb.Run())
}
