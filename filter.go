package main

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"net/http"
	"restweb"
	"strconv"
	"strings"
	"time"
)

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
		if ctx.R.Method == restweb.GET {
			ctx.Redirect("/sess", http.StatusFound)
		} else {
			ctx.W.WriteHeader(http.StatusUnauthorized)
		}
		return true
	}
	return false
}

func requireContest(ctx *restweb.Context) bool {
	url := ctx.R.URL.Path
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
				userlist := strings.Split(ContestDetail.Argument.(string), "\r\n")
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
