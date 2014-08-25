package model

import (
	log "GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model/class"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type Contest struct {
	Cid      int         `json:"cid"bson:"cid"`
	Title    string      `json:"title"bson:"title"`
	Encrypt  int         `json:"encrypt"bson:"encrypt"`
	Argument interface{} `json:"argument"bson:"argument"`
	Type     string      `json:"type"bson:"type"` //the type of Contest,acm Contest or normal exercise

	Start int64 `json:"start"bson:"start"`
	End   int64 `json:"end"bson:"end"`

	Status int    `json:"status"bson:"status"`
	Create string `'json:"create"bson:"create"`

	List []int `json:"list"bson:"list"` //problem list
}

var cDetailSelector = bson.M{"_id": 0}
var cListSelector = bson.M{"_id": 0, "cid": 1, "title": 1, "encrypt": 1, "argument": 1, "start": 1, "end": 1, "status": 1}

type ContestModel struct {
	class.Model
}

// 参数cid，返回指定cid的contest
func (this *ContestModel) Detail(cid int) (*Contest, error) {
	log.Logger.Debug("Server ContestModel Detail")

	err := this.OpenDB()
	if err != nil {
		return nil, DBErr
	}
	defer this.CloseDB()
	one := &Contest{}
	err = this.DB.C("Contest").Find(bson.M{"cid": cid}).Select(cDetailSelector).One(one)
	if err == mgo.ErrNotFound {
		return nil, NotFoundErr
	} else if err != nil {
		return nil, OpErr
	}
	return one, nil
}

// 删除指定cid的contest
func (this *ContestModel) Delete(cid int) error {
	log.Logger.Debug("Server ContestModel Delete")

	err := this.OpenDB()
	if err != nil {
		return DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("Contest").Remove(bson.M{"cid": cid})
	if err == mgo.ErrNotFound {
		return NotFoundErr
	} else if err != nil {
		return OpErr
	}

	return nil
}

// POST /Contest?insert
func (this *ContestModel) Insert(one Contest) error {
	log.Logger.Debug("Server ContestModel Insert")

	err := this.OpenDB()
	if err != nil {
		return DBErr
	}
	defer this.CloseDB()

	one.Status = config.StatusReverse

	one.Create = this.GetTime()
	one.Cid, err = this.GetID("Contest")
	if err != nil {
		return IDErr
	}

	err = this.DB.C("Contest").Insert(&one)
	if err != nil {
		return OpErr
	}

	// b, err := json.Marshal(map[string]interface{}{
	// 	"cid":    one.Cid,
	// 	"status": one.Status,
	// })
	return nil
}

// POST /Contest?update/cid?<cid>
func (this *ContestModel) Update(cid int, ori Contest) error {
	log.Logger.Debug("Server ContestModel Update")

	alt := make(map[string]interface{})
	alt["title"] = ori.Title
	alt["start"] = ori.Start
	alt["end"] = ori.End
	alt["encrypt"] = ori.Encrypt
	alt["Argument"] = ori.Argument
	alt["list"] = ori.List
	alt["type"] = ori.Type

	err := this.OpenDB()
	if err != nil {
		return DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("Contest").Update(bson.M{"cid": cid}, bson.M{"$set": alt})
	if err == mgo.ErrNotFound {
		return NotFoundErr
	} else if err != nil {
		return OpErr
	}

	return nil
}

// POST /Contest?status/cid?<cid>/action?<0/2>
func (this *ContestModel) Status(cid, status int) error {
	log.Logger.Debug("Server ContestModel Status")

	err := this.OpenDB()
	if err != nil {
		return DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("Contest").Update(bson.M{"cid": cid}, bson.M{"$set": bson.M{"status": status}})
	if err == mgo.ErrNotFound {
		return NotFoundErr
	} else if err != nil {
		return OpErr
	}

	return nil
}

// POST /Contest?push/cid?<cid>
func (this *ContestModel) Push(cid int, list []int) error {
	log.Logger.Debug("Server ContestModel Push")

	err := this.OpenDB()
	if err != nil {
		return DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("Contest").Update(bson.M{"cid": cid}, bson.M{"$addToSet": bson.M{"list": bson.M{"$each": list}}})
	if err == mgo.ErrNotFound {
		return NotFoundErr
	} else if err != nil {
		return OpErr
	}

	return nil
}

// POST /Contest?list/type?<type>/offset?<offset>/limit?<limit>/pid?<pid>/title?<title>/
func (this *ContestModel) List(args map[string]string) ([]*Contest, error) {
	log.Logger.Debug("Server ContestModel List")

	query, err := this.CheckQuery(args)
	if err != nil {
		return nil, ArgsErr
	}

	err = this.OpenDB()
	if err != nil {
		return nil, DBErr
	}
	defer this.CloseDB()

	q := this.DB.C("Contest").Find(query).Select(cListSelector).Sort("cid")

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

	var list []*Contest
	err = q.All(&list)
	if err != nil {
		return nil, QueryErr
	}

	return list, nil
}

func (this *ContestModel) CheckQuery(args map[string]string) (query bson.M, err error) {
	query = make(bson.M)

	if v, ok := args["cid"]; ok {
		var cid int
		cid, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		query["cid"] = cid
	}
	if v, ok := args["title"]; ok {
		query["title"] = bson.M{"$regex": bson.RegEx{v, "i"}}
	}
	if v, ok := args["type"]; ok {
		query["type"] = v
	}
	return
}
