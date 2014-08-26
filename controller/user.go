package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserController struct {
	class.Controller
}

func (this *UserController) Signin(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Login")
	this.Init(w, r)

	this.Data["Title"] = "User Sign In"
	this.Data["IsUserSignIn"] = true
	err := this.Execute(w, "view/layout.tpl", "view/user_signin.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *UserController) Login(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Login")
	this.Init(w, r)

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
		this.SetSession(w, r, "Uid", uid)
		this.SetSession(w, r, "Privilege", strconv.Itoa(ret.Privilege))
		w.WriteHeader(200)
	}
	return
}

func (this *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Sign Up")
	this.Init(w, r)

	this.Data["Title"] = "User Sign Up"
	this.Data["IsUserSignUp"] = true
	err := this.Execute(w, "view/layout.tpl", "view/user_signup.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *UserController) Register(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Register")
	this.Init(w, r)

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
		qry := make(map[string]string)
		qry["uid"] = uid
		ret, err := userModel.List(qry)
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else if len(ret) > 0 {
			ok, hint["uid"] = 0, "This handle is currently in use."
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
	if ok == 1 {
		one.Uid = uid
		one.Nick = nick
		one.Pwd = pwd
		one["privilege"] = config.PrivilegePU
		//one.Privilege = config.PrivilegeAD

		err := userModel.Insert(one)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		this.SetSession(w, r, "Uid", uid)
		this.SetSession(w, r, "Privilege", "1")
		w.WriteHeader(200)
	} else {
		b, err := json.Marshal(&hint)
		if err != nil {
			http.Error(w, "json error", 500)
			return
		}

		w.WriteHeader(400)
		w.Write(b)
	}
}

func (this *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Logout")
	this.Init(w, r)

	this.DeleteSession(w, r)
	w.WriteHeader(200)
}

func (this *UserController) Detail(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Detail")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	uid := args["uid"]
	userModel := model.UserModel{}
	one, err := userModel.Detail(uid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	this.Data["Detail"] = one

	solutionModle := model.SolutionModel{}
	solvedList, err := solutionModle.Achieve(uid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	this.Data["List"] = solvedList
	//class.Logger.Debug(solvedList)
	this.Data["Title"] = "User Detail"
	if uid != "" && uid == this.Uid {
		this.Data["IsSettings"] = true
		this.Data["IsSettingsDetail"] = true
	}

	err = this.Execute(w, "view/layout.tpl", "view/user_detail.tpl")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (this *UserController) Settings(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Settings")
	this.Init(w, r)

	if this.Privilege == config.PrivilegeNA {
		this.Data["Title"] = "Warning"
		this.Data["Info"] = "You must login!"
		err := this.Execute(w, "view/layout.tpl", "view/400.tpl")
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
		return
	}

	userModel := model.UserModel{}

	one, err := userModel.Detail(this.Uid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	this.Data["Detail"] = one

	solutionModel := model.SolutionModel{}
	solvedList, err := solutionModel.Achieve(this.Uid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	this.Data["List"] = solvedList

	this.Data["Title"] = "User Settings"
	this.Data["IsSettings"] = true
	this.Data["IsSettingsDetail"] = true

	err = this.Execute(w, "view/layout.tpl", "view/user_detail.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *UserController) Edit(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Edit")
	this.Init(w, r)

	if this.Privilege == config.PrivilegeNA {
		this.Data["Title"] = "Warning"
		this.Data["Info"] = "You must login!"
		err := this.Execute(w, "view/layout.tpl", "view/400.tpl")
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
		return
	}

	uid := this.Uid
	userModel := model.UserModel{}
	one, err := userModel.Detail(uid)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	this.Data["Detail"] = one

	this.Data["Title"] = "User Edit"
	this.Data["IsSettings"] = true
	this.Data["IsSettingsEdit"] = true

	err = this.Execute(w, "view/layout.tpl", "view/user_edit.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *UserController) Update(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Update")
	this.Init(w, r)

	ok := 1
	hint := make(map[string]string)
	hint["uid"] = this.Uid

	var one model.User
	one.Nick = r.FormValue("user[nick]")
	one.Mail = r.FormValue("user[mail]")
	one.School = r.FormValue("user[school]")
	one.Motto = r.FormValue("user[motto]")

	if one.Nick == "" {
		ok, hint["nick"] = 0, "Nick should not be empty."
	}

	if ok == 1 {
		userModel := model.UserModel{}
		err := userModel.Update(this.Uid, one)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
	}

	b, err := json.Marshal(&hint)
	if err != nil {
		http.Error(w, "json error", 400)
		return
	}
	w.Write(b)
}

func (this *UserController) Pagepassword(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Password Page")
	this.Init(w, r)

	if this.Privilege == config.PrivilegeNA {
		this.Data["Title"] = "Warning"
		this.Data["Info"] = "You must login!"
		err := this.Execute(w, "view/layout.tpl", "view/400.tpl")
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
		return
	}

	this.Data["Title"] = "User Password"
	this.Data["IsSettings"] = true
	this.Data["IsSettingsPassword"] = true

	err := this.Execute(w, "view/layout.tpl", "view/user_password.tpl")
	if err != nil {
		http.Error(w, "tpl error", 400)
		return
	}
}

func (this *UserController) Password(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Password")
	this.Init(w, r)

	ok := 1
	hint := make(map[string]string)
	hint["uid"] = this.Uid

	data := make(map[string]string)
	data["oldPassword"] = r.FormValue("user[oldPassword]")
	data["newPassword"] = r.FormValue("user[newPassword]")
	data["confirmPassword"] = r.FormValue("user[confirmPassword]")

	uid := this.Uid
	pwd := data["oldPassword"]

	userModel := model.UserModel{}
	ret, err := userModel.Login(uid, pwd)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if ret.Uid == "" {
		ok, hint["oldPassword"] = 0, "Old Password is Incorrect."
	}
	if len(data["newPassword"]) < 6 {
		ok, hint["newPassword"] = 0, "Password should contain at least six characters."
	}
	if data["newPassword"] != data["confirmPassword"] {
		ok, hint["confirmPassword"] = 0, "Confirmation mismatched."
	}

	if ok == 1 {
		pwd = data["newPassword"]
		err := userModel.Password(uid, pwd)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
	}
	b, err := json.Marshal(&hint)
	if err != nil {
		http.Error(w, "json error", 400)
		return
	}

	w.Write(b)
}
