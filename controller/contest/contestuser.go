package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"

	"net/http"
	"strconv"
	"strings"
	"time"
)

type ContestUserContorller struct {
	Contest
}

var RouterMap = map[string]class.Router{
	"problem":  ProblemController{},
	"status":   StatusController{},
	"ranklist": RanklistController{},
}

func (cuc ContestUserContorller) Route(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest User")
	cuc.InitContest(w, r)

	if cuc.Privilege < config.PrivilegeTC {
		if time.Now().Unix() < cuc.ContestDetail.Start || cuc.ContestDetail.Status == config.StatusReverse {
			info := "The contest has not started yet"
			if cuc.ContestDetail.Status == config.StatusReverse {
				info = "No such contest"
			}
			cuc.Err400(w, r, "Contest Detail "+strconv.Itoa(cuc.Cid), info)
			return
		} else if cuc.ContestDetail.Encrypt == config.EncryptPW {
			if cuc.Uid == "" {
				http.Redirect(w, r, "/user/signin", http.StatusFound)
				return
			} else if cuc.GetSession(w, r, strconv.Itoa(cuc.Cid)) != cuc.ContestDetail.Argument.(string) {
				cuc.Password(w, r)
				return
			}
		} else if cuc.ContestDetail.Encrypt == config.EncryptPT {
			if cuc.Uid == "" {
				http.Redirect(w, r, "/user/signin", http.StatusFound)
				return
			} else {
				userlist := strings.Split(cuc.ContestDetail.Argument.(string), "\n")
				flag := false
				for _, user := range userlist {
					if user == cuc.Uid {
						flag = true
						break
					}
				}
				if flag == false {
					cuc.Err400(w, r, cuc.ContestDetail.Title,
						"Sorry, the contest is private and you are not granted to participate in the contest.")
					return
				}
			}
		}
	}

	action := cuc.GetAction(r.URL.Path, 1)
	class.Logger.Debug(action)
	if v, ok := RouterMap[action]; ok {
		v.Route(w, r)
	} else {
		http.Error(w, "no such page", 404)
	}
}

func (cuc *ContestUserContorller) Password(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cuc.Execute(w, "view/layout.tpl", "view/contest/passwd.tpl")
	} else if r.Method == "POST" {
		passwd := r.FormValue("password")
		if passwd == cuc.ContestDetail.Argument.(string) {
			cuc.SetSession(w, r, strconv.Itoa(cuc.Cid), passwd)
			w.WriteHeader(200)
		} else {
			http.Error(w, "incorrect password", 400)
		}
	}
}
