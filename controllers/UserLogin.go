package controllers

import (
	"GoOnlineJudge/models"
	"encoding/json"
	"log"
	"net/http"
)

type Result struct {
	Uid string
	Ok  int
}

type UserLoginController struct {
}

func (this *UserLoginController) POST(w http.ResponseWriter, r *http.Request) {
	log.Println("User Login")

	w.Header().Set("content-type", "application/json")

	uid := r.FormValue("uid")
	pwd := r.FormValue("pwd")

	m := &models.UserModel{}
	if m.Login(uid, pwd) {
		log.Println("User Login Successfully")

		cookie := http.Cookie{Name: "uid", Value: uid, Path: "/"}
		http.SetCookie(w, &cookie)

		out := &Result{}
		out.Uid = uid
		out.Ok = 1

		b, _ := json.Marshal(out)
		w.Write(b)
	} else {
		log.Println("User Login Failed")

		out := &Result{}
		out.Uid = uid
		out.Ok = 0

		b, _ := json.Marshal(out)
		w.Write(b)
	}
}
