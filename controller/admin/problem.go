package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

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

	problemModel := model.ProblemModel{}
	one, err := problemModel.Detail(pid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	this.Data["Detail"] = one
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

	problemModel := model.ProblemModel{}
	qry := make(map[string]string)
	proList, err := problemModel.List(qry)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	this.Data["Problem"] = proList
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

	one := model.Problem{}
	one.Title = r.FormValue("title")
	time, err := strconv.Atoi(r.FormValue("time"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	one.Time = time
	memory, err := strconv.Atoi(r.FormValue("memory"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	one.Memory = memory
	if special := r.FormValue("special"); special == "" {
		one.Special = 0
	} else {
		one.Special = 1
	}

	in := r.FormValue("in")
	out := r.FormValue("out")
	one.Description = template.HTML(r.FormValue("description"))
	one.Input = template.HTML(r.FormValue("input"))
	one.Output = template.HTML(r.FormValue("output"))
	one.In = in
	one.Out = out
	one.Source = r.FormValue("source")
	one.Hint = r.FormValue("hint")

	problemModel := model.ProblemModel{}
	pid, err := problemModel.Insert(one)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = os.Mkdir(config.Datapath+strconv.Itoa(pid), os.ModePerm)
	if err != nil {
		class.Logger.Debug("create dir error")
		return
	}

	infile, err := os.Create(config.Datapath + strconv.Itoa(pid) + "/sample.in")
	if err != nil {
		class.Logger.Debug(err)
	}
	defer infile.Close()
	var cr rune = 13
	crStr := string(cr)
	in = strings.Replace(in, "\r\n", "\n", -1)
	in = strings.Replace(in, crStr, "\n", -1)
	infile.WriteString(in)
	outfile, err := os.Create(config.Datapath + strconv.Itoa(pid) + "/sample.out")
	if err != nil {
		class.Logger.Debug(err)
	}
	defer outfile.Close()
	out = strings.Replace(out, "\r\n", "\n", -1)
	out = strings.Replace(out, crStr, "\n", -1)
	outfile.WriteString(out)

	http.Redirect(w, r, "/admin/problem?list", http.StatusFound)
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

	problemModel := model.ProblemModel{}
	one, err := problemModel.Detail(pid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	this.Data["Detail"] = one
	var status int
	switch one.Status {
	case config.StatusAvailable:
		status = config.StatusReverse
	case config.StatusReverse:
		status = config.StatusAvailable
	}
	err = problemModel.Status(pid, status)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/admin/problem?list", http.StatusFound)
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

	problemModel := model.ProblemModel{}
	problemModel.Delete(pid)

	w.WriteHeader(200)
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

	problemModel := model.ProblemModel{}
	one, err := problemModel.Detail(pid)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	this.Data["Detail"] = one
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

	one := model.Problem{}
	one.Title = r.FormValue("title")
	time, err := strconv.Atoi(r.FormValue("time"))
	if err != nil {
		http.Error(w, "conv error", 500)
		return
	}
	one.Time = time
	memory, err := strconv.Atoi(r.FormValue("memory"))
	if err != nil {
		http.Error(w, "conv error", 500)
		return
	}
	one.Memory = memory
	if special := r.FormValue("special"); special == "" {
		one.Special = 0
	} else {
		one.Special = 1
	}

	var cr rune = 13
	crStr := string(cr)
	in := r.FormValue("in")
	out := r.FormValue("out")

	one.Description = template.HTML(r.FormValue("description"))
	one.Input = template.HTML(r.FormValue("input"))
	one.Output = template.HTML(r.FormValue("output"))
	one.In = in
	one.Out = out
	one.Source = r.FormValue("source")
	one.Hint = r.FormValue("hint")

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

	problemModel := model.ProblemModel{}
	err = problemModel.Update(pid, one)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/admin/problem?detail/pid?"+strconv.Itoa(pid), http.StatusFound)
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

		solutionModel := model.SolutionModel{}
		sol, err := solutionModel.Detail(sid)
		if err != nil {
			class.Logger.Debug(err)
			http.Error(w, err.Error(), 400)
			return
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

		problemModel := model.ProblemModel{}
		pro, err := problemModel.Detail(sol.Pid)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
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
