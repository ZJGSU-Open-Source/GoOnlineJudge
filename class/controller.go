package class

import (
	"GoOnlineJudge/config"
	"net/http"
	"restweb"
	"strconv"
)

type Controller struct {
	restweb.Controller
	Uid       string
	Privilege int
}

func (ct *Controller) Init(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	ct.Data = make(map[string]interface{})

	session := restweb.SessionManager.StartSession(w, r)
	ct.Uid = session.Get("Uid")

	ct.Data["CurrentUser"] = ct.Uid
	ct.Data["Privilege"] = ct.Privilege

	if ct.Uid != "" {
		ct.Data["IsCurrentUser"] = true
		var err error
		ct.Privilege, err = strconv.Atoi(session.Get("Privilege"))
		if err != nil {
			http.Error(w, "args error", 400)
			return
		}
		if ct.Privilege == config.PrivilegeAD {
			ct.Data["IsShowAdmin"] = true
			ct.Data["IsAdmin"] = true
			ct.Data["RejudgePrivilege"] = true
		} else if ct.Privilege == config.PrivilegeTC {
			ct.Data["IsShowTeacher"] = true
			ct.Data["IsTeacher"] = true
			ct.Data["RejudgePrivilege"] = true
		}
	}
}

func (ct *Controller) Err400(w http.ResponseWriter, r *http.Request, title string, info string) {
	restweb.Logger.Info(r.RemoteAddr + " " + ct.Uid)
	ct.Data["Title"] = title
	ct.Data["Info"] = info
	ct.Execute(w, "view/layout.tpl", "view/400.tpl")
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
