package model

import (
	log "GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model/class"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"strconv"
)

type Problem struct {
	Pid int `json:"pid"bson:"pid"`

	Time    int    `json:"time"bson:"time"`
	Memory  int    `json:"memory"bson:"memory"`
	Special int    `json:"special"bson:"special"`
	Expire  string `json:"expire"bson:"expire"`

	Title       string        `json:"title"bson:"title"`
	Description template.HTML `json:"description"bson:"description"`
	Input       template.HTML `json:"input"bson:"input"`
	Output      template.HTML `json:"output"bson:"output"`
	Source      string        `json:"source"bson:"source"`
	Hint        string        `json:"hint"bson:"hint"`

	In  string `json:"in"bson:"in"`
	Out string `json:"out"bson:"out"`

	Solve  int `json:"solve"bson:"solve"`
	Submit int `json:"submit"bson:"submit"`

	Status int    `json:"status"bson:"status"`
	Create string `json:"create"bson:"create"`
}

var pDetailSelector = bson.M{"_id": 0}
var pListSelector = bson.M{"_id": 0, "pid": 1, "title": 1, "source": 1, "solve": 1, "submit": 1, "expire": 1, "status": 1}

type ProblemModel struct {
	class.Model
}

// 修改指定pid的problem的expire值
func (this *ProblemModel) Expire(pid int, expire string) error {
	log.Logger.Debug("Server ProblemModel Expire")

	alt := make(map[string]interface{})
	alt["expire"] = expire

	err := this.OpenDB()
	if err != nil {
		return DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("Problem").Update(bson.M{"pid": pid, "expire": bson.M{"$lt": expire}}, bson.M{"$set": alt})
	if err == mgo.ErrNotFound {
		return NotFoundErr
	} else if err != nil {
		return OpErr
	}

	return nil
}

// 查询指定pid的problem的详细信息
func (this *ProblemModel) Detail(pid int) (*Problem, error) {
	log.Logger.Debug("Server ProblemModel Detail")

	err := this.OpenDB()
	if err != nil {
		return nil, DBErr
	}
	defer this.CloseDB()

	var one Problem
	err = this.DB.C("Problem").Find(bson.M{"pid": pid}).Select(pDetailSelector).One(&one)
	if err == mgo.ErrNotFound {
		return nil, NotFoundErr
	} else if err != nil {
		return nil, OpErr
	}

	return &one, nil
}

// 删除指定pid的problem
func (this *ProblemModel) Delete(pid int) error {
	log.Logger.Debug("Server ProblemModel Delete")

	err := this.OpenDB()
	if err != nil {
		return DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("Problem").Remove(bson.M{"pid": pid})
	if err == mgo.ErrNotFound {
		return NotFoundErr
	} else if err != nil {
		return OpErr
	}

	return nil
}

// 插入一个新的problem，不可指定solve,submit,status,create,exprie和pid参数，
// 成功则返回新的pid
func (this *ProblemModel) Insert(one Problem) (int, error) {
	log.Logger.Debug("Server ProblemModel Insert")

	err := this.OpenDB()
	if err != nil {
		return 0, DBErr
	}
	defer this.CloseDB()

	one.Solve = 0
	one.Submit = 0
	one.Status = config.StatusReverse
	one.Create = this.GetTime()
	one.Expire = one.Create
	one.Pid, err = this.GetID("Problem")
	if err != nil {
		return 0, IDErr
	}

	err = this.DB.C("Problem").Insert(&one)
	if err != nil {
		return 0, OpErr
	}

	// b, err := json.Marshal(map[string]int{
	// 	"pid":    one.Pid,
	// 	"status": one.Status,
	// })
	// if err != nil {
	// 	http.Error(w, "json error", 500)
	// 	return
	// }

	return one.Pid, nil
}

// 更新指定pid的problem
func (this *ProblemModel) Update(pid int, ori Problem) error {
	log.Logger.Debug("Server ProblemModel Update")

	alt := make(map[string]interface{})
	alt["title"] = ori.Title
	alt["description"] = ori.Description
	alt["input"] = ori.Input
	alt["output"] = ori.Output
	alt["source"] = ori.Source
	alt["hint"] = ori.Hint
	alt["in"] = ori.In
	alt["out"] = ori.Out
	alt["time"] = ori.Time
	alt["memory"] = ori.Memory
	alt["special"] = ori.Special

	err := this.OpenDB()
	if err != nil {
		return DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("Problem").Update(bson.M{"pid": pid}, bson.M{"$set": alt})
	if err == mgo.ErrNotFound {
		return NotFoundErr
	} else if err != nil {
		return OpErr
	}

	return nil
}

// 更新指定pid的problem的状态，状态值由status指定
func (this *ProblemModel) Status(pid, status int) error {
	log.Logger.Debug("Server ProblemModel Status")

	err := this.OpenDB()
	if err != nil {
		return DBErr
	}
	defer this.CloseDB()

	set := make(map[string]interface{})
	set["status"] = status
	err = this.DB.C("Problem").Update(bson.M{"pid": pid}, bson.M{"$set": set})
	if err == mgo.ErrNotFound {
		return NotFoundErr
	} else if err != nil {
		return OpErr
	}

	return nil
}

// 记录problem的solve和submit，每次记录时submit值总会加1，而slolve则由action指定是否记录
func (this *ProblemModel) Record(pid int, action string) error {
	log.Logger.Debug("Server ProblemModel Record")

	var inc int
	switch action {
	case "solve":
		inc = 1
	case "submit":
		inc = 0
	default:
		return ArgsErr
	}

	err := this.OpenDB()
	if err != nil {
		return DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("Problem").Update(bson.M{"pid": pid}, bson.M{"$inc": bson.M{"solve": inc, "submit": 1}})
	if err == mgo.ErrNotFound {
		return NotFoundErr
	} else if err != nil {
		return OpErr
	}

	return nil
}

// 列出符合参数args的所有problem，offset?<offset>/limit?<limit>/pid?<pid>/title?<title>/source?<source>
func (this *ProblemModel) List(args map[string]string) ([]*Problem, error) {
	log.Logger.Debug("Server ProblemModel List")

	query, err := this.CheckQuery(args)
	if err != nil {
		return nil, ArgsErr
	}

	err = this.OpenDB()
	if err != nil {
		return nil, DBErr
	}
	defer this.CloseDB()

	q := this.DB.C("Problem").Find(query).Select(pListSelector).Sort("pid")

	if v, ok := args["offset"]; ok {
		offset, err := strconv.Atoi(v)
		if err != nil {
			return nil, ArgsErr
		}
		q = q.Skip(offset)
	}

	if v, ok := args["limit"]; ok {
		limit, err := strconv.Atoi(v)
		if err != nil {
			return nil, ArgsErr
		}
		q = q.Limit(limit)
	}

	var list []*Problem
	err = q.All(&list)
	if err != nil {
		return nil, OpErr
	}

	return list, nil
}

// 计数符合参数args的problem的个数，count/title?<title>/source?<source>/status?<status>
func (this *ProblemModel) Count(args map[string]string) (int, error) {
	log.Logger.Debug("ProblemModel Count")

	query, err := this.CheckQuery(args)
	if err != nil {
		return 0, ArgsErr
	}

	err = this.OpenDB()
	if err != nil {
		return 0, DBErr
	}
	defer this.CloseDB()

	count, err := this.DB.C("Problem").Find(query).Count()
	if err != nil {
		return 0, QueryErr
	}

	return count, nil
}

//检查查询参数args，成功则返回一个query
func (this *ProblemModel) CheckQuery(args map[string]string) (query bson.M, err error) {
	query = make(bson.M)

	if v, ok := args["pid"]; ok {
		var pid int
		pid, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		query["pid"] = pid
	}
	if v, ok := args["title"]; ok {
		query["title"] = bson.M{"$regex": bson.RegEx{v, "i"}}
	}
	if v, ok := args["source"]; ok {
		query["source"] = bson.M{"$regex": bson.RegEx{v, "i"}}
	}

	return
}
