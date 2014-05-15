package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type user struct {
	Uid string `json:"uid"bson:"uid"`
	Pwd string `json:"pwd"bson:"pwd"`

	Nick   string `json:"nick"bson:"nick"`
	Mail   string `json:"mail"bson:"mail"`
	School string `json:"school"bson:"school"`
	Motto  string `json:"motto"bson:"motto"`

	Privilege int `json:"privilege"bson:"privilege"`

	Solve  int `json:"solve"bson:"solve"`
	Submit int `json:"submit"bson:"submit"`

	Status int    `json:"status"bson:"status"`
	Create string `json:"create"bson:'create'`
}

type UserController struct {
	class.Controller
}

func (this *UserController) List() {
	log.Println("Admin Problem List")
	this.Init(w, r)

	response, err := http.Post(config.PostHost+"/user/list", "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	one := make(map[string][]user)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		var len = len(one["list"])
		for i := 0; i < len; i++ {
			if one["list"][i].Privilege > config.PrivilegePU {
				one["list"][i].Index = count
				count += 1
			}
		}
		this.Data["User"] = one["list"]
	}

	funcMap := map[string]interface{}{
		"NumEqual": class.NumEqual,
	}

	t := template.New("layout.tpl").Funcs(funcMap)
	t, err = t.ParseFiles("view/admin/layout.tpl", "view/admin/userlist.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Privilege User List"
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
func (this *UserController) Privilege(w http.ResponseWriter, r *http.Request) {

}
