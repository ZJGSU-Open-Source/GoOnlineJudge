package class

import (
	"GoOnlineJudge/config"
	"encoding/json"
	"io"
	//"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Controller struct {
	Data      map[string]interface{}
	Uid       string
	Privilege int
}

func (this *Controller) Init(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	this.Data = make(map[string]interface{})

	this.Uid = this.GetSession(w, r, "CurrentUser")
	if this.Uid != "" {
		this.Data["IsCurrentUser"] = true
		this.Data["CurrentUser"] = this.Uid

		var err error
		this.Privilege, err = strconv.Atoi(this.GetSession(w, r, "CurrentPrivilege"))
		if err != nil {
			http.Error(w, "args error", 400)
			return
		}

		if this.Privilege > config.PrivilegeSB {
			this.Data["IsShowAdmin"] = true
		}
		if this.Privilege == config.PrivilegeAD {
			this.Data["IsAdmin"] = true
		}
	}
}

func (this *Controller) ParseURL(url string) (args map[string]string) {
	args = make(map[string]string)
	path := strings.Trim(url, "/")
	list := strings.Split(path, "/")

	for i := 1; i < len(list); i += 2 {
		args[list[i-1]] = list[i]
	}
	return
}

func (this *Controller) GetTime() (ft string) {
	t := time.Now().Unix()
	ft = time.Unix(t, 0).Format("2006-01-02 15:04:05")
	return
}

func (this *Controller) LoadJson(r io.Reader, i interface{}) (err error) {
	err = json.NewDecoder(r).Decode(i)
	return
}

func (this *Controller) PostReader(i interface{}) (r io.Reader, err error) {
	b, err := json.Marshal(i)
	if err != nil {
		return
	}
	r = strings.NewReader(string(b))
	return
}

func (this *Controller) SetSession(w http.ResponseWriter, r *http.Request, name string, value string) {
	s := Session{
		Name:  name,
		Value: value,
	}
	s.Set(w, r)
}

func (this *Controller) GetSession(w http.ResponseWriter, r *http.Request, name string) (value string) {
	s := Session{
		Name: name,
	}
	value = s.Get(w, r)
	return
}

func (this *Controller) DeleteSession(w http.ResponseWriter, r *http.Request, name string) {
	s := Session{
		Name: name,
	}
	s.Delete(w, r)
}

func (this *Controller) GetPage(page int, pageCount int) (ret map[string]interface{}) {
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

func (this *Controller) GetCodeLen(strLen int) (codeLen int) {
	codeLen = strLen
	return
}
