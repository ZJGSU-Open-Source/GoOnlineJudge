package contest

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
	Contest
}

func (this *ProblemController) List(w http.ResponseWriter, r *http.Request) {
	log.Println("Contest Problem List")
	this.InitContest(w, r)

	if (this.GetTime() < this.ContestDetail.Start || this.ContestDetail.Status == config.StatusReverse) && this.Privilege <= config.PrivilegePU {
		t := template.New("layout.tpl")
		t, err := t.ParseFiles("view/layout.tpl", "view/400.tpl")
		if err != nil {
			log.Println(err)
			http.Error(w, "tpl error", 500)
			return
		}

		this.Data["Info"] = "The contest has not started yet"
		if this.ContestDetail.Status == config.StatusReverse {
			this.Data["Info"] = "No such contest"
		}
		err = t.Execute(w, this.Data)
		if err != nil {
			log.Println(err)
			http.Error(w, "tpl error", 500)
			return
		}
		return
	}

	list := make([]problem, len(this.ContestDetail.List))
	for k, v := range this.ContestDetail.List {
		response, err := http.Post(config.PostHost+"/problem/detail/pid/"+strconv.Itoa(v), "application/json", nil)
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
			one.Pid = k
			query := "/pid/" + strconv.Itoa(v) + "/action/accept"
			one.Solve, err = this.GetCount(query)
			if err != nil {
				http.Error(w, "count error", 500)
				return
			}
			query = "/pid/" + strconv.Itoa(v) + "/action/submit"
			one.Submit, err = this.GetCount(query)
			if err != nil {
				http.Error(w, "count error", 500)
				return
			}

			list[k] = one
		}
	}
	this.Data["Problem"] = list
	t := template.New("layout.tpl").Funcs(template.FuncMap{
		"ShowRatio": class.ShowRatio})
	t, err := t.ParseFiles("view/layout.tpl", "view/contest/problem_list.tpl")
	if err != nil {
		log.Println(err)
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["IsContestProblem"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		log.Println(err)
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
	log.Println("Contest Problem Detail")
	this.InitContest(w, r)

	if (this.ContestDetail.Status == config.StatusReverse || this.GetTime() < this.ContestDetail.Start) && this.Privilege <= config.PrivilegePU {
		t := template.New("layout.tpl")
		t, err := t.ParseFiles("view/layout.tpl", "view/400.tpl")
		if err != nil {
			log.Println(err)
			http.Error(w, "tpl error", 500)
			return
		}

		this.Data["Info"] = "The contest has not started yet"
		if this.ContestDetail.Status == config.StatusReverse {
			this.Data["Info"] = "No such contest"
		}
		err = t.Execute(w, this.Data)
		if err != nil {
			log.Println(err)
			http.Error(w, "tpl error", 500)
			return
		}
		return
	}

	args := this.ParseURL(r.URL.Path[8:])
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/problem/detail/pid/"+strconv.Itoa(this.ContestDetail.List[pid]), "application/json", nil)
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
	this.Data["Pid"] = pid
	this.Data["Status"] = this.ContestDetail.Status

	t := template.New("layout.tpl").Funcs(template.FuncMap{
		"ShowRatio":   class.ShowRatio,
		"ShowStatus":  class.ShowStatus,
		"ShowSpecial": class.ShowSpecial})
	t, err = t.ParseFiles("view/layout.tpl", "view/contest/problem_detail.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

/////////Todo submit ,need to updata------

func (this *ProblemController) Submit(w http.ResponseWriter, r *http.Request) {
	log.Println("Contest Problem Submit")
	this.InitContest(w, r)

	args := this.ParseURL(r.URL.Path[8:])

	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	pid = this.ContestDetail.List[pid]
	uid := this.Uid
	if uid == "" {
		w.WriteHeader(401)
		return
	}

	one := make(map[string]interface{})
	one["pid"] = pid
	one["uid"] = uid
	one["mid"] = this.ContestDetail.Cid
	one["module"] = config.ModuleC
	/////TODO. Judge

	response, err := http.Post(config.PostHost+"/problem/detail/pid/"+strconv.Itoa(pid), "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	var pro problem
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &pro)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
	}

	action := "submit"
	one["judge"], one["time"], one["memory"] = config.JudgeAC, 1000, 888
	//sljudge.SJudge(1, pro.Time, pro.Memory, pid, r.FormValue("code")) //solution judge 最好做成外部程序
	if one["judge"] == config.JudgeAC { //Judge whether the solution is accepted
		action = "solve"
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

	query := "/pid/" + strconv.Itoa(pid) + "/uid/" + this.Uid + "/action/solve"
	cnt, err := this.GetCount(query)
	if err != nil {
		http.Error(w, "query erro", 500)
	}
	///end count
	///count if the problem has been solved
	if cnt >= 1 && action == "solve" {
		action = "submit"
	}
	response, err = http.Post(config.PostHost+"/user/record/uid/"+uid+"/action/"+action, "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	response, err = http.Post(config.PostHost+"/problem/record/pid/"+strconv.Itoa(pid)+"/action/"+action, "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	/////end Judge

	one["status"] = config.StatusAvailable
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
