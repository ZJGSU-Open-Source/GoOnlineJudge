package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"net/http"
	//"time"
	"sort"
	"strconv"
)

type ranklist struct {
	Uid      string
	costtime int
}
type RanklistController struct {
	Contest
}

func (this *RanklistController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("RankList")
	this.InitContest(w, r)
	response, err := http.Post(config.PostHost+"/solution/list/module/"+strconv.Itoa(config.ModuleC)+"/mid/"+strconv.Itoa(this.Cid), "application/json", nil)
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
		if user.Uid == "" {
			user.ProblemList = make([]*probleminfo, len(this.Index), len(this.Index))
		}
		user.Uid = v.Uid
		pid := this.Index[v.Pid]
		pro = user.ProblemList[pid]
		if pro.Judge == config.JudgeAC {
			continue
		}
		pro.Pid = pid
		if v.Judge != config.JudgeAC && v.Judge != config.JudgePD && v.Judge != config.JudgeRJ {
			pro.count++
			pro.time += 20 * 60 //罚时20分钟
		} else if v.Judge == config.JudgeAC {
			//pro.time += time.Now().Unix() - this.ContestDetail.Start
			pro.Judge = config.JudgeAC
			user.Time += pro.time
		}
	}
	UserList := newSorter(UserMap)
	sort.Sort(UserList)
	this.Data["UserList"] = UserList
	this.Data["Cid"] = this.Cid
	this.Data["ProblemList"] = this.Index
	return
}

type userRank struct {
	Uid         string
	ProblemList []*probleminfo
	Time        int64
}

type probleminfo struct {
	Pid   int
	time  int64
	count int
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
	return u[i].Time < u[j].Time
}

func (u UserSorter) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
