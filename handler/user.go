package handler

import (
	"ojapi/config"
	"ojapi/middleware"
	"ojapi/model"

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
		Handle   string `json:"handle"`
		Nick     string `json:"nick"`
		Password string `json:"password"`
		Mail     string `json:"mail"`
		School   string `json:"school"`
		Motto    string `json:"motto"`
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

	var (
		user = middleware.ToUser(c)
	)

	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	in := struct {
		Nick      string `json:"nick"`
		Mail      string `json:"mail"`
		School    string `json:"school"`
		Motto     string `json:"motto"`
		ShareCode bool   `json:"share_code"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.WriteHeader(400)
		return
	}

	user.Nick = in.Nick
	user.Mail = in.Mail
	user.School = in.School
	user.Motto = in.Motto
	user.ShareCode = in.ShareCode

	if user.Nick == "" {
		hint := make(map[string]string)
		hint["nick"] = "Nick should not be empty."
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(hint)
	} else {
		userModel := model.UserModel{}
		err := userModel.Update(user.Uid, *user)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	}

}

//@URL: /api/account @method: PUT
func Password(c web.C, w http.ResponseWriter, r *http.Request) {

	var (
		user = middleware.ToUser(c)
	)

	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	in := struct {
		oldPwd string `json:"old_password"`
		newPwd string `json:"new_password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.WriteHeader(400)
		return
	}

	userModel := model.UserModel{}
	ret, err := userModel.Login(user.Uid, in.oldPwd)
	if err != nil || ret.Uid == "" {
		w.WriteHeader(400)
		return
	}

	err = userModel.Password(user.Uid, in.newPwd)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)

}
