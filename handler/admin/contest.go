package admin

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/middleware"
	"GoOnlineJudge/model"

	"github.com/zenazn/goji/web"

	"net/http"
	"strconv"
	"strings"
	"time"
)

// 插入比赛
//@URL:/api/contests/ @method:POST
func PostContest(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD || user.Privilege != config.PrivilegeTC {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	one := contest(user, r)
	contestModel := model.ContestModel{}
	err := contestModel.Insert(one)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)

}

//更改contest状态
//@URL:/api/contests/:cid/status/ @method:PUT
func StatusContest(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		Cid  = c.URLParams["cid"]
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD || user.Privilege != config.PrivilegeTC {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	contestModel := model.ContestModel{}
	one, _ := contestModel.Detail(cid)
	if one.Creator != user.Uid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var status int
	switch one.Status {
	case config.StatusAvailable:
		status = config.StatusReverse
	default:
		status = config.StatusAvailable
	}

	err = contestModel.Status(cid, status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

//删除竞赛
//@URL: /api/contests/:cid/ @method:DELETE
func DeleteContest(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		Cid  = c.URLParams["cid"]
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD || user.Privilege != config.PrivilegeTC {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	contestModel := model.ContestModel{}
	contestModel.Detail(cid)

	err = contestModel.Delete(cid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

// 更新竞赛
//@URL:/api/contests/:cid @method:PUT
func PutContest(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		Cid  = c.URLParams["cid"]
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD || user.Privilege != config.PrivilegeTC {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	contestModel := model.ContestModel{}
	old, _ := contestModel.Detail(cid)
	if old.Creator != user.Uid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	one := contest(user, r)

	err = contestModel.Update(cid, one)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}

func contest(user *model.User, r *http.Request) (one model.Contest) {

	one.Title = r.FormValue("title")
	year, _ := strconv.Atoi(r.FormValue("startTimeYear"))
	month, _ := strconv.Atoi(r.FormValue("startTimeMonth"))
	day, _ := strconv.Atoi(r.FormValue("startTimeDay"))
	hour, _ := strconv.Atoi(r.FormValue("startTimeHour"))
	min, _ := strconv.Atoi(r.FormValue("startTimeMinute"))
	start := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one.Start = start.Unix()

	year, _ = strconv.Atoi(r.FormValue("endTimeYear"))
	month, _ = strconv.Atoi(r.FormValue("endTimeMonth"))
	day, _ = strconv.Atoi(r.FormValue("endTimeDay"))
	hour, _ = strconv.Atoi(r.FormValue("endTimeHour"))
	min, _ = strconv.Atoi(r.FormValue("endTimeMinute"))
	end := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one.End = end.Unix()

	if start.After(end) {
		return
	}

	switch r.FormValue("type") {
	case "public":
		one.Encrypt = config.EncryptPB
	case "private": //TODO 设置argument为一个string数组
		one.Encrypt = config.EncryptPT
		argument := r.FormValue("userlist")
		var cr rune = 13
		crStr := string(cr)
		argument = strings.Trim(argument, crStr)
		argument = strings.Trim(argument, "\r\n")
		one.Argument = argument
	case "password":
		one.Encrypt = config.EncryptPW
		one.Argument = r.FormValue("password")
	default:
		return
	}

	problemString := r.FormValue("problemList")
	problemString = strings.Trim(problemString, " ")
	problemString = strings.Trim(problemString, ";")
	problemList := strings.Split(problemString, ";")
	var list []int
	for _, v := range problemList {
		pid, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		problemModel := model.ProblemModel{}
		_, err = problemModel.Detail(pid) //检查题目是否存在
		if err != nil {
			continue
		}
		list = append(list, pid)
	}
	one.List = list
	one.Creator = user.Uid
	return one
}
