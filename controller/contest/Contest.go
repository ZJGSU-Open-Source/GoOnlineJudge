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

	Start string `json:"start"bson:"start"`
	End   string `json:"end"bson:"end"`

	Status int    `json:"status"bson:"status"`
	Create string `'json:"create"bson:"create"`

	List []int `json:"list"bson:"list"`
}

type Contest struct {
	Cid    int
	Detail *contest
	class.Controller
}

func (this *Contest) InitContest(w http.ResponseWriter, r *http.Request) {
	args := this.ParseURL(r.URL.Path[8:])
	cid, err := strconv.Atoi(args["cid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	this.Cid = cid

	response, err := http.Post(config.PostHost+"/contest/detail/cid/"+strconv.Itoa(cid), "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &this.Detail)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
	}
}
