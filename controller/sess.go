package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/model"

	"restweb"

	// "encoding/json"
	"strconv"
)

type SessController struct {
	class.Controller
} //@Controller

//@URL: /sess @method: GET
func (s *SessController) Get() {
	restweb.Logger.Debug("User Login")

	s.Output["Title"] = "User Sign In"
	s.Output["IsUserSignIn"] = true

	s.RenderTemplate("view/layout.tpl", "view/user_signin.tpl")
}

//@URL: /sess @method: POST
func (s *SessController) Post() {
	restweb.Logger.Debug("User Login")

	var out = struct {
		Handle   string `json:"handle"`
		Password string `json:"password"`
	}{}

	out.Handle = s.Input.Get("user[handle]")
	out.Password = s.Input.Get("user[password]")

	var accessToken string
	body, _ := s.JsonReader(out)
	req, _ := apiClient.NewRequest("POST", "/sess", "", body)
	apiClient.Do(req, &accessToken)

	restweb.Logger.Debug(accessToken)
	var ret *model.User
	req, _ = apiClient.NewRequest("GET", "/profile", accessToken, nil)
	apiClient.Do(req, &ret)

	if ret == nil || ret.Uid == "" {
		s.W.WriteHeader(400)
	} else {
		s.SetSession("Uid", ret.Uid)
		s.SetSession("Privilege", strconv.Itoa(ret.Privilege))
		s.SetSession("AccessToken", accessToken)
		s.W.WriteHeader(201)

		// // remoteAddr := s.R.Header.Get("X-Real-IP")           // if you set niginx as reverse proxy
		// remoteAddr := strings.Split(s.R.RemoteAddr, ":")[0] // otherwise
		// userModel.RecordIP(ret.Uid, remoteAddr, time.Now().Unix())
	}
}

//@URL: /sess @method: Delete
func (s *SessController) Delete() {
	restweb.Logger.Debug("User Logout")

	s.DeleteSession()
	s.W.WriteHeader(200)
}
