package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"

	"net/http"
)

type AdminUserController struct {
	class.Controller
}

var RouterMap = map[string]class.Router{
	"news":     NewsController{},
	"problem":  ProblemController{},
	"contest":  ContestController{},
	"testdata": TestdataController{},
	"image":    ImageController{},
	"user":     UserController{},
}

func (auc AdminUserController) Route(w http.ResponseWriter, r *http.Request) {
	auc.Init(w, r)

	if auc.Privilege <= config.PrivilegePU {
		auc.Err400(w, r, "Warning", "You are not admin!")
		return
	} else if r.URL.Path == "/admin/" {
		c := HomeController{}
		c.Home(w, r)
	} else {
		action := auc.GetAction(r.URL.Path, 1)
		if v, ok := RouterMap[action]; ok {
			v.Route(w, r)
		} else {
			http.Error(w, "no such page", 404)
		}
	}
}
