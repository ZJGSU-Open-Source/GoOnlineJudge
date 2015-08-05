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
} //@Controller

//@URL: /users/new @method: GET
func (uc *UserController) Signup() {
	restweb.Logger.Debug("User Sign Up")

	uc.Output["Title"] = "User Sign Up"
	uc.Output["IsUserSignUp"] = true
	uc.RenderTemplate("view/layout.tpl", "view/user_signup.tpl")

}

//@URL: /users @method:POST
func (uc *UserController) Register() {
	restweb.Logger.Debug("User Register")

	var one model.User
	userModel := model.UserModel{}

	uid := uc.Input.Get("user[handle]")
	nick := uc.Input.Get("user[nick]")
	pwd := uc.Input.Get("user[password]")
	pwdConfirm := uc.Input.Get("user[confirmPassword]")
	one.Mail = uc.Input.Get("user[mail]")
	one.School = uc.Input.Get("user[school]")
	one.Motto = uc.Input.Get("user[motto]")

	valid := restweb.Validation{}
	valid.MinSize(uid, 4, "uid")
	valid.Match(uid, "\\w+", "uid")

	if !valid.HasError {
		_, err := userModel.Detail(uid)
		if err != nil && err != model.NotFoundErr {
			http.Error(uc.W, err.Error(), 500)
			return
		} else if err == nil {
			valid.AppendError("uid", "Handle is currently in use.")
		}
	}

	valid.Required(nick, "nick")
	valid.MinSize(pwd, 6, "pwd")
	valid.Equal(pwd, pwdConfirm, "pwdConfirm")
	valid.Mail(one.Mail, "mail")

	if !valid.HasError {
		one.Uid = uid
		one.Nick = nick
		one.Pwd = pwd
		one.Privilege = config.PrivilegePU

		err := userModel.Insert(one)
		if err != nil {
			uc.Error(err.Error(), 500)
			return
		}

		uc.W.Header().Add("Location", "/users/"+uid)
		uc.W.WriteHeader(201)
	} else {
		hint := valid.RenderErrMap()
		b, _ := json.Marshal(&hint)
		uc.W.WriteHeader(400)
		uc.W.Write(b)
	}
}

//@URL: /users/(.+) @method: GET
func (uc *UserController) Detail(uid string) {
	restweb.Logger.Debug("User Detail", uid)

	userModel := model.UserModel{}
	one, err := userModel.Detail(uid)
	if err != nil {
		uc.Error(err.Error(), 400)
		return
	}
	uc.Output["Detail"] = one

	solutionModle := model.SolutionModel{}
	solvedList, err := solutionModle.Achieve(uid)
	if err != nil {
		uc.Error(err.Error(), 400)
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
	}

	achieveList := sort.IntSlice(solvedList)
	achieveList.Sort()
	uc.Output["List"] = achieveList
	uc.Output["IpList"] = ips
	uc.Output["Title"] = "User Detail"
	if uid != "" && uid == uc.Uid {
		uc.Output["IsSettings"] = true
		uc.Output["IsSettingsDetail"] = true
	}

	uc.RenderTemplate("view/layout.tpl", "view/user_detail.tpl")
}

//@URL: /settings @method: GET
func (uc *UserController) Settings() {
	restweb.Logger.Debug("User Settings xx", uc.Uid)

	userModel := model.UserModel{}
	one, err := userModel.Detail(uc.Uid)
	if err != nil {
		uc.Error(err.Error(), 400)
		return
	}
	uc.Output["Detail"] = one

	solutionModel := model.SolutionModel{}
	solvedList, err := solutionModel.Achieve(uc.Uid)
	if err != nil {
		uc.Error(err.Error(), 400)
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
	uc.Output["List"] = achieveList
	uc.Output["IpList"] = ips
	uc.Output["Title"] = "User Settings"
	uc.Output["IsSettings"] = true
	uc.Output["IsSettingsDetail"] = true

	uc.RenderTemplate("view/layout.tpl", "view/user_detail.tpl")
}

//@URL: /profile @method: GET
func (uc *UserController) Edit() {
	restweb.Logger.Debug("User Edit")

	uid := uc.Uid
	userModel := model.UserModel{}
	one, err := userModel.Detail(uid)
	if err != nil {
		uc.Error(err.Error(), 400)
		return
	}
	uc.Output["Detail"] = one

	uc.Output["Title"] = "User Edit"
	uc.Output["IsSettings"] = true
	uc.Output["IsSettingsEdit"] = true

	uc.RenderTemplate("view/layout.tpl", "view/user_edit.tpl")
}

//@URL: /profile @method: POST
func (uc *UserController) Update() {
	restweb.Logger.Debug("User Update")

	var one model.User
	one.Nick = uc.Input.Get("user[nick]")
	one.Mail = uc.Input.Get("user[mail]")
	one.School = uc.Input.Get("user[school]")
	one.Motto = uc.Input.Get("user[motto]")
	one.ShareCode, _ = strconv.ParseBool(uc.Input.Get("user[share_code]"))
	restweb.Logger.Debug(uc.Input.Get("user[share_code]"))
	restweb.Logger.Debug(one.ShareCode)

	if one.Nick == "" {
		hint := make(map[string]string)
		hint["nick"] = "Nick should not be empty."
		uc.W.WriteHeader(400)
		b, _ := json.Marshal(&hint)
		uc.W.Write(b)
	} else {
		userModel := model.UserModel{}
		err := userModel.Update(uc.Uid, one)
		if err != nil {
			http.Error(uc.W, err.Error(), 500)
			return
		}
		uc.W.WriteHeader(200)
	}
}

//@URL: /account @method: GET
func (uc *UserController) Pagepassword() {
	restweb.Logger.Debug("User Password Page")

	uc.Output["Title"] = "User Password"
	uc.Output["IsSettings"] = true
	uc.Output["IsSettingsPassword"] = true

	uc.RenderTemplate("view/layout.tpl", "view/user_password.tpl")
}

//@URL: /account @method: POST
func (uc *UserController) Password() {
	restweb.Logger.Debug("User Password")

	valid := restweb.Validation{}

	uid := uc.Uid
	// valid.AppendError("uid", uid)

	oldPwd := uc.Input.Get("user[oldPassword]")
	newPwd := uc.Input.Get("user[newPassword]")
	confirmPwd := uc.Input.Get("user[confirmPassword]")

	userModel := model.UserModel{}
	ret, err := userModel.Login(uid, oldPwd)
	if err != nil {
		uc.Error(err.Error(), 500)
		return
	}

	if ret.Uid == "" {
		valid.AppendError("oldPassword", "Old Password is Incorrect.")
	}
	valid.MinSize(newPwd, 6, "newPassword")
	valid.Equal(newPwd, confirmPwd, "confirmPassword")

	if !valid.HasError {
		err := userModel.Password(uid, newPwd)
		if err != nil {
			uc.Error(err.Error(), 400)
			return
		}

		uc.W.WriteHeader(200)
	} else {
		uc.W.WriteHeader(400)
	}
	hint := valid.RenderErrMap()
	b, _ := json.Marshal(&hint)
	uc.W.Write(b)
}
