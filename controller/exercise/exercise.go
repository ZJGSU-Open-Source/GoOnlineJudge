package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"net/http"
	"strconv"
)

type exercise struct {
	Cid      int         `json:"cid"bson:"cid"`
	Title    string      `json:"title"bson:"title"`
	Encrypt  int         `json:"encrypt"bson:"encrypt"`
	Argument interface{} `json:"argument"bson:"argument"`

	Start string `json:"start"bson:"start"`
	End   string `json:"end"bson:"end"`

	Status int    `json:"status"bson:"status"`
	Create string `'json:"create"bson:"create"`

	List []int `json:"list"bson:"list"`
}

type Exercise struct {
	Cid            int
	ExerciseDetail *exercise
	Index          map[int]int
	class.Controller
}

func (this *Exercise) InitExercise(w http.ResponseWriter, r *http.Request) {
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path[8:])
	cid, err := strconv.Atoi(args["cid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	this.Cid = cid

	response, err := http.Post(config.PostHost+"/exercise/detail/cid/"+strconv.Itoa(cid), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &this.ExerciseDetail)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
	}

	this.Index = make(map[int]int)
	for k, v := range this.ExerciseDetail.List {
		this.Index[v] = k
	}
	this.Data["Cid"] = strconv.Itoa(this.Cid)
	this.Data["Title"] = "Exercise Detail " + strconv.Itoa(this.Cid)
	this.Data["Exercise"] = this.ExerciseDetail.Title
	this.Data["IsExerciseDetail"] = true
}

func (this *Exercise) GetCount(query string) (count int, err error) {
	response, err := http.Post(config.PostHost+"/solution/count/module/"+strconv.Itoa(config.ModuleC)+"/mid/"+strconv.Itoa(this.Cid)+query, "application/json", nil)
	if err != nil {
		return
	}
	defer response.Body.Close()

	one := make(map[string]int)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			return
		}
		count = one["count"]
	}
	return
}
