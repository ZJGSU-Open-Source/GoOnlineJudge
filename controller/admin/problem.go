package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	//"encoding/json"
	"html/template"
	"net/http"
	"os"
	//"os/exec"
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
		http.Error(w, "The value 'Time' is neither too short nor too large", 400)
		return
	}
	one.Time = time
	memory, err := strconv.Atoi(r.FormValue("memory"))
	if err != nil {
		http.Error(w, "The value 'Memory' is neither too short nor too large", 400)
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
		http.Error(w, "The value 'Time' is neither too short nor too large", 500)
		return
	}
	one.Time = time
	memory, err := strconv.Atoi(r.FormValue("memory"))
	if err != nil {
		http.Error(w, "The value 'memory' is neither too short nor too large", 500)
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

/*
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

	hint := make(map[string]string)

	if types == "Pid" {
		pid := id
		proModel := model.ProblemModel{}
		pro, err := proModel.Detail(pid)
		if err != nil {
			class.Logger.Debug(err)
			hint["uid"] = "Problem does not exist!"

			b, err := json.Marshal(&hint)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(400)
			w.Write(b)

			return
		}
		qry := make(map[string]string)
		qry["pid"] = strconv.Itoa(pro.Pid)

		solutionModel := model.SolutionModel{}
		list, err := solutionModel.List(qry)

		for i := range list {
			sid := list[i].Sid

			go func() {
				cmd := exec.Command("./RunServer", "-sid", strconv.Itoa(sid), "-time", strconv.Itoa(pro.Time), "-memory", strconv.Itoa(pro.Memory)) //Run Judge
				err = cmd.Run()
				if err != nil {
					class.Logger.Debug(err)
					return
				}
			}()
		}
	} else if types == "Sid" {
		sid := id

		solutionModel := model.SolutionModel{}
		sol, err := solutionModel.Detail(sid)
		if err != nil {
			class.Logger.Debug(err)
			hint["uid"] = "Solution does not exist!"

			b, err := json.Marshal(&hint)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(400)
			w.Write(b)

			return
		}

		problemModel := model.ProblemModel{}
		pro, err := problemModel.Detail(sol.Pid)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		go func() {
			cmd := exec.Command("./RunServer", "-sid", strconv.Itoa(sid), "-time", strconv.Itoa(pro.Time), "-memory", strconv.Itoa(pro.Memory)) //Run Judge
			err = cmd.Run()
			if err != nil {
				class.Logger.Debug(err)
				return
			}
		}()
	}
}
*/
