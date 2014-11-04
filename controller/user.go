package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	class.Controller
}

func (uc UserController) Route(w http.ResponseWriter, r *http.Request) {
	uc.Init(w, r)
	action := uc.GetAction(r.URL.Path, 1)
	defer func() {
		if e := recover(); e != nil {
			http.Error(w, "no such page", 404)
		}
	}()
	rv := class.GetReflectValue(w, r)
	class.CallMethod(&uc, strings.Title(action), rv)
}

func (uc *UserController) Signin(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Login")

	uc.Data["Title"] = "User Sign In"
	uc.Data["IsUserSignIn"] = true

	uc.Execute(w, "view/layout.tpl", "view/user_signin.tpl")
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Login")

	uid := r.FormValue("user[handle]")
	pwd := r.FormValue("user[password]")

	userModel := model.UserModel{}
	ret, err := userModel.Login(uid, pwd)
	if err != nil {
		class.Logger.Debug(err)
		http.Error(w, err.Error(), 500)
		return
	}

	if ret.Uid == "" {
		w.WriteHeader(400)
	} else {
		uc.SetSession(w, r, "Uid", uid)
		uc.SetSession(w, r, "Privilege", strconv.Itoa(ret.Privilege))
		w.WriteHeader(200)

		class.Logger.Debug(r.RemoteAddr)
		//remoteAddr := r.Header.Get("X-Real-IP") // if you set niginx as reverse proxy
		//class.Logger.Debug(remoteAddr)
		remoteAddr := strings.Split(r.RemoteAddr, ":")[0] // otherwise
		userModel.RecordIP(uid, remoteAddr, time.Now().Unix())
	}
}

func (uc *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Sign Up")

	uc.Data["Title"] = "User Sign Up"
	uc.Data["IsUserSignUp"] = true
	uc.Execute(w, "view/layout.tpl", "view/user_signup.tpl")

}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Register")

	var one model.User
	userModel := model.UserModel{}

	uid := r.FormValue("user[handle]")
	nick := r.FormValue("user[nick]")
	pwd := r.FormValue("user[password]")
	pwdConfirm := r.FormValue("user[confirmPassword]")
	one.Mail = r.FormValue("user[mail]")
	one.School = r.FormValue("user[school]")
	one.Motto = r.FormValue("user[motto]")

	ok := 1
	hint := make(map[string]string)

	if uid == "" {
		ok, hint["uid"] = 0, "Handle should not be empty."
	} else {
		_, err := userModel.Detail(uid)
		if err != nil && err != model.NotFoundErr {
			http.Error(w, err.Error(), 500)
			return
		} else if err == nil {
			ok, hint["uid"] = 0, "uc handle is currently in use."
		}
	}

	if nick == "" {
		ok, hint["nick"] = 0, "Nick should not be empty."
	}
	if len(pwd) < 6 {
		ok, hint["pwd"] = 0, "Password should contain at least six characters."
	}
	if pwd != pwdConfirm {
		ok, hint["pwdConfirm"] = 0, "Confirmation mismatched."
	}
	if one.Mail != "" && class.CheckMail(one.Mail) == false {
		ok, hint["mail"] = 0, "Wrong mail."
	}
	if ok == 1 {
		one.Uid = uid
		one.Nick = nick
		one.Pwd = pwd
		one.Privilege = config.PrivilegePU

		err := userModel.Insert(one)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		uc.SetSession(w, r, "Uid", uid)
		uc.SetSession(w, r, "Privilege", strconv.Itoa(config.PrivilegePU))
		w.WriteHeader(200)
	} else {
		b, _ := json.Marshal(&hint)
		w.WriteHeader(400)
		w.Write(b)
	}
}

func (uc *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Logout")

	uc.DeleteSession(w, r)
	w.WriteHeader(200)
}

func (uc *UserController) Detail(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Detail")

	args := r.URL.Query()
	uid := args.Get("uid")
	userModel := model.UserModel{}
	one, err := userModel.Detail(uid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	uc.Data["Detail"] = one

	solutionModle := model.SolutionModel{}
	solvedList, err := solutionModle.Achieve(uid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	type IPs struct {
		Time int64
		IP   string
	}
	var ips []IPs
	ipo := IPs{}

	for i, lenth := 0, len(one.IPRecord); i < lenth; i++ {
		ipo.Time = one.TimeRecord[i]
		ipo.IP = one.IPRecord[i]
		ips = append(ips, ipo)
		class.Logger.Debug(ips[i].IP)
	}

	achieveList := sort.IntSlice(solvedList)
	achieveList.Sort()
	uc.Data["List"] = achieveList
	uc.Data["IpList"] = ips
	uc.Data["Title"] = "User Detail"
	if uid != "" && uid == uc.Uid {
		uc.Data["IsSettings"] = true
		uc.Data["IsSettingsDetail"] = true
	}

	uc.Execute(w, "view/layout.tpl", "view/user_detail.tpl")
}

func (uc *UserController) Settings(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Settings")

	if uc.Uid == "" {
		http.Redirect(w, r, "/user/signin", http.StatusFound)
	}

	userModel := model.UserModel{}
	one, err := userModel.Detail(uc.Uid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	uc.Data["Detail"] = one

	solutionModel := model.SolutionModel{}
	solvedList, err := solutionModel.Achieve(uc.Uid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	type IPs struct {
		Time int64
		IP   string
	}
	var ips []IPs
	ipo := IPs{}

	for i, lenth := 0, len(one.IPRecord); i < lenth; i++ {
		ipo.Time = one.TimeRecord[i]
		ipo.IP = one.IPRecord[i]
		ips = append(ips, ipo)
		class.Logger.Debug(ips[i].IP)
	}

	achieveList := sort.IntSlice(solvedList)
	achieveList.Sort()
	uc.Data["List"] = achieveList
	uc.Data["IpList"] = ips
	uc.Data["Title"] = "User Settings"
	uc.Data["IsSettings"] = true
	uc.Data["IsSettingsDetail"] = true

	uc.Execute(w, "view/layout.tpl", "view/user_detail.tpl")
}

func (uc *UserController) Edit(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Edit")

	if uc.Uid == "" {
		http.Redirect(w, r, "/user/signin", http.StatusFound)
		return
	}

	uid := uc.Uid
	userModel := model.UserModel{}
	one, err := userModel.Detail(uid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	uc.Data["Detail"] = one

	uc.Data["Title"] = "User Edit"
	uc.Data["IsSettings"] = true
	uc.Data["IsSettingsEdit"] = true

	uc.Execute(w, "view/layout.tpl", "view/user_edit.tpl")
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Update")

	var one model.User
	one.Nick = r.FormValue("user[nick]")
	one.Mail = r.FormValue("user[mail]")
	one.School = r.FormValue("user[school]")
	one.Motto = r.FormValue("user[motto]")

	if one.Nick == "" {
		hint := make(map[string]string)
		hint["nick"] = "Nick should not be empty."
		w.WriteHeader(400)
		b, _ := json.Marshal(&hint)
		w.Write(b)
	} else {
		userModel := model.UserModel{}
		err := userModel.Update(uc.Uid, one)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(200)
	}
}

func (uc *UserController) Pagepassword(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Password Page")

	if uc.Uid == "" {
		http.Redirect(w, r, "/user/signin", http.StatusFound)
		return
	}

	uc.Data["Title"] = "User Password"
	uc.Data["IsSettings"] = true
	uc.Data["IsSettingsPassword"] = true

	uc.Execute(w, "view/layout.tpl", "view/user_password.tpl")
}

func (uc *UserController) Password(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Password")

	ok := 1
	uid := uc.Uid
	hint := make(map[string]string)
	hint["uid"] = uid

	oldPwd := r.FormValue("user[oldPassword]")
	newPwd := r.FormValue("user[newPassword]")
	confirmPwd := r.FormValue("user[confirmPassword]")

	userModel := model.UserModel{}
	ret, err := userModel.Login(uid, oldPwd)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if ret.Uid == "" {
		ok, hint["oldPassword"] = 0, "Old Password is Incorrect."
	}
	if len(newPwd) < 6 {
		ok, hint["newPassword"] = 0, "Password should contain at least six characters."
	}
	if newPwd != confirmPwd {
		ok, hint["confirmPassword"] = 0, "Confirmation mismatched."
	}

	if ok == 1 {
		err := userModel.Password(uid, newPwd)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
	}
	b, _ := json.Marshal(&hint)
	w.Write(b)
}
