package handler

import (
    "ojapi/model"
    "ojapi/session"

    "github.com/zenazn/goji/web"

    "encoding/json"
    "log"
    "net/http"
    "strings"
    "time"
)

//@URL: /sess @method: POST
func PostSess(c web.C, w http.ResponseWriter, r *http.Request) {
    log.Println("post sess")

    in := struct {
        Handle   string `json:"handle"`
        Password string `json:"password"`
    }{}

    err := json.NewDecoder(r.Body).Decode(&in)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    userModel := model.UserModel{}
    ret, err := userModel.Login(in.Handle, in.Password)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    if ret.Uid == "" {
        w.WriteHeader(http.StatusNotFound)
    } else {
        session.DeleteSession(r)
        s := session.StartSession()
        s.Set("Uid", ret.Uid)
        json.NewEncoder(w).Encode(s.Sid)

        // remoteAddr := r.Header.Get("X-Real-IP") // if you set niginx as reverse proxy
        remoteAddr := strings.Split(r.RemoteAddr, ":")[0] // otherwise
        userModel.RecordIP(ret.Uid, remoteAddr, time.Now().Unix())
    }
}

//@URL: /sess @method: Delete
func DeleteSess(c web.C, w http.ResponseWriter, r *http.Request) {
    log.Println("delete sess")

    session.DeleteSession(r)
    w.WriteHeader(http.StatusNoContent)
}
