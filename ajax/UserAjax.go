package ajax

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

type UserAjax struct {
}

func (this *UserAjax) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("User Login")

	uid := r.FormValue("uid")
	pwd := r.FormValue("pwd")

	w.Header().Set("content-type", "application/json")

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

func (this *UserAjax) Logout(w http.ResponseWriter, r *http.Request) {
	log.Println("User Logout")

	w.Header().Set("content-type", "application/json")

	cookie := http.Cookie{Name: "uid", Value: "", Path: "/"}
	http.SetCookie(w, &cookie)

	out := &Result{}
	out.Uid = ""
	out.Ok = 1

	b, _ := json.Marshal(out)
	w.Write(b)
}
