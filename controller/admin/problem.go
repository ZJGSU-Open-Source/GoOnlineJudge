package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"log"
	"net/http"
	"os"
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

func (this *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Problem Detail")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path[6:])
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/problem/detail/pid/"+strconv.Itoa(pid), "application/json", nil)
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
	} else {
		http.Error(w, "resp error", 500)
		return
	}

	t := template.New("layout.tpl").Funcs(template.FuncMap{
		"ShowRatio":   class.ShowRatio,
		"ShowSpecial": class.ShowSpecial,
		"ShowStatus":  class.ShowStatus,
		"LargePU":     class.LargePU})
	t, err = t.ParseFiles("view/admin/layout.tpl", "view/problem_detail.tpl")
	if err != nil {
		log.Println(err)
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Admin - Problem Detail"
	this.Data["IsProblem"] = true
	this.Data["IsList"] = false

	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ProblemController) List(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Problem List")
	this.Init(w, r)

	response, err := http.Post(config.PostHost+"/problem/list", "application/json", nil)
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

	t := template.New("layout.tpl").Funcs(template.FuncMap{"ShowStatus": class.ShowStatus})
	t, err = t.ParseFiles("view/admin/layout.tpl", "view/admin/problem_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Admin - Problem List"
	this.Data["IsProblem"] = true
	this.Data["IsList"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ProblemController) Add(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Problem Add")
	this.Init(w, r)

	t := template.New("layout.tpl")
	t, err := t.ParseFiles("view/admin/layout.tpl", "view/admin/problem_add.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Admin - Problem Add"
	this.Data["IsProblem"] = true
	this.Data["IsAdd"] = true
	this.Data["IsEdit"] = true

	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ProblemController) Insert(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Problem Insert")
	this.Init(w, r)

	//TODO r.body to json
	one := make(map[string]interface{})
	one["title"] = r.FormValue("title")
	time, err := strconv.Atoi(r.FormValue("time"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	one["time"] = time
	memory, err := strconv.Atoi(r.FormValue("memory"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	one["memory"] = memory
	if special := r.FormValue("special"); special == "" {
		one["special"] = 0
	} else {
		one["special"] = 1
	}
	one["description"] = r.FormValue("description")
	one["input"] = r.FormValue("input")
	one["output"] = r.FormValue("output")
	one["in"] = r.FormValue("in")
	one["out"] = r.FormValue("out")
	one["source"] = r.FormValue("source")
	one["hint"] = r.FormValue("hint")

	reader, err := this.PostReader(&one)
	if err != nil {
		http.Error(w, "read error", 500)
		return
	}

	response, err := http.Post(config.PostHost+"/problem/insert", "application/json", reader)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	ret := make(map[string]int)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &ret)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}

		err = os.Mkdir(config.Datapath+strconv.Itoa(ret["pid"]), os.ModePerm)
		if err != nil {
			log.Println("create dir error")
			return
		}

		infile, err := os.Create(config.Datapath + strconv.Itoa(ret["pid"]) + "/sample.in")
		if err != nil {
			log.Println(err)
		}
		defer infile.Close()
		infile.WriteString(r.FormValue("in"))
		outfile, err := os.Create(config.Datapath + strconv.Itoa(ret["pid"]) + "/sample.out")
		if err != nil {
			log.Println(err)
		}
		defer outfile.Close()
		outfile.WriteString(r.FormValue("out"))

		http.Redirect(w, r, "/admin/problem/list", http.StatusFound)
	}
}

func (this *ProblemController) Status(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Problem Status")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path[6:])
	log.Println(args)
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/problem/detail/pid/"+strconv.Itoa(pid), "application/json", nil)
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
	} else {
		http.Error(w, "resp error", 500)
		return
	}
	var action int
	switch one.Status {
	case config.StatusAvailable:
		action = config.StatusReverse
	case config.StatusReverse:
		action = config.StatusAvailable
	}
	response, err = http.Post(config.PostHost+"/problem/status/pid/"+strconv.Itoa(pid)+"/action/"+strconv.Itoa(action), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		http.Redirect(w, r, "/admin/problem/list", http.StatusFound)
	}
}

func (this *ProblemController) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Problem Delete")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path[6:])
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/problem/delete/pid/"+strconv.Itoa(pid), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	w.WriteHeader(response.StatusCode)
}

func (this *ProblemController) Edit(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Problem Edit")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path[6:])
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/problem/detail/pid/"+strconv.Itoa(pid), "application/json", nil)
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
	} else {
		http.Error(w, "resp error", 500)
		return
	}

	t := template.New("layout.tpl").Funcs(template.FuncMap{"ShowRatio": class.ShowRatio, "ShowSpecial": class.ShowSpecial})
	t, err = t.ParseFiles("view/admin/layout.tpl", "view/admin/problem_edit.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Admin - Problem Edit"
	this.Data["IsProblem"] = true
	this.Data["IsList"] = false
	this.Data["IsEdit"] = true

	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ProblemController) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Problem Update")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path[6:])
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	one := make(map[string]interface{})
	one["title"] = r.FormValue("title")
	time, err := strconv.Atoi(r.FormValue("time"))
	if err != nil {
		http.Error(w, "conv error", 500)
		return
	}
	one["time"] = time
	memory, err := strconv.Atoi(r.FormValue("memory"))
	if err != nil {
		http.Error(w, "conv error", 500)
		return
	}
	one["memory"] = memory
	if special := r.FormValue("special"); special == "" {
		one["special"] = 0
	} else {
		one["special"] = 1
	}
	one["description"] = r.FormValue("description")
	one["input"] = r.FormValue("input")
	one["output"] = r.FormValue("output")
	one["in"] = r.FormValue("in")
	one["out"] = r.FormValue("out")
	one["source"] = r.FormValue("source")
	one["hint"] = r.FormValue("hint")

	infile, err := os.Create(config.Datapath + args["pid"] + "/sample.in")
	if err != nil {
		log.Println(err)
	}
	defer infile.Close()
	infile.WriteString(r.FormValue("in"))
	outfile, err := os.Create(config.Datapath + args["pid"] + "/sample.out")
	if err != nil {
		log.Println(err)
	}
	defer outfile.Close()
	outfile.WriteString(r.FormValue("out"))

	reader, err := this.PostReader(&one)
	if err != nil {
		http.Error(w, "read error", 500)
		return
	}

	response, err := http.Post(config.PostHost+"/problem/update/pid/"+strconv.Itoa(pid), "application/json", reader)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		http.Redirect(w, r, "/admin/problem/detail/pid/"+strconv.Itoa(pid), http.StatusFound)
	}
}
