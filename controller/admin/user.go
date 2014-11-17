package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

type privilegeUser struct {
	model.User
	Index int `json:"index"bson:"index"`
}

type UserController struct {
	class.Controller
}

//显示具有特殊权限的用户，url:/admin/user/list
func (uc *UserController) List() {
	restweb.Logger.Debug("Admin Privilege User List")

	if uc.Privilege != config.PrivilegeAD {
		restweb.Logger.Info(uc.Uid + " try to visit Admin page")
		uc.Data["Title"] = "Warning"
		uc.Data["Info"] = "You are not admin!"
		uc.RenderTemplate("view/layout.tpl", "view/400.tpl")
		return
	}

	userModel := model.UserModel{}
	userlist, err := userModel.List(nil)
	if err != nil {
		uc.Error(err.Error(), 500)
		return
	}

	uc.Data["User"] = userlist
	uc.Data["Title"] = "Privilege User List"
	uc.Data["IsUser"] = true
	uc.Data["IsList"] = true
	uc.RenderTemplate("view/admin/layout.tpl", "view/admin/user_list.tpl")
}

//密码设置页面,url: /admin/user/pagepassword
func (uc *UserController) Pagepassword() {
	restweb.Logger.Debug("Admin Password Page")

	uc.Data["Title"] = "Admin Password"
	uc.Data["IsSettings"] = true
	uc.Data["IsSettingsPassword"] = true
	uc.Data["IsUser"] = true
	uc.Data["IsPwd"] = true

	uc.RenderTemplate("view/admin/layout.tpl", "view/admin/user_password.tpl")
}

//设置用户密码，url:/admin/user/password, method: POST
func (uc *UserController) Password() {
	restweb.Logger.Debug("Admin Password")

	ok := 1
	hint := make(map[string]string)
	r := uc.Requset
	data := make(map[string]string)
	data["userHandle"] = r.FormValue("user[Handle]")
	data["newPassword"] = r.FormValue("user[newPassword]")
	data["confirmPassword"] = r.FormValue("user[confirmPassword]")

	uid := r.FormValue("user[Handle]")

	if uid == "" {
		ok, hint["uid"] = 0, "Handle should not be empty"
	} else {
		userModel := model.UserModel{}
		_, err := userModel.Detail(uid)
		if err == model.NotFoundErr {
			ok, hint["uid"] = 0, "uc handle does not exist!"
		} else if err != nil {
			uc.Error(err.Error(), 400)
			return
		}

	}

	if len(data["newPassword"]) < 6 {
		ok, hint["newPassword"] = 0, "Password should contain at least six characters."
	}
	if data["newPassword"] != data["confirmPassword"] {
		ok, hint["confirmPassword"] = 0, "Confirmation mismatched."
	}

	if ok == 1 {
		pwd := data["newPassword"]
		userModel := model.UserModel{}
		err := userModel.Password(uid, pwd)
		if err != nil {
			uc.Error(err.Error(), 400)
			return
		}

		uc.Response.WriteHeader(200)
	} else {
		uc.Response.WriteHeader(400)
	}
	b, _ := json.Marshal(&hint)
	uc.Response.Write(b)
}

// 设置用户权限
func (uc *UserController) Privilegeset() {
	restweb.Logger.Debug("User Privilege")

	r := uc.Requset
	args := r.URL.Query()
	uid := args.Get("uid")
	privilegeStr := args.Get("type")

	privilege := config.PrivilegeNA
	switch privilegeStr {
	case "Admin":
		privilege = config.PrivilegeAD
	case "TC":
		privilege = config.PrivilegeTC
	case "PU":
		privilege = config.PrivilegePU
	default:
		uc.Error("args error", 400)
	}

	ok := 1
	hint := make(map[string]string)

	if uid == "" {
		ok, hint["hint"] = 0, "Handle should not be empty."
	} else if uid == uc.Uid {
		ok, hint["hint"] = 0, "You cannot delete yourself!"
	} else {
		userModel := model.UserModel{}
		_, err := userModel.Detail(uid)
		if err == model.NotFoundErr {
			ok, hint["hint"] = 0, "uc handle does not exist!"
		} else if err != nil {
			uc.Error(err.Error(), 400)
			return
		}
	}

	if ok == 1 {
		userModel := model.UserModel{}
		err := userModel.Privilege(uid, privilege)
		if err != nil {
			uc.Error(err.Error(), 400)
			return
		}

		uc.Response.WriteHeader(200)
	} else {
		b, _ := json.Marshal(&hint)

		uc.Response.WriteHeader(400)
		uc.Response.Write(b)
	}
}

//Generate 生成指定数量的用户账号，/admin/user/generate
func (uc *UserController) Generate() {
	r := uc.Requset
	if r.Method == "GET" {
		uc.Data["Title"] = "Admin User Generate"
		uc.Data["IsUser"] = true
		uc.Data["IsGenerate"] = true
		uc.RenderTemplate("view/admin/layout.tpl", "view/admin/user_generate.tpl")

	} else if r.Method == "POST" {
		prefix := r.FormValue("prefix")
		module, _ := strconv.Atoi(r.FormValue("module"))
		module %= 2
		amount, _ := strconv.Atoi(r.FormValue("amount"))
		if amount > 100 {
			amount = 100
		}

		count := 0
		tmp := amount
		for tmp > 0 {
			tmp /= 10
			count++
		}

		format := "%0" + strconv.Itoa(count) + "d"
		usermodel := &model.UserModel{}
		accountlist := "Uid \tPassword\n"

		for i, nxt := 0, 1; i < amount; {
			uid := prefix + fmt.Sprintf(format, nxt)
			password := RandPassword()
			restweb.Logger.Debug(uid, password)
			one := model.User{}
			one.Uid = uid
			one.Pwd = password
			one.Module = module
			one.Module, _ = strconv.Atoi(r.FormValue("module"))
			if err := usermodel.Insert(one); err == nil {
				accountlist += uid + " \t" + password + "\n"
				i++
			}
			nxt++
		}

		uc.Response.Header().Add("ContentType", "application/octet-stream")
		uc.Response.Header().Add("Content-disposition", "attachment; filename=accountlist.txt")
		uc.Response.Header().Add("Content-Length", strconv.Itoa(len(accountlist)))
		uc.Response.Write([]byte(accountlist))
	}
}

//RandPassword 生成随机8位密码
func RandPassword() string {
	b := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)[:8]
}
