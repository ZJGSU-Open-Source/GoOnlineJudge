package main

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"html/template"
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
		Err400(ctx.W, "Admin", "Error occured")
		return true
	}
	if prt < config.PrivilegeTC {
		Err400(ctx.W, "Admin", "No privilege")
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
			Err400(ctx.W, "Contest", info)
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
					Err400(ctx.W, "Contest",
						"Sorry, the contest is private and you are not granted to participate in the contest.")
					return true
				}
			}
		}
	}
	return false
}
func Err400(w http.ResponseWriter, title string, info string) {
	Output := make(map[string]interface{})
	Output["Title"] = title
	Output["Info"] = info
	t, err := template.ParseFiles("view/layout.tpl", "view/400.tpl")
	if err == nil {
		err = t.Execute(w, Output)
	}
}
