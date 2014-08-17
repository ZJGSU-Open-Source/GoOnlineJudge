package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"net/http"
	// "strconv"
	"strings"
)

type contest struct {
	Cid      int         `json:"cid"bson:"cid"`
	Title    string      `json:"title"bson:"title"`
	Encrypt  int         `json:"encrypt"bson:"encrypt"`
	Argument interface{} `json:"argument"bson:"argument"`
	Type     string      `json:"type"bson:"type"` //the type of contest,acm contest or normal exercise

	Start int64 `json:"start"bson:"start"`
	End   int64 `json:"end"bson:"end"`

	Status int    `json:"status"bson:"status"`
	Create string `'json:"create"bson:"create"`

	List []int `json:"list"bson:"list"`
}

type ContestController struct {
	class.Controller
	Type string
}

func (this *ContestController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest List")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	//class.Logger.Debug(args)
	Type := args["type"]
	response, err := http.Post(config.PostHost+"/contest/list/type/"+Type, "application", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	one := make(map[string][]*contest)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["Contest"] = one["list"]
	}
	t := template.New("layout.tpl").Funcs(template.FuncMap{
		"ShowStatus":  class.ShowStatus,
		"ShowExpire":  class.ShowExpire,
		"ShowEncrypt": class.ShowEncrypt,
		"LargePU":     class.LargePU})
	t, err = t.ParseFiles("view/layout.tpl", "view/contest_list.tpl")
	if err != nil {
		class.Logger.Debug(err)
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Time"] = this.GetTime()
	this.Data["Type"] = Type
	this.Data["Title"] = strings.Title(Type) + " List"
	this.Data["Is"+strings.Title(Type)] = true
	this.Data["Privilege"] = this.Privilege
	err = t.Execute(w, this.Data)
	if err != nil {
		class.Logger.Debug(err)
		http.Error(w, "tpl error", 500)
		return
	}
}
