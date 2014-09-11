package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"encoding/json"
	"net/http"
)

type privilegeUser struct {
	model.User
	Index int `json:"index"bson:"index"`
}

type UserController struct {
	class.Controller
}

func (this *UserController) Route(w http.ResponseWriter, r *http.Request) {
	this.Init(w, r)
	action := this.GetAction(r.URL.Path, 2)
	switch action {
	case "list":
		this.List(w, r)
	case "Privilegeset":
		this.Privilegeset(w, r)
	case "Pagepassword":
		this.Pagepassword(w, r)
	case "password":
		this.Password(w, r)
	default:
		http.Error(w, "no such page", 404)
	}
}

func (this *UserController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Privilege User List")

	if this.Privilege != config.PrivilegeAD {
		class.Logger.Info(r.RemoteAddr + " " + this.Uid + " try to visit Admin page")
		this.Data["Title"] = "Warning"
		this.Data["Info"] = "You are not admin!"
		err := this.Execute(w, "view/layout.tpl", "view/400.tpl")
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
		return
	}
	userModel := model.UserModel{}
	userlist, err := userModel.List(nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	this.Data["User"] = userlist
	this.Data["Title"] = "Privilege User List"
	this.Data["IsUser"] = true
	this.Data["IsList"] = true

	err = this.Execute(w, "view/admin/layout.tpl", "view/admin/user_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *UserController) Pagepassword(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Password Page")
	this.Init(w, r)

	this.Data["Title"] = "Admin Password"
	this.Data["IsSettings"] = true
	this.Data["IsSettingsPassword"] = true
	this.Data["IsUser"] = true
	this.Data["IsPwd"] = true

	err := this.Execute(w, "view/admin/layout.tpl", "view/admin/admin_password.tpl")
	if err != nil {
		http.Error(w, "tpl error", 400)
		return
	}
}

func (this *UserController) Password(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Password")
	this.Init(w, r)

	ok := 1
	hint := make(map[string]string)

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
			ok, hint["uid"] = 0, "This handle does not exist!"
		} else if err != nil {
			http.Error(w, err.Error(), 400)
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

func (this *UserController) Privilegeset(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Privilege")
	this.Init(w, r)

	args := this.ParseURL(r.URL.String())
	uid := args["uid"]
	privilegeStr := args["type"]

	privilege := config.PrivilegeNA
	switch privilegeStr {
	case "Admin":
		privilege = config.PrivilegeAD
	case "TC":
		privilege = config.PrivilegeTC
	case "PU":
		privilege = config.PrivilegePU
	default:
		http.Error(w, "args error", 400)
	}

	ok := 1
	hint := make(map[string]string)

	if uid == "" {
		ok, hint["uid"] = 0, "Handle should not be empty."
	} else if uid == this.Uid {
		ok, hint["uid"] = 0, "You cannot delete yourself"
	} else {
		userModel := model.UserModel{}
		_, err := userModel.Detail(uid)
		if err == model.NotFoundErr {
			ok, hint["uid"] = 0, "This handle does not exist!"
		} else if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	if ok == 1 {
		userModel := model.UserModel{}
		err := userModel.Privilege(uid, privilege)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

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
