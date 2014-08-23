package model

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/model/class"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
	"time"
)

type Solution struct {
	Sid int `json:"sid"bson:"sid"`

	Pid      int    `json:"pid"bson:"pid"`
	Uid      string `json:"uid"bson:"uid"`
	Judge    int    `json:"judge"bson:"judge"`
	Time     int    `json:"time"bson:"time"`
	Memory   int    `json:"memory"bson:"memory"`
	Length   int    `json:"length"bson:"length"`
	Language int    `json:"language"bson:"language"`

	Module int `json:"module"bson:"module"`
	Mid    int `json:"mid"bson:"mid"`

	Code string `json:"code"bson:"code"`

	Status int   `json:"status"bson:"status"`
	Create int64 `json:"create"bson:"create"`
}

var sDetailSelector = bson.M{"_id": 0}
var sAchieveSelector = bson.M{"_id": 0, "pid": 1}
var sListSelector = bson.M{"_id": 0, "sid": 1, "pid": 1, "uid": 1, "judge": 1, "time": 1, "memory": 1, "length": 1, "language": 1, "create": 1, "status": 1}

type SolutionModel struct {
	class.Model
}

// POST /Solution?detail/sid?<sid>
func (this *SolutionModel) Detail(sid int) (*Solution, error) {
	log.Println("Server SolutionModel Detail")

	err := this.OpenDB()
	if err != nil {
		return nil, class.DBErr
	}
	defer this.CloseDB()

	var one Solution
	err = this.DB.C("Solution").Find(bson.M{"sid": sid}).Select(sDetailSelector).One(&one)
	if err == mgo.ErrNotFound {
		return nil, class.NotFoundErr
	} else if err != nil {
		return nil, class.OpErr
	}

	return &one, nil
}

// POST /Solution?delete/sid?<sid>
func (this *SolutionModel) Delete(sid int) error {
	log.Println("Server SolutionModel Delete")

	err := this.OpenDB()
	if err != nil {
		return class.DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("Solution").Remove(bson.M{"sid": sid})
	if err == mgo.ErrNotFound {
		return class.NotFoundErr
	} else if err != nil {
		return class.OpErr
	}

	return nil
}

// POST /Solution?insert
func (this *SolutionModel) Insert(one Solution) (int, error) {
	log.Println("Server SolutionModel Insert")

	err := this.OpenDB()
	if err != nil {
		return 0, class.DBErr
	}
	defer this.CloseDB()

	one.Create = time.Now().Unix()
	one.Sid, err = this.GetID("Solution")
	if err != nil {
		return 0, class.IDErr
	}

	err = this.DB.C("Solution").Insert(&one)
	if err != nil {
		return 0, class.OpErr
	}

	// b, err := json.Marshal(map[string]interface{}{
	// 	"sid":    one.Sid,
	// 	"status": one.Status,
	// })

	return one.Sid, nil
}

// POST /Solution?update/sid?<sid>
func (this *SolutionModel) Update(sid int, ori Solution) error {
	log.Println("Server SolutionModel Update")

	alt := make(map[string]interface{})
	alt["judge"] = ori.Judge
	alt["time"] = ori.Time
	alt["memory"] = ori.Memory

	err := this.OpenDB()
	if err != nil {
		return class.DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("Solution").Update(bson.M{"sid": sid}, bson.M{"$set": alt})
	if err == mgo.ErrNotFound {
		return class.NotFoundErr
	} else if err != nil {
		return class.OpErr
	}

	return nil
}

// POST /Solution?status/sid?<sid>/action?<0/1/2>
func (this *SolutionModel) Status(sid, status int) error {
	log.Println("Server SolutionModel Status")

	err := this.OpenDB()
	if err != nil {
		return class.DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("Solution").Update(bson.M{"sid": sid}, bson.M{"$inc": bson.M{"status": status}})
	if err == mgo.ErrNotFound {
		return class.NotFoundErr
	} else if err != nil {
		return class.OpErr
	}

	return nil
}

// POST /Solution?count/pid?<pid>/uid?<uid>/module?<module>/mid?<mid>/action?<accept/solve/submit>
func (this *SolutionModel) Count(args map[string]string) (int, error) {
	log.Println("Server SolutionModel Count")

	query, err := this.CheckQuery(args)
	if err != nil {
		return 0, class.ArgsErr
	}

	err = this.OpenDB()
	if err != nil {
		return 0, class.DBErr
	}
	defer this.CloseDB()

	var count int
	c := this.DB.C("Solution")
	switch v := args["action"]; v {
	case "submit":
		count, err = c.Find(query).Count()
		if err != nil {
			return 0, class.QueryErr
		}
	case "accept":
		query["judge"] = config.JudgeAC
		count, err = c.Find(query).Count()
		if err != nil {
			return 0, class.QueryErr
		}
	case "solve":
		var list []string
		query["judge"] = config.JudgeAC
		err = c.Find(query).Distinct("uid", &list)
		if err != nil {
			return 0, class.QueryErr
		}
		count = len(list)
	default:
		return 0, class.ArgsErr
	}

	return count, nil
}

// POST /Solution?achieve/uid?<uid>
func (this *SolutionModel) Achieve(uid string) ([]int, error) {
	log.Println("Server SolutionModel Achieve")

	err := this.OpenDB()
	if err != nil {
		return nil, class.DBErr
	}
	defer this.CloseDB()

	var list []int
	err = this.DB.C("Solution").Find(bson.M{"uid": uid, "judge": config.JudgeAC}).Sort("pid").Distinct("pid", &list)
	if err != nil {
		return nil, class.OpErr
	}

	return list, nil
}

// POST /Solution?list/offset?<offset>/limit?<limit>/sid?<sid>/pid?<pid>/uid?<uid>/language?<language>/judge?<judge>/module?<module>/mid?<mid>/from?<from>/sort?<resort/sort>
//默认按sid从大到小拍,resort可以设置为从小到大
func (this *SolutionModel) List(args map[string]string) ([]*Solution, error) {
	log.Println("Server SolutionModel List")

	query, err := this.CheckQuery(args)
	if err != nil {
		return nil, class.ArgsErr
	}

	err = this.OpenDB()
	if err != nil {
		return nil, class.DBErr
	}
	defer this.CloseDB()

	sort := "-sid"
	if args["sort"] == "resort" {
		sort = "sid"
	}
	q := this.DB.C("Solution").Find(query).Select(sListSelector).Sort(sort)

	if v, ok := args["offset"]; ok {
		offset, err := strconv.Atoi(v)
		if err != nil {
			return nil, class.ArgsErr
		}
		q = q.Skip(offset)
	}

	if v, ok := args["limit"]; ok {
		limit, err := strconv.Atoi(v)
		if err != nil {
			return nil, class.ArgsErr
		}
		q = q.Limit(limit)
	}

	var list []*Solution
	err = q.All(&list)
	if err != nil {
		return nil, class.QueryErr
	}

	return list, nil
}

func (this *SolutionModel) CheckQuery(args map[string]string) (query bson.M, err error) {
	query = make(bson.M)

	if v, ok := args["sid"]; ok {
		var sid int
		sid, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		query["sid"] = sid
	}
	if v, ok := args["pid"]; ok {
		var pid int
		pid, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		query["pid"] = pid
	}
	if v, ok := args["cid"]; ok {
		var cid int
		cid, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		query["cid"] = cid
	}
	if v, ok := args["uid"]; ok {
		query["uid"] = v
	}
	if v, ok := args["language"]; ok {
		var language int
		language, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		query["language"] = language
	}
	if v, ok := args["judge"]; ok {
		var judge int
		judge, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		query["judge"] = judge
	}
	if v, ok := args["module"]; ok {
		var module int
		module, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		query["module"] = module
	}
	if v, ok := args["mid"]; ok {
		var mid int
		mid, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		query["mid"] = mid
	}
	if v, ok := args["from"]; ok {
		var from int
		from, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		query["sid"] = bson.M{"$gte": from}
	}
	return
}
