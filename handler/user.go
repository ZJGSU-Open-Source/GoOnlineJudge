package handler

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

//@URL: /api/users @method:POST
func (uc *UserController) Register() {

	restweb.Logger.Debug("User Register")

	var one model.User
	userModel := model.UserModel{}

	in := struct {
		Handle   string
		Nick     string
		Password string
		Mail     string
		School   string
		Motto    string
	}{}
	if err := json.NewDecoder(sc.R.Body).Decode(&in); err != nil {
		sc.Error(err.Error(), http.StatusBadRequest)
		return
	}

	one.Privilege = config.PrivilegePU

	err := userModel.Insert(one)
	if err != nil {
		uc.Error(err.Error(), 500)
		return
	}

	uc.W.Header().Add("Location", "/users/"+uid)
	uc.W.WriteHeader(201)
}

//@URL: /api/users/(.+) @method: GET
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
	uc.RenderJson()
}

//@URL: /api/profile @method: GET
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

	uc.RenderJson()
}

//@URL: /api/profile @method: PUT
func (uc *UserController) Update() {

	restweb.Logger.Debug("User Update")

	var one model.User

	in := struct {
		Handle    string
		Nick      string
		Mail      string
		School    string
		Motto     string
		ShareCode string
	}{}
	if err := json.NewDecoder(sc.R.Body).Decode(&in); err != nil {
		sc.Error(err.Error(), http.StatusBadRequest)
		return
	}

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

//@URL: /api/account @method: PUT
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
