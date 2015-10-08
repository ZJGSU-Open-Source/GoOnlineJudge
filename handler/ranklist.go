package handler

import (
	// "ojapi/class"
	"ojapi/config"
	"ojapi/model"

	"github.com/zenazn/goji/web"

	"encoding/json"
	"net/http"
)

// 排名
type rank struct {
	model.User
	Index int `json:"index"bson:"index"`
}

//@URL: /api/ranklist @method: GET
func Ranklist(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	qry := make(map[string]string)
	if v := r.FormValue("offset"); len(v) > 0 {
		qry["offset"] = v
	}
	if v := r.FormValue("limit"); len(v) > 0 {
		qry["limit"] = v
	}

	userModel := model.UserModel{}
	userList, err := userModel.List(qry)
	if err != nil {

	}

	list := make([]rank, len(userList), len(userList))
	count := 1
	for i, one := range userList {
		list[i].User = *one
		if one.Status == config.StatusAvailable {
			list[count-1].Index = count //+ in.Offset*in.Limit
			count += 1
		}
	}

	json.NewEncoder(w).Encode(list)
}
