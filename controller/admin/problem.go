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
	"time"
)

type ProblemController struct {
	class.Controller
}

func (this ProblemController) Route(w http.ResponseWriter, r *http.Request) {
	this.Init(w, r)
	action := this.GetAction(r.URL.Path, 2)
	defer func() {
		if e := recover(); e != nil {
			http.Error(w, "no such page", 404)
		}
	}()
	rv := class.GetReflectValue(w, r)
	class.CallMethod(&this, strings.Title(action), rv)
}

func (this *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Problem Detail")

	pid, err := strconv.Atoi(r.URL.Query().Get("pid"))
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

	if this.Privilege != config.PrivilegeAD {
		this.Err400(w, r, "Warning", "Error Privilege to Add problem")
		return
	}
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
	if r.Method != "POST" {
		this.Err400(w, r, "Error", "Error Method to Insert problem")
		return
	}

	if this.Privilege != config.PrivilegeAD {
		this.Err400(w, r, "Warning", "Error Privilege to Insert problem")
		return
	}

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

	createfile(config.Datapath+strconv.Itoa(pid), "sample.in", in)
	createfile(config.Datapath+strconv.Itoa(pid), "sample.out", out)

	http.Redirect(w, r, "/admin/problem/list", http.StatusFound)
}

func createfile(path, filename string, context string) {

	err := os.Mkdir(path, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		class.Logger.Debug("create dir error")
		return
	}

	file, err := os.Create(path + "/" + filename)
	if err != nil {
		class.Logger.Debug(err)
	}
	defer file.Close()

	var cr rune = 13
	crStr := string(cr)
	context = strings.Replace(context, "\r\n", "\n", -1)
	context = strings.Replace(context, crStr, "\n", -1)
	file.WriteString(context)
}

func (this *ProblemController) Status(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Problem Status")
	if r.Method != "POST" {
		this.Err400(w, r, "Error", "Error Method to Change problem status")
		return
	}

	if this.Privilege != config.PrivilegeAD {
		this.Err400(w, r, "Warning", "Error Privilege to Change problem status")
		return
	}

	pid, err := strconv.Atoi(r.URL.Query().Get("pid"))
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

	http.Redirect(w, r, "/admin/problem/list", http.StatusFound)
}

func (this *ProblemController) Delete(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Problem Delete")
	if r.Method != "POST" {
		this.Err400(w, r, "Error", "Error Method to Delete problem")
		return
	}

	if this.Privilege != config.PrivilegeAD {
		this.Err400(w, r, "Warning", "Error Privilege to Delete problem")
		return
	}

	pid, err := strconv.Atoi(r.URL.Query().Get("pid"))
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

	if this.Privilege != config.PrivilegeAD {
		this.Err400(w, r, "Warning", "Error Privilege to Edit problem")
		return
	}

	pid, err := strconv.Atoi(r.URL.Query().Get("pid"))
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
	if r.Method != "POST" {
		this.Err400(w, r, "Error", "Error Method to Update problem")
		return
	}

	if this.Privilege != config.PrivilegeAD {
		this.Err400(w, r, "Warning", "Error Privilege to Update problem")
		return
	}

	pid, err := strconv.Atoi(r.URL.Query().Get("pid"))
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

	in := r.FormValue("in")
	out := r.FormValue("out")

	one.Description = template.HTML(r.FormValue("description"))
	one.Input = template.HTML(r.FormValue("input"))
	one.Output = template.HTML(r.FormValue("output"))
	one.In = in
	one.Out = out
	one.Source = r.FormValue("source")
	one.Hint = r.FormValue("hint")

	createfile(config.Datapath+strconv.Itoa(pid), "sample.in", in)
	createfile(config.Datapath+strconv.Itoa(pid), "sample.out", out)

	problemModel := model.ProblemModel{}
	err = problemModel.Update(pid, one)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/admin/problem/detail?pid="+strconv.Itoa(pid), http.StatusFound)
}

func (this *ProblemController) Rejudgepage(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Rejudge Page")

	if this.Privilege < config.PrivilegeTC {
		this.Err400(w, r, "Warning", "Error Privilege to Rejudge problem")
		return
	}

	this.Data["Title"] = "Problem Rejudge"
	this.Data["RejudgePrivilege"] = true
	this.Data["IsProblem"] = true
	this.Data["IsRejudge"] = true

	err := this.Execute(w, "view/admin/layout.tpl", "view/admin/problem_rejudge.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ProblemController) Rejudge(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Problem Rejudge")

	if this.Privilege < config.PrivilegeTC {
		this.Err400(w, r, "Warning", "Error Privilege to Rejudge problem")
		return
	}

	args := r.URL.Query()
	types := args.Get("type")
	id, err := strconv.Atoi(args.Get("id"))
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
			hint["info"] = "Problem does not exist!"

			b, _ := json.Marshal(&hint)
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

			time.Sleep(1 * time.Second)
			go func() {
				cmd := exec.Command("./RunServer", "-sid", strconv.Itoa(sid), "-time", strconv.Itoa(pro.Time), "-memory", strconv.Itoa(pro.Memory), "-rejudge", strconv.Itoa(1)) //Run Judge
				err = cmd.Run()
				if err != nil {
					class.Logger.Debug(err)
				}
			}()
		}
	} else if types == "Sid" {
		sid := id

		solutionModel := model.SolutionModel{}
		sol, err := solutionModel.Detail(sid)
		if err != nil {
			class.Logger.Debug(err)

			hint["info"] = "Solution does not exist!"
			b, _ := json.Marshal(&hint)
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
			cmd := exec.Command("./RunServer", "-sid", strconv.Itoa(sid), "-time", strconv.Itoa(pro.Time), "-memory", strconv.Itoa(pro.Memory), "-rejudge", strconv.Itoa(1)) //Run Judge
			err = cmd.Run()
			if err != nil {
				class.Logger.Debug(err)
			}
		}()
	}
}
