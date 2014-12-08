package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/model"

	"restweb"

	"strconv"
	"strings"
	"time"
)

type SessController struct {
	class.Controller
}

func (s SessController) Get() {
	restweb.Logger.Debug("User Login")

	s.Output["Title"] = "User Sign In"
	s.Output["IsUserSignIn"] = true

	s.RenderTemplate("view/layout.tpl", "view/user_signin.tpl")
}

func (s SessController) Post() {
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
		s.Response.WriteHeader(400)
	} else {
		s.SetSession("Uid", uid)
		s.SetSession("Privilege", strconv.Itoa(ret.Privilege))
		s.Response.WriteHeader(200)

		restweb.Logger.Debug(s.Requset.RemoteAddr)
		//remoteAddr := r.Header.Get("X-Real-IP") // if you set niginx as reverse proxy
		//restweb.Logger.Debug(remoteAddr)
		remoteAddr := strings.Split(s.Requset.RemoteAddr, ":")[0] // otherwise
		userModel.RecordIP(uid, remoteAddr, time.Now().Unix())
	}
}

func (s SessController) Delete() {
	restweb.Logger.Debug("User Logout")

	s.DeleteSession()
	s.Response.WriteHeader(200)
}
