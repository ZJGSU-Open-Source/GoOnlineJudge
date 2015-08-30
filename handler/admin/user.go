package admin

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/middleware"
	"GoOnlineJudge/model"

	"github.com/zenazn/goji/web"

	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type privilegeUser struct {
	model.User
	Index int `json:"index"bson:"index"`
}

//显示具有特殊权限的用户
//@URL: /admin/users/ @method: GET
func ListUsers(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		user = middleware.ToUser(c)
	)
	if user.Privilege != config.PrivilegeAD {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	userModel := model.UserModel{}
	userlist, err := userModel.List(nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(userlist)
}

// //设置用户密码
// //@URL: /admin/users/password @method: POST
// func Password(c web.C, w http.ResponseWriter, r *http.Request) {

//     ok := 1
//     hint := make(map[string]string)
//     data := make(map[string]string)

//     // data["userHandle"] = r.FormValue("user[Handle]")
//     // data["newPassword"] = uc.Input.Get("user[newPassword]")
//     // data["confirmPassword"] = uc.Input.Get("user[confirmPassword]")

//     uid := uc.Input.Get("user[Handle]")

//     if uid == "" {
//         ok, hint["uid"] = 0, "Handle should not be empty"
//     } else {
//         userModel := model.UserModel{}
//         _, err := userModel.Detail(uid)
//         if err == model.NotFoundErr {
//             w.WriteHeader(http.StatusNotFound)
//             return
//         } else if err != nil {
//             w.WriteHeader(http.StatusInternalServerError)
//             return
//         }

//     }

//     if len(data["newPassword"]) < 6 {
//         ok, hint["newPassword"] = 0, "Password should contain at least six characters."
//     }
//     if data["newPassword"] != data["confirmPassword"] {
//         ok, hint["confirmPassword"] = 0, "Confirmation mismatched."
//     }

//     if ok == 1 {
//         pwd := data["newPassword"]
//         userModel := model.UserModel{}
//         err := userModel.Password(uid, pwd)
//         if err != nil {
//             w.WriteHeader(http.StatusInternalServerError)
//             return
//         }

//         w.WriteHeader(200)
//     } else {
//         w.WriteHeader(400)
//     }
//     b, _ := json.Marshal(&hint)
//     uc.W.Write(b)
// }

// 设置用户权限
//@URL: /admin/privilegeset @method: POST
func Privilegeset(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		user = middleware.ToUser(c)
	)

	uid := r.Form.Get("uid")
	privilegeStr := r.Form.Get("type")

	privilege := config.PrivilegeNA
	switch privilegeStr {
	case "Admin":
		privilege = config.PrivilegeAD
	case "TC":
		privilege = config.PrivilegeTC
	case "PU":
		privilege = config.PrivilegePU
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

	ok := 1
	hint := make(map[string]string)

	if uid == "" {
		ok, hint["hint"] = 0, "Handle should not be empty."
	} else if uid == user.Uid {
		ok, hint["hint"] = 0, "You cannot delete yourself!"
	} else {
		userModel := model.UserModel{}
		_, err := userModel.Detail(uid)
		if err == model.NotFoundErr {
			w.WriteHeader(http.StatusNotFound)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if ok == 1 {
		userModel := model.UserModel{}
		err := userModel.Privilege(uid, privilege)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
	} else {
		b, _ := json.Marshal(&hint)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(b)
	}
}

//@URL: /admin/users/generation @method: POST
func Generate(c web.C, w http.ResponseWriter, r *http.Request) {
	prefix := r.Form.Get("prefix")
	module, _ := strconv.Atoi(r.Form.Get("module"))
	module %= 2
	amount, _ := strconv.Atoi(r.Form.Get("amount"))
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

		one := model.User{Uid: uid, Pwd: password, Module: module}
		one.Privilege = config.PrivilegePU
		if err := usermodel.Insert(one); err == nil {
			accountlist += uid + " \t" + password + "\n"
			i++
		}
		nxt++
	}

	w.Header().Set("ContentType", "application/octet-stream")
	w.Header().Add("Content-disposition", "attachment; filename=accountlist.txt")
	w.Header().Add("Content-Length", strconv.Itoa(len(accountlist)))
	w.Write([]byte(accountlist))
}

//RandPassword 生成随机8位密码
func RandPassword() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Int63()
	return fmt.Sprintf("%08d", n%100000000)
}
