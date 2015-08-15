package handler

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/middleware"
	"GoOnlineJudge/model"
	"github.com/zenazn/goji/web"

	"encoding/json"
	"net/http"
	"sort"
)

//@URL: /api/users @method:POST
func PostUser(c web.C, w http.ResponseWriter, r *http.Request) {

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
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.WriteHeader(400)
		return
	}

	one.Privilege = config.PrivilegePU

	err := userModel.Insert(one)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Location", "/users/"+one.Uid)
	w.WriteHeader(201)
}

//@URL: /api/users/:user @method: GET
func GetUser(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		uid = c.URLParams["uid"]
	)

	userModel := model.UserModel{}
	one, err := userModel.Detail(uid)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	output := make(map[string]interface{})
	output["Detail"] = one

	solutionModle := model.SolutionModel{}
	solvedList, err := solutionModle.Achieve(uid)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	type IPs struct {
		Time int64  `json:"time"`
		IP   string `json:"ip"`
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
	output["List"] = achieveList
	output["IpList"] = ips

	json.NewEncoder(w).Encode(output)
}

//@URL: /api/profile @method: GET
func GetProfile(c web.C, w http.ResponseWriter, r *http.Request) {

	var (
		user = middleware.ToUser(c)
	)

	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user.Pwd = ""

	json.NewEncoder(w).Encode(user)
}

//@URL: /api/profile @method: PUT
func PutUser(c web.C, w http.ResponseWriter, r *http.Request) {

	var one model.User

	var (
		user = middleware.ToUser(c)
	)

	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	in := struct {
		Handle    string
		Nick      string
		Mail      string
		School    string
		Motto     string
		ShareCode string
	}{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.WriteHeader(400)
		return
	}

	if one.Nick == "" {
		hint := make(map[string]string)
		hint["nick"] = "Nick should not be empty."
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(hint)
	} else {
		userModel := model.UserModel{}
		err := userModel.Update(user.Uid, one)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	}

}

// //@URL: /api/account @method: PUT
// func Password(c web.C, w http.ResponseWriter, r *http.Request) {

//     uid := ""
//     // valid.AppendError("uid", uid)

//     oldPwd := uc.Input.Get("user[oldPassword]")
//     newPwd := uc.Input.Get("user[newPassword]")
//     confirmPwd := uc.Input.Get("user[confirmPassword]")

//     userModel := model.UserModel{}
//     ret, err := userModel.Login(uid, oldPwd)
//     if err != nil {
//         uc.Error(err.Error(), 500)
//         return
//     }

//     if ret.Uid == "" {
//         valid.AppendError("oldPassword", "Old Password is Incorrect.")
//     }

//     if !valid.HasError {
//         err := userModel.Password(uid, newPwd)
//         if err != nil {
//             uc.Error(err.Error(), 400)
//             return
//         }

//         uc.W.WriteHeader(200)
//     } else {
//         uc.W.WriteHeader(400)
//     }
//     hint := valid.RenderErrMap()
//     b, _ := json.Marshal(&hint)
//     uc.W.Write(b)
// }
