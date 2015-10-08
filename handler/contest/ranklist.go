package contest

import (
	"ojapi/config"
	"ojapi/model"

	"restweb"

	"encoding/csv"
	"io"
	"os"
	"sort"
	"strconv"
)

type ContestRanklist struct {
	Contest
} //@Controller

//Download 下载contest排名csv文件
//@URL: /contests/(\d+)/rankfile @method: GET
func (rc *ContestRanklist) Download(Cid string) {
	rc.InitContest(Cid)
	filename := strconv.Itoa(rc.Cid) + ".csv"
	f, err := os.Create(filename)
	if err != nil {
		restweb.Logger.Debug(err)
		return
	}
	defer os.Remove(filename)

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	rankcsv := csv.NewWriter(f)
	rankcsv.Write([]string{"Rank", "Team", "Solved", "Penalty"})

	for rank, user := range rc.ranklist() {
		rankcsv.Write([]string{strconv.Itoa(rank + 1), user.Uid, strconv.Itoa(user.Solved), restweb.ShowGapTime(user.Time)})
	}
	rankcsv.Flush()
	f.Close()

	file, _ := os.Open(filename)
	defer file.Close()

	w := rc.W
	finfo, _ := file.Stat()
	w.Header().Set("ContentType", "application/octet-stream")
	w.Header().Add("Content-disposition", "attachment; filename="+filename)
	w.Header().Add("Content-Length", strconv.Itoa(int(finfo.Size())))

	_, err = io.Copy(w, file)
	if err != nil {
		restweb.Logger.Debug(err)
	}

}

//Home ranklist 列表主页
//@URL: /contests/(\d+)/ranklist @method: GET
func (rc *ContestRanklist) Home(Cid string) {
	rc.InitContest(Cid)
	rc.Output["UserList"] = rc.ranklist()
	rc.Output["IsContestRanklist"] = true
	rc.Output["Cid"] = rc.Cid
	rc.Output["ProblemList"] = rc.ContestDetail.List
	rc.RenderTemplate("view/layout.tpl", "view/contest/ranklist.tpl")
}

//ranklist 实时计算排名
func (rc *ContestRanklist) ranklist() UserSorter {
	qry := make(map[string]string)
	qry["module"] = strconv.Itoa(config.ModuleC)
	qry["mid"] = strconv.Itoa(rc.Cid)
	qry["sort"] = "resort"

	solutionModel := model.SolutionModel{}
	solutionList, err := solutionModel.List(qry)
	if err != nil {
		restweb.Logger.Debug(err)
		return nil
	}

	UserMap := make(map[string]*userRank)
	var pro *probleminfo
	var user *userRank
	for _, v := range solutionList {
		user = UserMap[v.Uid]
		if user == nil {
			user = &userRank{}
			UserMap[v.Uid] = user
			user.ProblemList = make([]*probleminfo, len(rc.Index), len(rc.Index))
		}
		user.Uid = v.Uid
		pid := rc.Index[v.Pid]
		pro = user.ProblemList[pid]
		if pro == nil {
			pro = &probleminfo{}
			user.ProblemList[pid] = pro
		}
		if pro.Judge == config.JudgeAC {
			continue
		}
		pro.Pid = pid
		if v.Judge != config.JudgeAC && v.Judge != config.JudgePD && v.Judge != config.JudgeRJ && v.Judge != config.JudgeNA {
			pro.Count++
			pro.Time += 20 * 60 //罚时20分钟
		} else if v.Judge == config.JudgeAC {
			user.Time += pro.Time
			pro.Time = v.Create - rc.ContestDetail.Start
			pro.Judge = config.JudgeAC
			user.Time += pro.Time
			user.Solved += 1
		}
	}
	UserList := newSorter(UserMap)
	sort.Sort(UserList)
	return UserList
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
