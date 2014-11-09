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
	//"GoOnlineJudge/class"
	"GoOnlineJudge/controller"
	// "GoOnlineJudge/controller/admin"
	// "GoOnlineJudge/controller/contest"

	"restweb"

	"net/http"
)

func main() {
	restweb.AddRouter("/", controller.HomeController{})
	restweb.AddRouter("/news/", controller.NewsController{})
	restweb.AddRouter("/problem/", controller.ProblemController{})
	restweb.AddRouter("/status/", controller.StatusController{})
	restweb.AddRouter("/ranklist/", controller.RanklistController{})
	restweb.AddRouter("/contestlist", controller.ContestController{})
	restweb.AddRouter("/user/", controller.UserController{})
	// class.AddRouter("/contest/", contest.ContestUserContorller{})
	// class.AddRouter("/admin/", admin.AdminUserController{})
	restweb.AddRouter("/faq/", controller.FAQController{})
	restweb.AddRouter("/osc/", controller.OSCController{})

	restweb.AddFile("/static/", http.FileServer(http.Dir(".")))
	restweb.Run()
}
