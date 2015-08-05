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
	Uid       string
	Privilege int
}

func (ct *Controller) Init() {
	ct.W.Header().Set("Content-Type", "text/html")

	ct.Uid = ct.GetSession("Uid")

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

func (ct *Controller) GetPage(page int, pageCount int) (ret map[string]interface{}) {
	ret = make(map[string]interface{})
	if page > 1 {
		ret["IsPreviousPage"] = true
	}
	if page < pageCount {
		ret["IsNextPage"] = true
	}

	var firstBlock bool = (page-config.PageMidLimit > config.PageHeadLimit+1)
	var secondBlock bool = (page+config.PageMidLimit < pageCount-config.PageTailLimit)

	if firstBlock && secondBlock {
		ret["IsPageHead"] = true
		s1 := make([]int, 0, 0)
		for i := 1; i <= config.PageHeadLimit; i++ {
			s1 = append(s1, i)
		}
		ret["PageHeadList"] = s1
		ret["IsPageMid"] = true
		s2 := make([]int, 0, 0)
		for i := page - config.PageMidLimit; i <= page+config.PageMidLimit; i++ {
			s2 = append(s2, i)
		}
		ret["PageMidList"] = s2
		ret["IsPageTail"] = true
		s3 := make([]int, 0, 0)
		for i := pageCount - config.PageTailLimit + 1; i <= pageCount; i++ {
			s3 = append(s3, i)
		}
		ret["PageTailList"] = s3
	} else if !firstBlock && !secondBlock {
		ret["IsPageHead"] = true
		s := make([]int, 0, 0)
		for i := 1; i <= pageCount; i++ {
			s = append(s, i)
		}
		ret["PageHeadList"] = s
	} else if firstBlock && !secondBlock {
		ret["IsPageHead"] = true
		s1 := make([]int, 0, 0)
		for i := 1; i <= config.PageHeadLimit; i++ {
			s1 = append(s1, i)
		}
		ret["PageHeadList"] = s1
		ret["IsPageMid"] = true
		s2 := make([]int, 0, 0)
		for i := page - config.PageMidLimit; i <= pageCount; i++ {
			s2 = append(s2, i)
		}
		ret["PageMidList"] = s2
	} else {
		ret["IsPageHead"] = true
		s1 := make([]int, 0, 0)
		for i := 1; i <= page+config.PageMidLimit; i++ {
			s1 = append(s1, i)
		}
		ret["PageHeadList"] = s1
		ret["IsPageTail"] = true
		s2 := make([]int, 0, 0)
		for i := pageCount - config.PageTailLimit + 1; i <= pageCount; i++ {
			s2 = append(s2, i)
		}
		ret["PageTailList"] = s2
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
