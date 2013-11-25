package controllers

import (
	"GoOnlineJudge/models"
	"log"
	"net/http"
)

type UserLoginController struct {
}

func (this *UserLoginController) POST(w http.ResponseWriter, r *http.Request) {
	log.Println("User Login")

	uid := r.FormValue("uid")
	pwd := r.FormValue("pwd")

	m := &models.UserModel{}
	if m.SignIn(uid, pwd) {
		cookie := http.Cookie{Name: "uid", Value: uid, Path: "/"}
		http.SetCookie(w, &cookie)

		log.Println("User Login Successfully")
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
