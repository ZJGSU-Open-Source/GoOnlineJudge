package controller

import (
    "GoOnlineJudge/class"
    "GoOnlineJudge/model"

    "restweb"

    "strconv"
    "time"
)

type SessController struct {
    class.Controller
}   //@Controller

//@URL: /api/sess @method: POST
func (s *SessController) Post() {
    restweb.Logger.Debug("User Login")

    uid := s.Input.Get("user[handle]")
    pwd := s.Input.Get("user[password]")

    userModel := model.UserModel{}
    ret, err := userModel.Login(uid, pwd)
    if err != nil {
        restweb.Logger.Debug(err)
        s.Error(err.Error(), 500)
        return
    }

    if ret.Uid == "" {
        s.W.WriteHeader(400)
    } else {
        s.SetSession("Uid", uid)
        s.SetSession("Privilege", strconv.Itoa(ret.Privilege))
        s.W.WriteHeader(201)

        remoteAddr := s.R.Header.Get("X-Real-IP") // if you set niginx as reverse proxy
        // remoteAddr := strings.Split(s.R.RemoteAddr, ":")[0] // otherwise
        userModel.RecordIP(uid, remoteAddr, time.Now().Unix())
    }
}

//@URL: /api/sess @method: Delete
func (s *SessController) Delete() {
    restweb.Logger.Debug("User Logout")

    s.DeleteSession()
    s.W.WriteHeader(200)
}
