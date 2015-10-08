package admin

import (
	"ojapi/config"
	"ojapi/middleware"
	"ojapi/model"

	"github.com/zenazn/goji/web"

	"encoding/json"
	"errors"
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

	one, err := contest(user, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contestModel := model.ContestModel{}
	err = contestModel.Insert(*one)
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
		one  = middleware.ToContest(c)
		user = middleware.ToUser(c)
	)

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

	contestModel := model.ContestModel{}
	err := contestModel.Status(one.Cid, status)
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
		contest = middleware.ToContest(c)
		user    = middleware.ToUser(c)
	)

	if contest.Creator != user.Uid || user.Privilege != config.PrivilegeAD {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	contestModel := model.ContestModel{}
	err := contestModel.Delete(contest.Cid)
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
		old  = middleware.ToContest(c)
		user = middleware.ToUser(c)
	)

	if old.Creator != user.Uid {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	one, err := contest(user, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contestModel := model.ContestModel{}
	err = contestModel.Update(one.Cid, *one)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}

func contest(user *model.User, r *http.Request) (one *model.Contest, err error) {

	in := struct {
		Title       string
		Start       int64
		End         int64
		Type        string
		Password    string
		Userlist    string
		ProblemList string
	}{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		return nil, errors.New("Bad json payload")
	}

	one.Title = in.Title
	one.Start = in.Start
	one.End = in.End

	if time.Unix(one.Start, 0).After(time.Unix(one.End, 0)) {
		return nil, errors.New("Bad start or end time")
	}

	switch in.Type {
	case "public":
		one.Encrypt = config.EncryptPB
	case "private": //TODO 设置argument为一个string数组
		one.Encrypt = config.EncryptPT
		one.Argument = in.Userlist
	case "password":
		one.Encrypt = config.EncryptPW
		one.Argument = in.Password
	default:
		return nil, errors.New("Bad contest time")
	}

	problemString := in.ProblemList
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
	return one, nil
}
