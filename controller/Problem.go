package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type problem struct {
	Pid int `json:"pid"bson:"pid"`

	Time    int    `json:"time"bson:"time"`
	Memory  int    `json:"memory"bson:"memory"`
	Special int    `json:"special"bson:"special"`
	Expire  string `json:"expire"bson:"expire"`

	Title       string        `json:"title"bson:"title"`
	Description template.HTML `json:"description"bson:"description"`
	Input       template.HTML `json:"input"bson:"input"`
	Output      template.HTML `json:"output"bson:"output"`
	Source      string        `json:"source"bson:"source"`
	Hint        string        `json:"hint"bson:"hint"`

	In  string `json:"in"bson:"in"`
	Out string `json:"out"bson:"out"`

	Solve  int `json:"solve"bson:"solve"`
	Submit int `json:"submit"bson:"submit"`

	Status int    `json:"status"bson:"status"`
	Create string `json:"create"bson:"create"`
}

type ProblemController struct {
	class.Controller
}

func (this *ProblemController) List(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem List")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path)
	url := "/problem/list"
	searchUrl := ""

	// Search
	if v, ok := args["pid"]; ok {
		searchUrl += "/pid/" + v
		this.Data["SearchPid"] = true
		this.Data["SearchValue"] = v
	}
	if v, ok := args["title"]; ok {
		searchUrl += "/title/" + v
		this.Data["SearchTitle"] = true
		this.Data["SearchValue"] = v
	}
	if v, ok := args["source"]; ok {
		searchUrl += "/source/" + v
		this.Data["SearchSource"] = true
		this.Data["SearchValue"] = v
	}
	url += searchUrl
	this.Data["URL"] = url

	// Page
	if _, ok := args["page"]; !ok {
		args["page"] = "1"
	}

	response, err := http.Post(config.PostHost+"/problem/count"+searchUrl, "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

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
	var pageCount = (count-1)/config.ProblemPerPage + 1
	page, err := strconv.Atoi(args["page"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	if page > pageCount {
		http.Error(w, "args error", 400)
		return
	}
	url += "/offset/" + strconv.Itoa((page-1)*config.ProblemPerPage) + "/limit/" + strconv.Itoa(config.ProblemPerPage)
	pageData := this.GetPage(page, pageCount)
	for k, v := range pageData {
		this.Data[k] = v
	}

	//
	response, err = http.Post(config.PostHost+url, "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	one := make(map[string][]problem)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["Problem"] = one["list"]
	}

	funcMap := map[string]interface{}{
		"ShowRatio":  class.ShowRatio,
		"ShowStatus": class.ShowStatus,
		"ShowExpire": class.ShowExpire,
		"NumEqual":   class.NumEqual,
		"NumAdd":     class.NumAdd,
		"NumSub":     class.NumSub,
	}
	t := template.New("layout.tpl").Funcs(funcMap)
	t, err = t.ParseFiles("view/layout.tpl", "view/problem_list.tpl")
	if err != nil {
		http.Error(w, "tpl1 error", 500)
		return
	}

	this.Data["Time"] = this.GetTime()
	this.Data["Title"] = "Problem List"
	this.Data["IsProblem"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem Detail")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path[2:])
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/problem/detail/pid/"+strconv.Itoa(pid), "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	var one problem
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["Detail"] = one
	}

	t := template.New("layout.tpl").Funcs(template.FuncMap{"ShowRatio": class.ShowRatio, "ShowSpecial": class.ShowSpecial})
	t, err = t.ParseFiles("view/layout.tpl", "view/problem_detail.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Problem Detial " + strconv.Itoa(pid)
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

// URL /problem/submit/pid/<pid>

func (this *ProblemController) Submit(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem Submit")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path)
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	uid := this.Uid
	if uid == "" {
		w.WriteHeader(401)
		return
	}

	one := make(map[string]interface{})
	one["pid"] = pid
	one["uid"] = uid
	one["module"] = config.ModuleP
	one["mid"] = config.ModuleP
	/////TODO. Judge
	one["judge"] = config.JudgeAC
	one["time"] = 1000
	one["memory"] = 888
	action := "submit"
	if one["judge"] == config.JudgeAC { //Judge whether the solution is accepted
		action = "solve"
	}
	response, err := http.Post(config.PostHost+"/problem/record/pid/"+strconv.Itoa(pid)+"/action/"+action, "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	response, err = http.Post(config.PostHost+"/user/record/uid/"+uid+"/action/"+action, "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	/////
	one["code"] = r.FormValue("code")
	one["len"] = this.GetCodeLen(len(r.FormValue("code")))
	one["language"], _ = strconv.Atoi(r.FormValue("compiler_id"))

	reader, err := this.PostReader(&one)
	if err != nil {
		http.Error(w, "read error", 500)
		return
	}

	response, err = http.Post(config.PostHost+"/solution/insert", "application/json", reader)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	w.WriteHeader(200)
}
