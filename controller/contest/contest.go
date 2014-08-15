package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"net/http"
	"strconv"
)

type contest struct {
	Cid      int         `json:"cid"bson:"cid"`
	Title    string      `json:"title"bson:"title"`
	Encrypt  int         `json:"encrypt"bson:"encrypt"`
	Argument interface{} `json:"argument"bson:"argument"`
	Type     string      `json:"type"bson:"type"` //the type of contest,acm contest or normal exercise

	Start string `json:"start"bson:"start"`
	End   string `json:"end"bson:"end"`

	Status int    `json:"status"bson:"status"`
	Create string `'json:"create"bson:"create"`

	List []int `json:"list"bson:"list"`
}

type Contest struct {
	Cid           int
	ContestDetail *contest
	Index         map[int]int
	class.Controller
}

func (this *Contest) InitContest(w http.ResponseWriter, r *http.Request) {
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	cid, err := strconv.Atoi(args["cid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	this.Cid = cid

	response, err := http.Post(config.PostHost+"/contest/detail/cid/"+strconv.Itoa(cid), "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &this.ContestDetail)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
	}

	this.Index = make(map[int]int)
	for k, v := range this.ContestDetail.List {
		this.Index[v] = k
	}
	this.Data["Cid"] = strconv.Itoa(this.Cid)
	this.Data["Title"] = "Contest Detail " + strconv.Itoa(this.Cid)
	this.Data["Contest"] = this.ContestDetail.Title
	this.Data["IsContestDetail"] = true
}

func (this *Contest) GetCount(query string) (count int, err error) {
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
