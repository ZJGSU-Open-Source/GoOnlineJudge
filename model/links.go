package model

import (
    // "GoOnlineJudge/config"
    "GoOnlineJudge/model/class"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"

    "strconv"
    "time"
)

type Link struct {
    Lid int `json:"lid"bson:"lid"`

    Uid   string `json:"uid"bson:"uid"`
    Title string `json:"title"bson:"title"`
    Link  string `json:"link"bson:"link"`

    Create int64 `json:"create"bson:"create"`
}

var lDetailSelector = bson.M{"_id": 0}
var lListSelector = bson.M{"_id": 0,
    "lid":   1,
    "uid":   1,
    "title": 1,
    "link":  1,
}

type LinkModel struct {
    class.Model
}

// 查询指定sid的solution的所有详细信息
func (this *LinkModel) Detail(lid int) (*Link, error) {
    logger.Debug("Server link Detail")
    err := this.OpenDB()
    if err != nil {
        return nil, DBErr
    }
    defer this.CloseDB()

    var one Link
    err = this.DB.C("Link").Find(bson.M{"lid": lid}).Select(lDetailSelector).One(&one)
    if err == mgo.ErrNotFound {
        return nil, NotFoundErr
    } else if err != nil {
        return nil, OpErr
    }

    return &one, nil
}

// 删除指定sid的solution
func (this *LinkModel) Delete(lid int) error {
    logger.Debug("Server SolutionModel Delete")

    err := this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    err = this.DB.C("Link").Remove(bson.M{"lid": lid})
    if err == mgo.ErrNotFound {
        return NotFoundErr
    } else if err != nil {
        return OpErr
    }

    return nil
}

// 插入一个新的solution，不能指定create和sid
func (this *LinkModel) Insert(one Link) (int, error) {
    logger.Debug("Server Link Insert")

    err := this.OpenDB()
    if err != nil {
        return 0, DBErr
    }
    defer this.CloseDB()

    one.Create = time.Now().Unix()
    one.Lid, err = this.GetID("Link")
    if err != nil {
        return 0, IDErr
    }

    err = this.DB.C("Link").Insert(&one)
    if err != nil {
        return 0, OpErr
    }

    return one.Lid, nil
}

// 更新指定sid的solution，可更新参数包括judge、time、memory三个
func (this *LinkModel) Update(lid int, ori Link) error {
    logger.Debug("Server Link Update")

    alt := make(map[string]interface{})
    alt["title"] = ori.Title
    alt["link"] = ori.Link

    err := this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    err = this.DB.C("Link").Update(bson.M{"lid": lid}, bson.M{"$set": alt})
    if err == mgo.ErrNotFound {
        return NotFoundErr
    } else if err != nil {
        return OpErr
    }

    return nil
}

// // 更新指定sid的solution状态，状态值由status指定
// func (this *LinkModel) Status(sid, status int) error {
//     logger.Debug("Server SolutionModel Status")

//     err := this.OpenDB()
//     if err != nil {
//         return DBErr
//     }
//     defer this.CloseDB()

//     err = this.DB.C("Solution").Update(bson.M{"sid": sid}, bson.M{"$inc": bson.M{"status": status}})
//     if err == mgo.ErrNotFound {
//         return NotFoundErr
//     } else if err != nil {
//         return OpErr
//     }

//     return nil
// }

// 计数由参数args指定的solution个数，参数包括pid、uid、module、mid、action
func (this *LinkModel) Count(args map[string]string) (int, error) {
    logger.Debug("Server Links Count")

    query, err := this.CheckQuery(args)
    if err != nil {
        return 0, ArgsErr
    }

    err = this.OpenDB()
    if err != nil {
        return 0, DBErr
    }
    defer this.CloseDB()

    var count int
    c := this.DB.C("Link")

    count, err = c.Find(query).Count()
    if err != nil {
        return 0, OpErr
    }

    return count, nil
}

// 列出符合参数args的Solution，参数包括offset、limit、sid、pid、uid、language、judge、module、mid、from、sort
//默认按sid从大到小排,sort参数设为resort可以设置为从小到大
func (this *LinkModel) List(args map[string]string) ([]*Link, error) {
    logger.Debug("Server LinkModel List")

    query, err := this.CheckQuery(args)
    if err != nil {
        return nil, ArgsErr
    }

    err = this.OpenDB()
    if err != nil {
        return nil, DBErr
    }
    defer this.CloseDB()

    sort := "-lid"
    if args["sort"] == "resort" {
        sort = "lid"
    }
    q := this.DB.C("Link").Find(query).Select(lListSelector).Sort(sort)

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

    var list []*Link
    err = q.All(&list)
    if err != nil {
        return nil, QueryErr
    }

    return list, nil
}

func (this *LinkModel) CheckQuery(args map[string]string) (query bson.M, err error) {
    query = make(bson.M)

    if v, ok := args["lid"]; ok {
        var lid int
        lid, err = strconv.Atoi(v)
        if err != nil {
            return
        }
        query["lid"] = lid
    }
    if v, ok := args["uid"]; ok {
        query["uid"] = v
    }
    if v, ok := args["title"]; ok {
        query["title"] = v
    }
    if v, ok := args["from"]; ok {
        var from int
        from, err = strconv.Atoi(v)
        if err != nil {
            return
        }
        query["lid"] = bson.M{"$gte": from}
    }
    return
}
