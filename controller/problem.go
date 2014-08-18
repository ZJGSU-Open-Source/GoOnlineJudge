package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"net/http"
	"os/exec"
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
	class.Logger.Debug(r.RemoteAddr + "visit Problem List")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	url := "/problem?list"
	searchUrl := ""

	// Search
	if v, ok := args["pid"]; ok {
		searchUrl += "/pid?" + v
		this.Data["SearchPid"] = true
		this.Data["SearchValue"] = v
	}
	if v, ok := args["title"]; ok {
		searchUrl += "/title?" + v
		this.Data["SearchTitle"] = true
		this.Data["SearchValue"] = v
	}
	if v, ok := args["source"]; ok {
		searchUrl += "/source?" + v
		this.Data["SearchSource"] = true
		this.Data["SearchValue"] = v
	}
	url += searchUrl
	this.Data["URL"] = url

	// Page
	if _, ok := args["page"]; !ok {
		args["page"] = "1"
	}

	response, err := http.Post(config.PostHost+"/problem?count"+searchUrl, "application/json", nil)
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
	url += "/offset?" + strconv.Itoa((page-1)*config.ProblemPerPage) + "/limit?" + strconv.Itoa(config.ProblemPerPage)
	pageData := this.GetPage(page, pageCount)
	for k, v := range pageData {
		this.Data[k] = v
	}

	//
	response, err = http.Post(config.PostHost+url, "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

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
		//"ShowExpire": class.ShowExpire,
		"NumEqual": class.NumEqual,
		"NumAdd":   class.NumAdd,
		"NumSub":   class.NumSub,
		"LargePU":  class.LargePU,
	}
	t := template.New("layout.tpl").Funcs(funcMap)
	t, err = t.ParseFiles("view/layout.tpl", "view/problem_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
	this.Data["Privilege"] = this.Privilege
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
	class.Logger.Debug("Problem Detail")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/problem?detail/pid?"+strconv.Itoa(pid), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	var one problem
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["Detail"] = one
	}
	if this.Privilege <= config.PrivilegePU && one.Status == config.StatusReverse {
		t := template.New("layout.tpl")
		t, err = t.ParseFiles("view/layout.tpl", "view/400.tpl")
		if err != nil {
			class.Logger.Debug(err)
			http.Error(w, "tpl error", 500)
			return
		}

		this.Data["Info"] = "No such problem"
		this.Data["Title"] = "No such problem"
		err = t.Execute(w, this.Data)
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
		return
	}

	t := template.New("layout.tpl").Funcs(template.FuncMap{
		"ShowRatio":   class.ShowRatio,
		"ShowSpecial": class.ShowSpecial,
		"ShowStatus":  class.ShowStatus,
		"LargePU":     class.LargePU})
	t, err = t.ParseFiles("view/layout.tpl", "view/problem_detail.tpl")
	if err != nil {
		class.Logger.Debug(err)
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Privilege"] = this.Privilege
	this.Data["Title"] = "Problem — " + strconv.Itoa(pid)
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

// URL /problem?submit/pid?<pid>
func (this *ProblemController) Submit(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Problem Submit")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path)
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	uid := this.Uid
	if uid == "" {
		http.Error(w, "need sign in", 401)
		return
	}

	one := make(map[string]interface{})
	one["pid"] = pid
	one["uid"] = uid
	one["module"] = config.ModuleP
	one["mid"] = config.ModuleP

	response, err := http.Post(config.PostHost+"/problem?detail/pid?"+strconv.Itoa(pid), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	var pro problem
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &pro)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
	}
	code := r.FormValue("code")
	one["code"] = code
	one["length"] = this.GetCodeLen(len(r.FormValue("code")))
	one["language"], _ = strconv.Atoi(r.FormValue("compiler_id"))

	if code == "" || pro.Pid == 0 || (pro.Status == config.StatusReverse && this.Privilege <= config.PrivilegePU) {
		switch {
		case pro.Pid == 0 || (pro.Status == config.StatusReverse && this.Privilege <= config.PrivilegePU):
			this.Data["Info"] = "No such problem"
		case code == "":
			this.Data["Info"] = "Your source code is too short"
		}
		this.Data["Title"] = "Problem — " + strconv.Itoa(pid)

		t := template.New("layout.tpl")
		t, err = t.ParseFiles("view/layout.tpl", "view/400.tpl")
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
		err = t.Execute(w, this.Data)
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
		return
	}
	one["status"] = config.StatusAvailable
	one["judge"] = config.JudgePD

	reader, err := this.PostReader(&one)
	if err != nil {
		http.Error(w, "read error", 500)
		return
	}

	response, err = http.Post(config.PostHost+"/solution?insert", "application/json", reader)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	sl := make(map[string]int)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &sl)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}

	}
	w.WriteHeader(200)

	go func() {
		cmd := exec.Command("./RunServer", "-sid", strconv.Itoa(sl["sid"]), "-time", strconv.Itoa(pro.Time), "-memory", strconv.Itoa(pro.Memory)) //Run Judge
		err = cmd.Run()
		if err != nil {
			class.Logger.Debug(err)
		}
	}()
}
