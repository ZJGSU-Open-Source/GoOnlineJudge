package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"net/http"
)

type AdminUserController struct {
	class.Controller
}

func (this *AdminUserController) Register(w http.ResponseWriter, r *http.Request) {
	this.Init(w, r)

	if this.Privilege <= config.PrivilegePU {
		this.Err400(w, r, "Warning", "You are not admin!")
		return
	} else if r.URL.Path == "/admin/" {
		c := HomeController{}
		c.Home(w, r)
	} else {
		action := this.GetAction(r.URL.Path, 1)
		switch action {
		case "news":
			c := &NewsController{}
			c.Route(w, r)
		case "problem":
			c := &ProblemController{}
			c.Route(w, r)
		case "contest":
			c := &ContestController{}
			c.Route(w, r)
		case "testdata":
			c := &TestdataController{}
			c.Route(w, r)
		case "image":
			c := &ImageController{}
			c.Upload(w, r)
		case "user":
			if this.Privilege < config.PrivilegeAD {
				this.Err400(w, r, "Admin", "Privilege Error")
				return
			}
			c := &UserController{}
			c.Route(w, r)
		default:
			http.Error(w, "no such page", 404)
		}
	}
}
