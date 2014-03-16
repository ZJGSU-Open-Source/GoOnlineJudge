package controller

import (
	"GoOnlineJudge/class"
	"html/template"
	"log"
	"net/http"
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

func (this *UserController) Signin(w http.ResponseWriter, r *http.Request) {
	log.Println("User Login")
	this.Init(w, r)

	t := template.New("layout.tpl")
	t, err := t.ParseFiles("view/layout.tpl", "view/user_signin.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "User Sign In"
	this.Data["IsUserSignIn"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
