package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type contest struct {
	Cid      int         `json:"cid"bson:"cid"`
	Title    string      `json:"title"bson:"title"`
	Encrypt  int         `json:"encrypt"bson:"encrypt"`
	Argument interface{} `json:"argument"bson:"argument"`
	Type     string      `json:"type"bson:"type"` //the type of contest,acm contest or normal exercise

	Start int64 `json:"start"bson:"start"`
	End   int64 `json:"end"bson:"end"`

	Status int    `json:"status"bson:"status"`
	Create string `'json:"create"bson:"create"`

	List []int `json:"list"bson:"list"`
}

type ContestController struct {
	Cid           int
	ContestDetail *contest
	Index         map[int]int
	class.Controller
}

// url:/admin/contest/list/type/<contest,exercise>
func (this *ContestController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest List")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	Type := args["type"]

	response, err := http.Post(config.PostHost+"/contest?list/type?"+Type, "application", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	one := make(map[string][]*contest)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["Contest"] = one["list"]
	}

	this.Data["Title"] = "Admin - " + strings.Title(Type) + " List"
	this.Data["Is"+strings.Title(Type)] = true
	this.Data["IsList"] = true
	err = this.Execute(w, "view/admin/layout.tpl", "view/admin/contest_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

}

// url:/admin/contest/add/type/<contest,exercise>
func (this *ContestController) Add(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Contest Add")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	Type := args["type"]
	//class.Logger.Debug(Type)

	this.Data["Title"] = "Admin - " + strings.Title(Type) + " Add"
	this.Data["Is"+strings.Title(Type)] = true
	this.Data["IsAdd"] = true
	this.Data["Type"] = Type

	err := this.Execute(w, "view/admin/layout.tpl", "view/admin/contest_add.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

// url:/admin/contest?insert/type?<contest,exercise>
func (this *ContestController) Insert(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Contest Insert")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	Type := args["type"]

	one := make(map[string]interface{})
	one["title"] = r.FormValue("title")
	one["type"] = Type
	//class.Logger.Debug(one["type"])
	year, err := strconv.Atoi(r.FormValue("startTimeYear"))
	month, err := strconv.Atoi(r.FormValue("startTimeMonth"))
	day, err := strconv.Atoi(r.FormValue("startTimeDay"))
	hour, err := strconv.Atoi(r.FormValue("startTimeHour"))
	min, err := strconv.Atoi(r.FormValue("startTimeMinute"))
	start := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one["start"] = start.Unix()

	year, err = strconv.Atoi(r.FormValue("endTimeYear"))
	month, err = strconv.Atoi(r.FormValue("endTimeMonth"))
	day, err = strconv.Atoi(r.FormValue("endTimeDay"))
	hour, err = strconv.Atoi(r.FormValue("endTimeHour"))
	min, err = strconv.Atoi(r.FormValue("endTimeMinute"))
	end := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one["end"] = end.Unix()

	if start.After(end) {
		http.Error(w, "args error", 400)
		return
	}

	switch r.FormValue("type") {
	case "public":
		one["encrypt"] = config.EncryptPB
	case "private":
		one["encrypt"] = config.EncryptPT
		one["argument"] = r.FormValue("userlist")
	case "password":
		one["encrypt"] = config.EncryptPW
		one["argument"] = r.FormValue("password")
	default:
		http.Error(w, "args error", 400)
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
			http.Error(w, "conv error", 400)
			return
		}
		list = append(list, pid)
	}
	one["list"] = list //problemList 建议检查下problem是否存在，存在的将其在普通列表中不可见

	reader, err := this.PostReader(&one)
	if err != nil {
		http.Error(w, "read error", 500)
		return
	}

	response, err := http.Post(config.PostHost+"/contest?insert", "application/json", reader)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	ret := make(map[string]interface{})
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &ret)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		http.Redirect(w, r, "/admin/contest?list/type?"+Type, http.StatusFound)
	}
}

// url:/admin/contest/status/
func (this *ContestController) Status(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Contest Status")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	cid, err := strconv.Atoi(args["cid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/contest?detail/cid?"+strconv.Itoa(cid), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	var one contest
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
	}

	Type := one.Type

	var action int
	switch one.Status {
	case config.StatusAvailable:
		action = config.StatusReverse
	default:
		action = config.StatusAvailable
	}

	response, err = http.Post(config.PostHost+"/contest?status/cid?"+strconv.Itoa(cid)+"/action?"+strconv.Itoa(action), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		http.Redirect(w, r, "/admin/contest?list/type?"+strings.Title(Type), http.StatusFound)
	}
}

// url:/admin/contest/delete/
func (this *ContestController) Delete(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Contest Delete")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	cid, err := strconv.Atoi(args["cid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/contest?delete/cid?"+strconv.Itoa(cid), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	w.WriteHeader(response.StatusCode)
}

//// url:/admin/contest/edit/
func (this *ContestController) Edit(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Contest Edit")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	cid, err := strconv.Atoi(args["cid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/contest?detail/cid?"+strconv.Itoa(cid), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	var one struct {
		contest
		StartTimeYear   int
		StartTimeMonth  int
		StartTimeDay    int
		StartTimeHour   int
		StartTimeMinute int
		EndTimeYear     int
		EndTimeMonth    int
		EndTimeDay      int
		EndTimeHour     int
		EndTimeMinute   int
		ProblemList     string
		IsPublic        bool
		IsPrivate       bool
		IsPassword      bool
	}
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		start := time.Unix(one.Start, 0).Local()
		one.StartTimeYear = start.Year()
		one.StartTimeMonth = int(start.Month())
		one.StartTimeDay = start.Day()
		one.StartTimeHour = start.Hour()
		one.StartTimeMinute = start.Minute()

		end := time.Unix(one.End, 0).Local()
		one.EndTimeYear = end.Year()
		one.EndTimeMonth = int(end.Month())
		one.EndTimeDay = end.Day()
		one.EndTimeHour = end.Hour()
		one.EndTimeMinute = end.Minute()
		one.ProblemList = ""
		for _, v := range one.List {
			one.ProblemList += strconv.Itoa(v) + ";"
		}
		one.IsPublic = false
		one.IsPrivate = false
		one.IsPassword = false
		switch one.Encrypt {
		case config.EncryptPB:
			one.IsPublic = true
		case config.EncryptPT:
			one.IsPrivate = true
		case config.EncryptPW:
			one.IsPassword = true
		}
		this.Data["Detail"] = one
	} else {
		http.Error(w, "resp error", 500)
		return
	}

	Type := one.Type
	this.Data["Title"] = "Admin - " + strings.Title(Type) + " Edit"
	this.Data["Is"+strings.Title(Type)] = true
	this.Data["IsEdit"] = true

	err = this.Execute(w, "view/admin/layout.tpl", "view/admin/contest_edit.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

// url:/admin/contest/update/
func (this *ContestController) Update(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Contest Update")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	cid, err := strconv.Atoi(args["cid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	Type := args["type"]

	one := make(map[string]interface{})
	one["title"] = r.FormValue("title")
	one["type"] = Type
	year, _ := strconv.Atoi(r.FormValue("startTimeYear"))
	month, _ := strconv.Atoi(r.FormValue("startTimeMonth"))
	day, _ := strconv.Atoi(r.FormValue("startTimeDay"))
	hour, _ := strconv.Atoi(r.FormValue("startTimeHour"))
	min, _ := strconv.Atoi(r.FormValue("startTimeMinute"))

	start := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one["start"] = start.Unix()

	year, _ = strconv.Atoi(r.FormValue("endTimeYear"))
	month, _ = strconv.Atoi(r.FormValue("endTimeMonth"))
	day, _ = strconv.Atoi(r.FormValue("endTimeDay"))
	hour, _ = strconv.Atoi(r.FormValue("endTimeHour"))
	min, _ = strconv.Atoi(r.FormValue("endTimeMinute"))
	end := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)
	one["end"] = end.Unix()

	if start.After(end) {
		http.Error(w, "this.Query error", 400)
		return
	}

	switch r.FormValue("type") {
	case "public":
		one["encrypt"] = config.EncryptPB
		one["argument"] = ""
	case "private":
		one["encrypt"] = config.EncryptPT
		one["argument"] = r.FormValue("userlist")
	case "password":
		one["encrypt"] = config.EncryptPW
		one["argument"] = r.FormValue("password")
	default:
		http.Error(w, "args error", 400)
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
			http.Error(w, "conv error", 400)
			return
		}
		list = append(list, pid)
	}
	one["list"] = list

	reader, err := this.PostReader(&one)
	if err != nil {
		http.Error(w, "read error", 500)
		return
	}

	response, err := http.Post(config.PostHost+"/contest?update/cid?"+strconv.Itoa(cid), "application/json", reader)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		http.Redirect(w, r, "/admin/contest?list/type?"+Type, http.StatusFound)
	}
}
