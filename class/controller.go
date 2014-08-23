package class

import (
	"GoOnlineJudge/config"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Controller struct {
	Uid       string
	Privilege int
	Data      map[string]interface{}
}

func (this *Controller) Init(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	this.Data = make(map[string]interface{})

	session := SessionManager.StartSession(w, r)
	this.Uid = session.Get("Uid")

	this.Data["CurrentUser"] = this.Uid
	this.Data["Privilege"] = this.Privilege

	if this.Uid != "" {
		this.Data["IsCurrentUser"] = true
		var err error
		this.Privilege, err = strconv.Atoi(session.Get("Privilege"))
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

	for _, pair := range list {
		k_v := strings.Split(pair, "?")
		if len(k_v) <= 1 {
			args[k_v[0]] = "Index"
		} else {
			args[k_v[0]] = k_v[1]
		}
	}
	return
}

func (this *Controller) GetTime() (t int64) {
	t = time.Now().Unix()
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

func (this *Controller) SetSession(w http.ResponseWriter, r *http.Request, key string, value string) {
	session := SessionManager.StartSession(w, r)
	session.Set(key, value)
}

func (this *Controller) GetSession(w http.ResponseWriter, r *http.Request, key string) (value string) {
	session := SessionManager.StartSession(w, r)
	value = session.Get(key)
	return
}

func (this *Controller) DeleteSession(w http.ResponseWriter, r *http.Request) {
	SessionManager.DeleteSession(w, r)
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
func (c *Controller) Execute(w io.Writer, tplfiles ...string) error {
	t, err := ParseFiles(tplfiles...)
	if err == nil {
		err = t.Execute(w, c.Data)
	}
	return err
}
