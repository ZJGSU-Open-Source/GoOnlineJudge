package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type solution struct {
	Sid int `json:"sid"bson:"sid"`

	Pid      int    `json:"pid"bson:"pid"`
	Uid      string `json:"uid"bson:"uid"`
	Judge    int    `json:"judge"bson:"judge"`
	Time     int    `json:"time"bson:"time"`
	Memory   int    `json:"memory"bson:"memory"`
	Length   int    `json:"length"bson:"length"`
	Language int    `json:"language"bson:"language"`

	Module int `json:"module"bson:"module"`
	Mid    int `json:"mid"bson:"mid"`

	Code string `json:"code"bson:"code"`

	Status int    `json:"status"bson:"status"`
	Create string `json:"create"bson:"create"`
}

type StatusController struct {
	class.Controller
}

func (this *StatusController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Status List")
	this.Init(w, r)
	args := this.ParseURL(r.URL.String())
	url := "/solution?list"
	searchUrl := ""
	// Search
	if v, ok := args["uid"]; ok {
		searchUrl += "/uid?" + v
		this.Data["SearchUid"] = v
	}
	if v, ok := args["pid"]; ok {
		searchUrl += "/pid?" + v
		this.Data["SearchPid"] = v
	}
	if v, ok := args["judge"]; ok {
		searchUrl += "/judge?" + v
		this.Data["SearchJudge"+v] = v
	}
	if v, ok := args["language"]; ok {
		searchUrl += "/language?" + v
		this.Data["SearchLanguage"+v] = v
	}
	url += searchUrl
	this.Data["URL"] = "/status?list" + searchUrl

	// Page
	if _, ok := args["page"]; !ok {
		args["page"] = "1"
	}

	response, err := http.Post(config.PostHost+"/solution?count"+searchUrl+"/module?"+strconv.Itoa(config.ModuleP)+"/action?submit", "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	c := make(map[string]int)
	var count int
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &c)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		count = c["count"]
	}
	var pageCount = (count-1)/config.SolutionPerPage + 1

	page, err := strconv.Atoi(args["page"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	if page > pageCount {
		http.Error(w, "args error", 400)
		return
	}
	url += "/offset?" + strconv.Itoa((page-1)*config.SolutionPerPage) + "/limit?" + strconv.Itoa(config.SolutionPerPage)
	pageData := this.GetPage(page, pageCount)
	for k, v := range pageData {
		this.Data[k] = v
	}

	//
	response, err = http.Post(config.PostHost+url+"/module?"+strconv.Itoa(config.ModuleP), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	one := make(map[string][]solution)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["Solution"] = one["list"]
	}

	funcMap := map[string]interface{}{
		"ShowStatus":   class.ShowStatus,
		"ShowJudge":    class.ShowJudge,
		"ShowLanguage": class.ShowLanguage,
		"NumEqual":     class.NumEqual,
		"NumAdd":       class.NumAdd,
		"NumSub":       class.NumSub,
	}
	t := template.New("layout.tpl").Funcs(funcMap)
	t, err = t.ParseFiles("view/layout.tpl", "view/status_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Status List"
	this.Data["IsStatus"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		class.Logger.Debug(err.Error())
		e := err.Error()
		codefile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			class.Logger.Debug(err)
			return
		}
		defer codefile.Close()
		_, err = codefile.WriteString(e)
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *StatusController) Code(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Status Code")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	sid, err := strconv.Atoi(args["sid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/solution?detail/sid?"+strconv.Itoa(sid), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	var one solution
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["Solution"] = one
	}

	t := template.New("layout.tpl")
	t, err = t.ParseFiles("view/layout.tpl", "view/status_code.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "View Code"
	this.Data["IsCode"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
