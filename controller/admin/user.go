package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type privilegeUser struct {
	model.User
	Index int `json:"index"bson:"index"`
}

type UserController struct {
	class.Controller
}

func (this UserController) Route(w http.ResponseWriter, r *http.Request) {
	this.Init(w, r)
	if this.Privilege < config.PrivilegeAD {
		this.Err400(w, r, "Admin", "Privilege Error")
		return
	}
	action := this.GetAction(r.URL.Path, 2)
	defer func() {
		if e := recover(); e != nil {
			http.Error(w, "no such page", 404)
		}
	}()
	rv := class.GetReflectValue(w, r)
	class.CallMethod(&this, strings.Title(action), rv)
}

func (this *UserController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Privilege User List")

	if this.Privilege != config.PrivilegeAD {
		class.Logger.Info(r.RemoteAddr + " " + this.Uid + " try to visit Admin page")
		this.Data["Title"] = "Warning"
		this.Data["Info"] = "You are not admin!"
		this.Execute(w, "view/layout.tpl", "view/400.tpl")
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
	this.Execute(w, "view/admin/layout.tpl", "view/admin/user_list.tpl")
}

func (this *UserController) Pagepassword(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Password Page")

	this.Data["Title"] = "Admin Password"
	this.Data["IsSettings"] = true
	this.Data["IsSettingsPassword"] = true
	this.Data["IsUser"] = true
	this.Data["IsPwd"] = true

	this.Execute(w, "view/admin/layout.tpl", "view/admin/user_password.tpl")
}

func (this *UserController) Password(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Password")

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
	b, _ := json.Marshal(&hint)
	w.Write(b)
}

func (this *UserController) Privilegeset(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("User Privilege")

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
		http.Error(w, "args error", 400)
	}

	ok := 1
	hint := make(map[string]string)

	if uid == "" {
		ok, hint["hint"] = 0, "Handle should not be empty."
	} else if uid == this.Uid {
		ok, hint["hint"] = 0, "You cannot delete yourself!"
	} else {
		userModel := model.UserModel{}
		_, err := userModel.Detail(uid)
		if err == model.NotFoundErr {
			ok, hint["hint"] = 0, "This handle does not exist!"
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
		b, _ := json.Marshal(&hint)

		w.WriteHeader(400)
		w.Write(b)
	}
}

func (this *UserController) Generate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		this.Data["Title"] = "Admin User Generate"
		this.Data["IsUser"] = true
		this.Data["IsGenerate"] = true

		this.Execute(w, "view/admin/layout.tpl", "view/admin/user_generate.tpl")
	} else if r.Method == "POST" {
		prefix := r.FormValue("prefix")
		amount, err := strconv.Atoi(r.FormValue("amount"))
		if err != nil {
			http.Error(w, "args error", 400)
		}
		//TODO:account type
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
			class.Logger.Debug(uid, password)
			one := model.User{}
			one.Uid = uid
			one.Pwd = password
			one.Module, _ = strconv.Atoi(r.FormValue("module"))
			if err := usermodel.Insert(one); err == nil {
				accountlist += uid + " \t" + password + "\n"
				i++
			}
			nxt++
		}
		w.Header().Add("ContentType", "application/octet-stream")
		w.Header().Add("Content-disposition", "attachment; filename=accountlist.txt")
		w.Header().Add("Content-Length", strconv.Itoa(len(accountlist)))
		w.Write([]byte(accountlist))
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
