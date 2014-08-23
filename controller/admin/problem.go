package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

	Status int   `json:"status"bson:"status"`
	Create int64 `json:"create"bson:"create"`
}

type ProblemController struct {
	class.Controller
}

func (this *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Problem Detail")
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
	} else {
		http.Error(w, "resp error", 500)
		return
	}

	this.Data["Title"] = "Admin - Problem Detail"
	this.Data["IsProblem"] = true
	this.Data["IsList"] = false

	err = this.Execute(w, "view/admin/layout.tpl", "view/problem_detail.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ProblemController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Problem List")
	this.Init(w, r)

	response, err := http.Post(config.PostHost+"/problem?list", "application/json", nil)
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

	this.Data["Title"] = "Admin - Problem List"
	this.Data["IsProblem"] = true
	this.Data["IsList"] = true
	err = this.Execute(w, "view/admin/layout.tpl", "view/admin/problem_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ProblemController) Add(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Problem Add")
	this.Init(w, r)

	this.Data["Title"] = "Admin - Problem Add"
	this.Data["IsProblem"] = true
	this.Data["IsAdd"] = true
	this.Data["IsEdit"] = true

	err := this.Execute(w, "view/admin/layout.tpl", "view/admin/problem_add.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ProblemController) Insert(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Problem Insert")
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

	in := r.FormValue("in")
	out := r.FormValue("out")

	one["description"] = r.FormValue("description")
	one["input"] = r.FormValue("input")
	one["output"] = r.FormValue("output")
	one["in"] = in
	one["out"] = out
	one["source"] = r.FormValue("source")
	one["hint"] = r.FormValue("hint")

	reader, err := this.PostReader(&one)
	if err != nil {
		http.Error(w, "read error", 500)
		return
	}

	response, err := http.Post(config.PostHost+"/problem?insert", "application/json", reader)
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
			class.Logger.Debug("create dir error")
			return
		}

		infile, err := os.Create(config.Datapath + strconv.Itoa(ret["pid"]) + "/sample.in")
		if err != nil {
			class.Logger.Debug(err)
		}
		defer infile.Close()
		var cr rune = 13
		crStr := string(cr)
		in = strings.Replace(in, "\r\n", "\n", -1)
		in = strings.Replace(in, crStr, "\n", -1)
		infile.WriteString(in)
		outfile, err := os.Create(config.Datapath + strconv.Itoa(ret["pid"]) + "/sample.out")
		if err != nil {
			class.Logger.Debug(err)
		}
		defer outfile.Close()
		out = strings.Replace(out, "\r\n", "\n", -1)
		out = strings.Replace(out, crStr, "\n", -1)
		outfile.WriteString(out)

		http.Redirect(w, r, "/admin/problem?list", http.StatusFound)
	}
}

func (this *ProblemController) Status(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Problem Status")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	//class.Logger.Debug(args)
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
	response, err = http.Post(config.PostHost+"/problem?status/pid?"+strconv.Itoa(pid)+"/action?"+strconv.Itoa(action), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		http.Redirect(w, r, "/admin/problem?list", http.StatusFound)
	}
}

func (this *ProblemController) Delete(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Problem Delete")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/problem?delete/pid?"+strconv.Itoa(pid), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	w.WriteHeader(response.StatusCode)
}

func (this *ProblemController) Edit(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Problem Edit")
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
	} else {
		http.Error(w, "resp error", 500)
		return
	}

	this.Data["Title"] = "Admin - Problem Edit"
	this.Data["IsProblem"] = true
	this.Data["IsList"] = false
	this.Data["IsEdit"] = true

	err = this.Execute(w, "view/admin/layout.tpl", "view/admin/problem_edit.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ProblemController) Update(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Problem Update")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
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

	var cr rune = 13
	crStr := string(cr)
	in := r.FormValue("in")
	out := r.FormValue("out")

	one["description"] = r.FormValue("description")
	one["input"] = r.FormValue("input")
	one["output"] = r.FormValue("output")
	one["in"] = in
	one["out"] = out
	one["source"] = r.FormValue("source")
	one["hint"] = r.FormValue("hint")

	infile, err := os.Create(config.Datapath + args["pid"] + "/sample.in")
	if err != nil {
		class.Logger.Debug(err)
	}
	defer infile.Close()
	in = strings.Replace(in, "\r\n", "\n", -1)
	in = strings.Replace(in, crStr, "\n", -1)
	infile.WriteString(in)

	outfile, err := os.Create(config.Datapath + args["pid"] + "/sample.out")
	if err != nil {
		class.Logger.Debug(err)
	}
	defer outfile.Close()
	out = strings.Replace(out, "\r\n", "\n", -1)
	out = strings.Replace(out, crStr, "\n", -1)
	outfile.WriteString(out)

	reader, err := this.PostReader(&one)
	if err != nil {
		http.Error(w, "read error", 500)
		return
	}

	response, err := http.Post(config.PostHost+"/problem?update/pid?"+strconv.Itoa(pid), "application/json", reader)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		http.Redirect(w, r, "/admin/problem?detail/pid?"+strconv.Itoa(pid), http.StatusFound)
	}
}

func (this *ProblemController) Rejudgepage(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Rejudge Page")
	this.Init(w, r)

	this.Data["Title"] = "Admin - Problem Rejudge"
	this.Data["IsAdmin"] = true
	this.Data["IsProblem"] = true
	this.Data["IsRejudge"] = true

	err := this.Execute(w, "view/admin/layout.tpl", "view/admin/problem_rejudge.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ProblemController) Rejudge(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Rejudge")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	id, err := strconv.Atoi(args["id"])
	types := args["type"]

	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	ok := 1
	hint := make(map[string]string)

	if types == "Pid" {
		response, err := http.Post(config.PostHost+"/problem?detail/pid?"+strconv.Itoa(id), "application/json", nil)
		if err != nil {
			http.Error(w, "post error", 500)
			return
		}
		defer response.Body.Close()
	} else if types == "Sid" {
		sid := id
		response, err := http.Post(config.PostHost+"/solution?detail/sid?"+strconv.Itoa(sid), "application/json", nil)
		if err != nil {
			class.Logger.Debug(err)
			return
		}
		defer response.Body.Close()

		var sol solution
		if response.StatusCode == 200 {
			err = this.LoadJson(response.Body, &sol)
			if err != nil {
				class.Logger.Debug(err)
				return
			}
		}

		//one := make(map[string]interface{})

		//one["pid"] = sol.Pid
		//one["uid"] = sol.Uid
		//one["module"] = config.ModuleP
		//one["mid"] = config.ModuleP
		//one["code"] = sol.Code
		//one["length"] = sol.Length
		//one["language"] = sol.Language
		//one["status"] = config.StatusAvailable
		//one["judge"] = config.JudgePD
		//one["judge"] = config.JudgeRPD

		response, err = http.Post(config.PostHost+"/problem?detail/pid?"+strconv.Itoa(sol.Pid), "application/json", nil)
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
		/*
			reader, err := this.PostReader(&one)
			if err != nil {
				http.Error(w, "read error", 500)
				return
			}


					response, err = http.Post(config.PostHost+"/solution?rejudge/sid?"+strconv.Itoa(sid), "application/json", reader)
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
		*/

		go func() {
			cmd := exec.Command("./RunServer", "-sid", strconv.Itoa(sid), "-time", strconv.Itoa(pro.Time), "-memory", strconv.Itoa(pro.Memory)) //Run Judge
			err = cmd.Run()
			if err != nil {
				class.Logger.Debug(err)
			}
		}()
	}

	if ok == 1 {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
	}

	b, err := json.Marshal(&hint)
	if err != nil {
		http.Error(w, "json error", 500)
		return
	}
	w.Write(b)
}
