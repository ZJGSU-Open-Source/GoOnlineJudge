package handler

import (
	"GoOnlineJudge/model"
	"GoOnlineJudge/session"

	"github.com/zenazn/goji/web"

	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

//@URL: /sess @method: POST
func PostSess(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("post sess")

	uid := r.FormValue("handle")
	pwd := r.FormValue("password")

	userModel := model.UserModel{}
	ret, err := userModel.Login(uid, pwd)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if ret.Uid == "" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		session.DeleteSession(r)
		s := session.StartSession()
		s.Set("Uid", uid)
		json.NewEncoder(w).Encode(s.Sid)

		// remoteAddr := r.Header.Get("X-Real-IP") // if you set niginx as reverse proxy
		remoteAddr := strings.Split(s.R.RemoteAddr, ":")[0] // otherwise
		userModel.RecordIP(uid, remoteAddr, time.Now().Unix())
	}
}

//@URL: /sess @method: Delete
func DeleteSess(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("delete sess")

	session.DeleteSession(r)
	w.WriteHeader(http.StatusNoContent)
}
