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
	"GoOnlineJudge/config"
	"GoOnlineJudge/controller"
	"GoOnlineJudge/controller/admin"
	"GoOnlineJudge/controller/contest"
	"GoOnlineJudge/model"

	"restweb"

	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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
	restweb.RegisterFilters(restweb.ANY, `^/contests/\d+`, restweb.Before, requireContest)
	restweb.RegisterFilters(restweb.POST, `^/user/\w+`, restweb.Before, requireLogin)

	restweb.AddFile("/static/", http.FileServer(http.Dir(".")))
	log.Fatal(restweb.Run())
}

func requireAdmin(ctx *restweb.Context) bool {
	uid := ctx.GetSession("Uid")
	if uid == "" {
		ctx.Redirect("/sess", http.StatusFound)
		return true
	}

	prv := ctx.GetSession("Privilege")
	prt, err := strconv.Atoi(prv)
	if err != nil {
		ctx.Error("Error occured", http.StatusForbidden)
		return true
	}
	if prt < config.PrivilegeTC {
		ctx.Error("No privilege", http.StatusForbidden)
		return true
	}
	return false
}

func requireLogin(ctx *restweb.Context) bool {
	uid := ctx.GetSession("Uid")
	if uid == "" {
		if ctx.Requset.Method == restweb.GET {
			ctx.Redirect("/sess", http.StatusFound)
		} else {
			ctx.Response.WriteHeader(http.StatusUnauthorized)
		}
		return true
	}
	return false
}

func requireContest(ctx *restweb.Context) bool {
	url := ctx.Requset.URL.Path
	if restweb.GetAction(url, 2) == "password" {
		return false
	}
	Cid := restweb.GetAction(url, 1)
	cid, err := strconv.Atoi(Cid)
	if err != nil {
		return true
	}

	contestModel := model.ContestModel{}
	ContestDetail, err := contestModel.Detail(cid)
	if err != nil {
		return true
	}
	prvs := "0" + ctx.GetSession("Privilege")
	prv, err := strconv.Atoi(prvs)
	if err != nil {
		restweb.Logger.Debug(err)
		return true
	}
	Uid := ctx.GetSession("Uid")

	restweb.Logger.Debug(Uid, prv)
	if prv < config.PrivilegeTC {
		if time.Now().Unix() < ContestDetail.Start || ContestDetail.Status == config.StatusReverse {
			info := "The contest has not started yet"
			if ContestDetail.Status == config.StatusReverse {
				info = "No such contest"
			}
			ctx.Error(info, http.StatusForbidden)
			return true
		} else if ContestDetail.Encrypt == config.EncryptPW {
			if Uid == "" {
				ctx.Redirect("/sess", http.StatusFound)
			} else if ctx.GetSession(Cid+"pass") != ContestDetail.Argument.(string) {
				ctx.Redirect("/contests/"+Cid+"/password", http.StatusFound)
			} else {
				return false
			}
			return true
		} else if ContestDetail.Encrypt == config.EncryptPT {
			if Uid == "" {
				ctx.Redirect("/sess", http.StatusFound)
				return true
			} else {
				userlist := strings.Split(ContestDetail.Argument.(string), "\n")
				flag := false
				for _, user := range userlist {
					if user == Uid {
						flag = true
						break
					}
				}
				if flag == false {
					ctx.Error(
						"Sorry, the contest is private and you are not granted to participate in the contest.", http.StatusForbidden)
					return true
				}
			}
		}
	}
	return false
}
