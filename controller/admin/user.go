package admin

import (
    "GoOnlineJudge/class"
    "GoOnlineJudge/config"
    "GoOnlineJudge/model"

    "restweb"

    "encoding/json"
    "fmt"
    "math/rand"
    "strconv"
    "time"
)

type privilegeUser struct {
    model.User
    Index int `json:"index"bson:"index"`
}

type AdminUser struct {
    class.Controller
}   //@Controller

//显示具有特殊权限的用户
//@URL: /admin/users/ @method: GET
func (uc *AdminUser) List() {
    restweb.Logger.Debug("Admin Privilege User List")

    if uc.Privilege != config.PrivilegeAD {
        restweb.Logger.Info(uc.Uid + " try to visit Admin page")
        uc.Output["Title"] = "Warning"
        uc.Output["Info"] = "You are not admin!"
        uc.RenderTemplate("view/layout.tpl", "view/400.tpl")
        return
    }

    userModel := model.UserModel{}
    userlist, err := userModel.List(nil)
    if err != nil {
        uc.Error(err.Error(), 500)
        return
    }

    uc.Output["User"] = userlist
    uc.Output["Title"] = "Privilege User List"
    uc.Output["IsUser"] = true
    uc.Output["IsList"] = true
    uc.RenderTemplate("view/admin/layout.tpl", "view/admin/user_list.tpl")
}

//密码设置页面
//@URL: /admin/users/pagepassword @method: GET
func (uc *AdminUser) Pagepassword() {
    restweb.Logger.Debug("Admin Password Page")

    uc.Output["Title"] = "Admin Password"
    uc.Output["IsSettings"] = true
    uc.Output["IsSettingsPassword"] = true
    uc.Output["IsUser"] = true
    uc.Output["IsPwd"] = true

    uc.RenderTemplate("view/admin/layout.tpl", "view/admin/user_password.tpl")
}

//设置用户密码
//@URL: /admin/users/password @method: POST
func (uc *AdminUser) Password() {
    restweb.Logger.Debug("Admin Password")

    ok := 1
    hint := make(map[string]string)
    data := make(map[string]string)

    data["userHandle"] = uc.Input.Get("user[Handle]")
    data["newPassword"] = uc.Input.Get("user[newPassword]")
    data["confirmPassword"] = uc.Input.Get("user[confirmPassword]")

    uid := uc.Input.Get("user[Handle]")

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

        uc.W.WriteHeader(200)
    } else {
        uc.W.WriteHeader(400)
    }
    b, _ := json.Marshal(&hint)
    uc.W.Write(b)
}

// 设置用户权限
//@URL: /admin/privilegeset @method: POST
func (uc *AdminUser) Privilegeset() {
    restweb.Logger.Debug("User Privilege")

    uid := uc.Input.Get("uid")
    privilegeStr := uc.Input.Get("type")

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

        uc.W.WriteHeader(200)
    } else {
        b, _ := json.Marshal(&hint)

        uc.W.WriteHeader(400)
        uc.W.Write(b)
    }
}

//Generate 生成指定数量的用户账号
//@URL: /admin/users/generation @method: GET
func (uc *AdminUser) GeneratePage() {
    uc.Output["Title"] = "Admin User Generate"
    uc.Output["IsUser"] = true
    uc.Output["IsGenerate"] = true
    uc.RenderTemplate("view/admin/layout.tpl", "view/admin/user_generate.tpl")
}

//@URL: /admin/users/generation @method: POST
func (uc *AdminUser) Generate() {
    prefix := uc.Input["prefix"][0]
    module, _ := strconv.Atoi(uc.Input["module"][0])
    module %= 2
    amount, _ := strconv.Atoi(uc.Input["amount"][0])
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
        one := model.User{Uid: uid, Pwd: password, Module: module}
        one.Privilege = config.PrivilegePU
        if err := usermodel.Insert(one); err == nil {
            accountlist += uid + " \t" + password + "\n"
            i++
        }
        nxt++
    }

    uc.W.Header().Set("ContentType", "application/octet-stream")
    uc.W.Header().Add("Content-disposition", "attachment; filename=accountlist.txt")
    uc.W.Header().Add("Content-Length", strconv.Itoa(len(accountlist)))
    uc.W.Write([]byte(accountlist))
}

//RandPassword 生成随机8位密码
func RandPassword() string {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    n := r.Int63()
    return fmt.Sprintf("%08d", n%100000000)
}
