package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"encoding/json"
	"net/http"
	"sort"
	"strconv"
)

type UserController struct {
	class.Controller
}

func (uc *UserController) Signup() {
	restweb.Logger.Debug("User Sign Up")

	uc.Data["Title"] = "User Sign Up"
	uc.Data["IsUserSignUp"] = true
	uc.RenderTemplate("view/layout.tpl", "view/user_signup.tpl")

}

func (uc *UserController) Register() {
	restweb.Logger.Debug("User Register")

	var one model.User
	userModel := model.UserModel{}

	uid := uc.Requset.FormValue("user[handle]")
	nick := uc.Requset.FormValue("user[nick]")
	pwd := uc.Requset.FormValue("user[password]")
	pwdConfirm := uc.Requset.FormValue("user[confirmPassword]")
	one.Mail = uc.Requset.FormValue("user[mail]")
	one.School = uc.Requset.FormValue("user[school]")
	one.Motto = uc.Requset.FormValue("user[motto]")

	ok := 1
	hint := make(map[string]string)

	if uid == "" {
		ok, hint["uid"] = 0, "Handle should not be empty."
	} else {
		_, err := userModel.Detail(uid)
		if err != nil && err != model.NotFoundErr {
			http.Error(uc.Response, err.Error(), 500)
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
	if one.Mail != "" {
		ok, hint["mail"] = 0, "Wrong mail."
	}
	if ok == 1 {
		one.Uid = uid
		one.Nick = nick
		one.Pwd = pwd
		one.Privilege = config.PrivilegePU

		err := userModel.Insert(one)
		if err != nil {
			http.Error(uc.Response, err.Error(), 500)
			return
		}

		uc.SetSession("Uid", uid)
		uc.SetSession("Privilege", strconv.Itoa(config.PrivilegePU))
		uc.Response.WriteHeader(200)
	} else {
		b, _ := json.Marshal(&hint)
		uc.Response.WriteHeader(400)
		uc.Response.Write(b)
	}
}

func (uc *UserController) Detail() {
	restweb.Logger.Debug("User Detail")

	uid := uc.GetAction(uc.Requset.URL.Path, 1)
	userModel := model.UserModel{}
	one, err := userModel.Detail(uid)
	if err != nil {
		http.Error(uc.Response, err.Error(), 400)
		return
	}
	uc.Data["Detail"] = one

	solutionModle := model.SolutionModel{}
	solvedList, err := solutionModle.Achieve(uid)
	if err != nil {
		http.Error(uc.Response, err.Error(), 400)
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
		restweb.Logger.Debug(ips[i].IP)
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

	uc.RenderTemplate("view/layout.tpl", "view/user_detail.tpl")
}

func (uc *UserController) Settings() {
	restweb.Logger.Debug("User Settings")

	if uc.Uid == "" {
		uc.Redirect("/user/signin", http.StatusFound)
	}

	userModel := model.UserModel{}
	one, err := userModel.Detail(uc.Uid)
	if err != nil {
		http.Error(uc.Response, err.Error(), 400)
		return
	}
	uc.Data["Detail"] = one

	solutionModel := model.SolutionModel{}
	solvedList, err := solutionModel.Achieve(uc.Uid)
	if err != nil {
		http.Error(uc.Response, err.Error(), 400)
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
		restweb.Logger.Debug(ips[i].IP)
	}

	achieveList := sort.IntSlice(solvedList)
	achieveList.Sort()
	uc.Data["List"] = achieveList
	uc.Data["IpList"] = ips
	uc.Data["Title"] = "User Settings"
	uc.Data["IsSettings"] = true
	uc.Data["IsSettingsDetail"] = true

	uc.RenderTemplate("view/layout.tpl", "view/user_detail.tpl")
}

func (uc *UserController) Edit() {
	restweb.Logger.Debug("User Edit")

	if uc.Uid == "" {
		uc.Redirect("/user/signin", http.StatusFound)
		return
	}

	uid := uc.Uid
	userModel := model.UserModel{}
	one, err := userModel.Detail(uid)
	if err != nil {
		http.Error(uc.Response, err.Error(), 400)
		return
	}
	uc.Data["Detail"] = one

	uc.Data["Title"] = "User Edit"
	uc.Data["IsSettings"] = true
	uc.Data["IsSettingsEdit"] = true

	uc.RenderTemplate("view/layout.tpl", "view/user_edit.tpl")
}

func (uc *UserController) Update() {
	restweb.Logger.Debug("User Update")

	var one model.User
	one.Nick = uc.Requset.FormValue("user[nick]")
	one.Mail = uc.Requset.FormValue("user[mail]")
	one.School = uc.Requset.FormValue("user[school]")
	one.Motto = uc.Requset.FormValue("user[motto]")

	if one.Nick == "" {
		hint := make(map[string]string)
		hint["nick"] = "Nick should not be empty."
		uc.Response.WriteHeader(400)
		b, _ := json.Marshal(&hint)
		uc.Response.Write(b)
	} else {
		userModel := model.UserModel{}
		err := userModel.Update(uc.Uid, one)
		if err != nil {
			http.Error(uc.Response, err.Error(), 500)
			return
		}
		uc.Response.WriteHeader(200)
	}
}

func (uc *UserController) Pagepassword() {
	restweb.Logger.Debug("User Password Page")

	if uc.Uid == "" {
		uc.Redirect("/user/signin", http.StatusFound)
		return
	}

	uc.Data["Title"] = "User Password"
	uc.Data["IsSettings"] = true
	uc.Data["IsSettingsPassword"] = true

	uc.RenderTemplate("view/layout.tpl", "view/user_password.tpl")
}

func (uc *UserController) Password() {
	restweb.Logger.Debug("User Password")

	ok := 1
	uid := uc.Uid
	hint := make(map[string]string)
	hint["uid"] = uid

	oldPwd := uc.Requset.FormValue("user[oldPassword]")
	newPwd := uc.Requset.FormValue("user[newPassword]")
	confirmPwd := uc.Requset.FormValue("user[confirmPassword]")

	userModel := model.UserModel{}
	ret, err := userModel.Login(uid, oldPwd)
	if err != nil {
		http.Error(uc.Response, err.Error(), 500)
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
			http.Error(uc.Response, err.Error(), 400)
			return
		}

		uc.Response.WriteHeader(200)
	} else {
		uc.Response.WriteHeader(400)
	}
	b, _ := json.Marshal(&hint)
	uc.Response.Write(b)
}
