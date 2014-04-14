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

			one.Solve, err = this.GetCount(k, "solve")
			if err != nil {
				http.Error(w, "count error", 500)
				return
			}

			one.Submit, err = this.GetCount(k, "submit")
			if err != nil {
				http.Error(w, "count error", 500)
				return
			}

			list[k] = one
		}
	}
	this.Data["Problem"] = list

	t := template.New("layout.tpl").Funcs(template.FuncMap{"ShowRatio": class.ShowRatio, "ShowStatus": class.ShowStatus})
	t, err := t.ParseFiles("view/layout.tpl", "view/contest/problem_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["IsContestProblem"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ProblemController) Detail(w http.ResponseWriter, r *http.Request) {
	log.Println("Contest Problem Detail")
	this.InitContest(w, r)

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

	t := template.New("layout.tpl").Funcs(template.FuncMap{"ShowRatio": class.ShowRatio, "ShowSpecial": class.ShowSpecial})
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
