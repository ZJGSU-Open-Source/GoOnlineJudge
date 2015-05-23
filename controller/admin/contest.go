package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"net/http"
	"strconv"
	"strings"
	"time"
)

//竞赛控件
type AdminContest struct {
	Cid           int
	ContestDetail *model.Contest
	Index         map[int]int
	class.Controller
} //@Controller

//列出所有的比赛
//@URL: /api/admin/contests/ @method:GET
func (cc *AdminContest) List() {
	restweb.Logger.Debug("Contest List")

	qry := make(map[string]string)
	contestModel := model.ContestModel{}
	contestList, err := contestModel.List(qry)
	if err != nil {
		cc.Error(err.Error(), 400)
	}

	cc.Output["Contest"] = contestList
	cc.RenderJson()
}

// 插入比赛
//@URL:/api/admin/contests/ @method:POST
func (cc *AdminContest) Insert() {
	restweb.Logger.Debug("Admin Contest Insert")

	one := cc.contest()
	contestModel := model.ContestModel{}
	err := contestModel.Insert(one)
	if err != nil {
		cc.Error(err.Error(), 500)
		return
	}

	cc.W.WriteHeader(201)

}

//更改contest状态
//@URL:/api/admin/contests/:cid/status/ @method:PUT
func (cc *AdminContest) Status(Cid string) {

	restweb.Logger.Debug("Admin Contest Status")

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		cc.Error("args error", 400)
		return
	}
	contestModel := model.ContestModel{}
	one, _ := contestModel.Detail(cid)
	if one.Creator != cc.Uid {
		cc.Error("privilege error", 400)
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
		cc.Error(err.Error(), 500)
		return
	}

	cc.W.WriteHeader(200)
}

//删除竞赛
//@URL: /api/admin/contests/(\d+)/ @method:DELETE
func (cc *AdminContest) Delete(Cid string) {

	restweb.Logger.Debug("Admin Contest Delete")

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		cc.Error("args error", 400)
		return
	}
	contestModel := model.ContestModel{}
	old, _ := contestModel.Detail(cid)

	err = contestModel.Delete(cid)
	if err != nil {
		cc.Error(err.Error(), 400)
		return
	}
	cc.W.WriteHeader(200)
}

// 更新竞赛
//@URL:/api/admin/contests/(\d+)/ @method:PUT
func (cc *AdminContest) Update(Cid string) {

	restweb.Logger.Debug("Admin Contest Update")

	cid, err := strconv.Atoi(Cid)
	if err != nil {
		cc.Error("args error", 400)
		return
	}
	contestModel := model.ContestModel{}
	old, _ := contestModel.Detail(cid)
	if old.Creator != cc.Uid {
		cc.Error("privilege error", 400)
		return
	}

	one := cc.contest()

	err = contestModel.Update(cid, one)
	if err != nil {
		cc.Error(err.Error(), 400)
		return
	}
	cc.W.WriteHeader(200)
}

func (cc *AdminContest) contest() (one model.Contest) {

	one.Title = cc.Input.Get("title")
	year, _ := strconv.Atoi(cc.Input.Get("startTimeYear"))
	month, _ := strconv.Atoi(cc.Input.Get("startTimeMonth"))
	day, _ := strconv.Atoi(cc.Input.Get("startTimeDay"))
	hour, _ := strconv.Atoi(cc.Input.Get("startTimeHour"))
	min, _ := strconv.Atoi(cc.Input.Get("startTimeMinute"))
	start := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one.Start = start.Unix()

	year, _ = strconv.Atoi(cc.Input.Get("endTimeYear"))
	month, _ = strconv.Atoi(cc.Input.Get("endTimeMonth"))
	day, _ = strconv.Atoi(cc.Input.Get("endTimeDay"))
	hour, _ = strconv.Atoi(cc.Input.Get("endTimeHour"))
	min, _ = strconv.Atoi(cc.Input.Get("endTimeMinute"))
	end := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one.End = end.Unix()

	if start.After(end) {
		cc.Error("args error", 400)
		return
	}

	switch cc.Input.Get("type") {
	case "public":
		one.Encrypt = config.EncryptPB
	case "private": //TODO 设置argument为一个string数组
		one.Encrypt = config.EncryptPT
		argument := cc.Input.Get("userlist")
		var cr rune = 13
		crStr := string(cr)
		argument = strings.Trim(argument, crStr)
		argument = strings.Trim(argument, "\r\n")
		one.Argument = argument
		restweb.Logger.Debug(one.Argument)
	case "password":
		one.Encrypt = config.EncryptPW
		one.Argument = cc.Input.Get("password")
	default:
		cc.Error("args error", 400)
		return
	}

	problemString := cc.Input.Get("problemList")
	problemString = strings.Trim(problemString, " ")
	problemString = strings.Trim(problemString, ";")
	problemList := strings.Split(problemString, ";")
	var list []int
	for _, v := range problemList {
		pid, err := strconv.Atoi(v)
		if err != nil {
			restweb.Logger.Debug(err)
			continue
		}
		problemModel := model.ProblemModel{}
		_, err = problemModel.Detail(pid) //检查题目是否存在
		if err != nil {
			restweb.Logger.Debug(err)
			continue
		}
		list = append(list, pid)
	}
	one.List = list
	one.Creator = cc.Uid
	restweb.Logger.Debug(one.Creator)
	return one
}
