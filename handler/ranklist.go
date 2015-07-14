package handler

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"encoding/json"
	"net/http"
)

// 排名
type rank struct {
	model.User
	Index int `json:"index"bson:"index"`
}

// 排名控件
type RanklistController struct {
	class.Controller
} //@Controller

//@URL: /api/ranklist @method: GET
func (rc *RanklistController) Index() {

	restweb.Logger.Debug("Ranklist")

	// in := struct {
	//     Offset int
	//     Limit  int
	// }{}
	// if err := json.NewDecoder(rc.R.Body).Decode(&in); err != nil {
	//     rc.Error(err.Error(), http.StatusBadRequest)
	//     return
	// }
	userModel := model.UserModel{}

	qry := make(map[string]string)

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
	rc.Output["User"] = list
	rc.RenderJson()
}
