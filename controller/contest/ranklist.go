package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"net/http"
	"sort"
	"strconv"
)

type RanklistController struct {
	Contest
}

func (this *RanklistController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("RankList")
	this.InitContest(w, r)
	response, err := http.Post(config.PostHost+"/solution?list/module?"+strconv.Itoa(config.ModuleC)+"/mid?"+strconv.Itoa(this.Cid)+"/sort?resort/", "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return
	}

	one := make(map[string][]solution)
	err = this.LoadJson(response.Body, &one)
	if err != nil {
		http.Error(w, "load error", 400)
		return
	}

	UserMap := make(map[string]*userRank)
	var pro *probleminfo
	var user *userRank
	for _, v := range one["list"] {
		user = UserMap[v.Uid]
		if user == nil {
			user = &userRank{}
			UserMap[v.Uid] = user
			user.ProblemList = make([]*probleminfo, len(this.Index), len(this.Index))
		}
		user.Uid = v.Uid
		pid := this.Index[v.Pid]
		pro = user.ProblemList[pid]
		if pro == nil {
			pro = &probleminfo{}
			user.ProblemList[pid] = pro
		}
		if pro.Judge == config.JudgeAC {
			continue
		}
		pro.Pid = pid
		if v.Judge != config.JudgeAC && v.Judge != config.JudgePD && v.Judge != config.JudgeRJ {
			pro.Count++
			pro.Time += 20 * 60 //罚时20分钟
		} else if v.Judge == config.JudgeAC {
			pro.Time += v.Create - this.ContestDetail.Start
			pro.Judge = config.JudgeAC
			user.Time += pro.Time
			user.Solved += 1
		}
	}
	UserList := newSorter(UserMap)
	sort.Sort(UserList)

	t := template.New("layout.tpl").Funcs(template.FuncMap{
		"NumAdd": class.NumAdd})
	t, err = t.ParseFiles("view/layout.tpl", "view/contest/ranklist.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
	this.Data["UserList"] = UserList
	this.Data["IsContestRanklist"] = true
	this.Data["Cid"] = this.Cid
	this.Data["ProblemList"] = this.Index
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
	return
}

type userRank struct {
	Uid         string
	ProblemList []*probleminfo
	Time        int64
	Solved      int
}

type probleminfo struct {
	Pid   int
	Time  int64
	Count int
	Judge int
}

type UserSorter []*userRank

func newSorter(m map[string]*userRank) UserSorter {
	us := make([]*userRank, 0, len(m))
	for _, v := range m {
		us = append(us, v)
	}
	return us
}

func (u UserSorter) Len() int {
	return len(u)
}

func (u UserSorter) Less(i, j int) bool {
	if u[i].Solved > u[j].Solved {
		return true
	} else if u[i].Solved == u[j].Solved {
		if u[i].Time < u[j].Time {
			return true
		} else if u[i].Time >= u[j].Time {
			return false
		}
	}
	return false
}

func (u UserSorter) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
