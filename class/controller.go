package class

import (
	"GoOnlineJudge/config"
	"html/template"
	"io/ioutil"
	"restweb"
	"strconv"
)

type Controller struct {
	restweb.Controller
	Uid         string
	AccessToken string
	Privilege   int
}

func (ct *Controller) Init() {
	ct.W.Header().Set("Content-Type", "text/html")

	ct.Uid = ct.GetSession("Uid")
	ct.AccessToken = ct.GetSession("AccessToken")

	ct.Output["CurrentUser"] = ct.Uid
	ct.Output["Privilege"] = ct.Privilege

	if ct.Uid != "" {
		ct.Output["IsCurrentUser"] = true
		var err error
		ct.Privilege, err = strconv.Atoi(ct.GetSession("Privilege"))
		if err != nil {
			ct.Error("args error", 400)
			return
		}
		if ct.Privilege == config.PrivilegeAD {
			ct.Output["IsShowAdmin"] = true
			ct.Output["IsAdmin"] = true
			ct.Output["RejudgePrivilege"] = true
		} else if ct.Privilege == config.PrivilegeTC {
			ct.Output["IsShowTeacher"] = true
			ct.Output["IsTeacher"] = true
			ct.Output["RejudgePrivilege"] = true
		}
	}

	b, err := ioutil.ReadFile("view/admin/msg.txt")
	if err != nil {
		restweb.Logger.Debug(err)
	}

	ct.Output["Msg"] = template.HTML(string(b))
}

func (ct *Controller) Err400(title string, info string) {
	restweb.Logger.Info(ct.R.RemoteAddr + " " + ct.Uid)
	ct.Output["Title"] = title
	ct.Output["Info"] = info
	ct.RenderTemplate("view/layout.tpl", "view/400.tpl")
}

func (ct *Controller) GetPage(page int, tail bool) (ret map[string]interface{}) {
	ret = make(map[string]interface{})

	if page > 1 {
		ret["IsPreviousPage"] = true
	}

	if !tail {
		ret["IsNextPage"] = true
	}

	ret["CurrentPage"] = int(page)
	return
}

func (ct *Controller) GetCodeLen(strLen int) (codeLen int) {
	codeLen = strLen
	return
}

func importPrvi(priv string) int {
	if priv == "Admin" {
		return config.ProFull | config.ConFull | config.NewsFull | config.AccessAdmin | config.Notice | config.UserFull | config.CodeFull | config.ViewReverse
	} else if priv == "Teacher" {
		return config.CodeFull | config.AddContest | config.ReJudge | config.GenerateUser | config.AccessAdmin
	}
	return 0
}
