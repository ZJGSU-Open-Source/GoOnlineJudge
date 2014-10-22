// GoOnlineJudge is an online judge system for zjgsu.
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
	class.AddRouter("/faq/", controller.FAQController{})
	class.AddRouter("/osc/", controller.OSCController{})

	class.AddFile("/static/", http.FileServer(http.Dir(".")))
	class.Run()
}
