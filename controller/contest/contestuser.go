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

func (this *ContestUserContorller) Register(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Contest User")
	this.InitContest(w, r)

	if this.Privilege < config.PrivilegeTC {
		if time.Now().Unix() < this.ContestDetail.Start || this.ContestDetail.Status == config.StatusReverse {
			info := "The contest has not started yet"
			if this.ContestDetail.Status == config.StatusReverse {
				info = "No such contest"
			}
			this.Err400(w, r, "Contest Detail "+strconv.Itoa(this.Cid), info)
			return
		} else if this.ContestDetail.Encrypt == config.EncryptPW {
			if this.Uid == "" {
				http.Redirect(w, r, "/user?signin", http.StatusFound)
				return
			} else if this.GetSession(w, r, strconv.Itoa(this.Cid)) != this.ContestDetail.Argument.(string) {
				this.Password(w, r)
				return
			}
		} else if this.ContestDetail.Encrypt == config.EncryptPT {
			if this.Uid == "" {
				http.Redirect(w, r, "/user?signin", http.StatusFound)
			} else {
				userlist := strings.Split(this.ContestDetail.Argument.(string), "\n")
				flag := false
				for _, user := range userlist {
					if user == this.Uid {
						flag = true
						break
					}
				}
				if flag == false {
					this.Err400(w, r, this.ContestDetail.Title,
						"Sorry, the contest is private and you are not granted to participate in the contest.")
					return
				}
			}
		}
	}

	args := this.ParseURL(r.URL.String())
	class.Logger.Debug(args)
	if args["problem"] != "" {
		c := &ProblemController{}
		c.Route(w, r)
	} else if args["status"] != "" {
		c := &StatusController{}
		c.Route(w, r)
	} else if args["ranklist"] != "" {
		c := &RanklistController{}
		c.List(w, r)
	} else {
		class.Logger.Debug("args err")
		http.Error(w, "args err", 500)
		return
	}
}

func (this *ContestUserContorller) Password(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := this.Execute(w, "view/layout.tpl", "view/contest/passwd.tpl")
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
	} else if r.Method == "POST" {
		passwd := r.FormValue("password")
		if passwd == this.ContestDetail.Argument.(string) {
			this.SetSession(w, r, strconv.Itoa(this.Cid), passwd)
			w.WriteHeader(200)
		} else {
			http.Error(w, "incorrect password", 400)
		}
	}
}
