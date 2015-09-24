package model

import (
    "GoOnlineJudge/config"
    "GoOnlineJudge/model/class"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"

    "html/template"
    "strconv"
)

type News struct {
    Nid     int           `json:"nid"bson:"nid"`
    Title   string        `json:"title"bson:"title"`
    Content template.HTML `json:"content"bson:"content"`

    Status int    `json:"status"bson:"status"`
    Create string `json:"create"bson:'create'`
}

var nDetailSelector = bson.M{"_id": 0}
var nListSelector = bson.M{"_id": 0,
    "nid":     1,
    "title":   1,
    "content": 1,
    "status":  1,
    "create":  1}

type NewsModel struct {
    class.Model
}

// 获取指定nid的news
func (this *NewsModel) Detail(nid int) (*News, error) {
    logger.Debug("Server NewsModel Detail")

    err := this.OpenDB()
    if err != nil {
        return nil, DBErr
    }
    defer this.CloseDB()

    one := &News{}
    err = this.DB.C("News").Find(bson.M{"nid": nid}).Select(nDetailSelector).One(&one)
    if err == mgo.ErrNotFound {
        return nil, NotFoundErr
    } else if err != nil {
        return nil, OpErr
    }
    return one, nil
}

// 删除指定nid的news
func (this *NewsModel) Delete(nid int) error {
    logger.Debug("Server NewsModel Delete")

    err := this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    err = this.DB.C("News").Remove(bson.M{"nid": nid})
    if err == mgo.ErrNotFound {
        return NotFoundErr
    } else if err != nil {
        return OpErr
    }

    return nil
}

// 插入一个新的news，不能指定status和create
func (this *NewsModel) Insert(one News) error {
    logger.Debug("Server NewsModel Insert")

    err := this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    one.Status = config.StatusReverse
    one.Create = this.GetTime()
    one.Nid, err = this.GetID("News")
    if err != nil {
        return IDErr
    }

    err = this.DB.C("News").Insert(&one)
    if err != nil {
        return OpErr
    }

    return nil
}

// 更新指定nid的news
func (this *NewsModel) Update(nid int, ori News) error {
    logger.Debug("Server NewsModel Update")

    alt := make(map[string]interface{})
    alt["title"] = ori.Title
    alt["content"] = ori.Content

    err := this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    err = this.DB.C("News").Update(bson.M{"nid": nid}, bson.M{"$set": alt})
    if err == mgo.ErrNotFound {
        return NotFoundErr
    } else if err != nil {
        return OpErr
    }

    return nil
}

// 更新指定的news的status
func (this *NewsModel) Status(nid, status int) error {
    logger.Debug("Server NewsModel Status")

    err := this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    err = this.DB.C("News").Update(bson.M{"nid": nid}, bson.M{"$set": bson.M{"status": status}})
    if err == mgo.ErrNotFound {
        return NotFoundErr
    } else if err != nil {
        return OpErr
    }

    return nil
}

// 列出由offset，limit指定的news
func (this *NewsModel) List(args map[string]string, offset, limit int) ([]*News, error) {
    logger.Debug("Server NewsModel List")

    query, err := this.CheckQuery(args)
    if err != nil {
        return nil, ArgsErr
    }

    err = this.OpenDB()
    if err != nil {
        return nil, DBErr
    }
    defer this.CloseDB()

    q := this.DB.C("News").Find(query).Select(nListSelector).Sort("-nid")

    if offset > -1 {
        q = q.Skip(offset)
    }

    if limit > -1 {
        q = q.Limit(limit)
    }

    var list []*News
    err = q.All(&list)
    if err != nil {
        return nil, OpErr
    }

    return list, nil
}

func (this *NewsModel) CheckQuery(args map[string]string) (query bson.M, err error) {
    query = make(bson.M)

    if v, ok := args["nid"]; ok {
        var sid int
        sid, err = strconv.Atoi(v)
        if err != nil {
            return
        }
        query["nid"] = sid
    }

    if v, ok := args["status"]; ok {
        var status int
        status, err = strconv.Atoi(v)
        if err != nil {
            return
        }
        query["status"] = status
    }

    return
}
