package model

import (
    "GoOnlineJudge/config"
    "GoOnlineJudge/model/class"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "strconv"
)

const IPCNT = 5

type User struct {
    Uid string `json:"uid"bson:"uid"`
    Pwd string `json:"pwd"bson:"pwd"`

    Nick   string `json:"nick"bson:"nick"`
    Mail   string `json:"mail"bson:"mail"`
    School string `json:"school"bson:"school"`
    Motto  string `json:"motto"bson:"motto"`

    Privilege int `json:"privilege"bson:"privilege"`
    //module指定user类型，有普通normal(0)，比赛用team(1)两种类型
    Module int `json:"module"bson:"module"`

    Solve  int `json:"solve"bson:"solve"`
    Submit int `json:"submit"bson:"submit"`

    ShareCode  bool          `json:"share_code"bson:"share_code"`
    Status     int           `json:"status"bson:"status"`
    Create     string        `json:"create"bson:"create"`
    IPRecord   [IPCNT]string `json:"iprecord"bson:"iprecord"` //记录ip地址
    TimeRecord [IPCNT]int64  `json:"timerecord"bson:"timerecord"`
}

var uDetailSelector = bson.M{"_id": 0}
var uListSelector = bson.M{"_id": 0,
    "uid":        1,
    "nick":       1,
    "motto":      1,
    "privilege":  1,
    "solve":      1,
    "submit":     1,
    "status":     1,
    "iprecord":   1,
    "timerecord": 1}

type UserModel struct {
    class.Model
}

// 用户登入验证，需要提供uid和pwd两个参数，
// 如果用户存在并且uid和pwd匹配，则返回一个user
func (this *UserModel) Login(uid, pwd string) (*User, error) {
    logger.Debug("Server UserModel Login")

    var err error
    pwd, err = this.EncryptPassword(pwd)
    if err != nil {
        return nil, EncryptErr
    }

    err = this.OpenDB()
    if err != nil {
        return nil, DBErr
    }
    defer this.CloseDB()

    var alt User
    err = this.DB.C("User").Find(bson.M{"uid": uid}).Select(uDetailSelector).One(&alt)
    if err == mgo.ErrNotFound {
        return nil, NotFoundErr
    } else if err != nil {
        return nil, OpErr
    }

    if pwd == alt.Pwd {
        logger.Debug("Server UserModel Login Successfully")
        return &User{
            Uid:       alt.Uid,
            Nick:      alt.Nick,
            Privilege: alt.Privilege,
            Status:    alt.Status,
        }, nil
    }
    logger.Debug("Server UserModel Login Failed")
    return &User{
        Uid:       "",
        Nick:      "",
        Privilege: config.PrivilegeNA,
        Status:    config.StatusReverse,
    }, nil
}

func (this *UserModel) RecordIP(uid, IP string, time int64) error {
    err := this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    var alt User
    err = this.DB.C("User").Find(bson.M{"uid": uid}).Select(uDetailSelector).One(&alt)
    if err == mgo.ErrNotFound {
        return NotFoundErr
    } else if err != nil {
        return OpErr
    }

    ipRecord := alt.IPRecord
    timeRecord := alt.TimeRecord

    ipcnt := len(ipRecord)
    if ipcnt < IPCNT {
        ipRecord[ipcnt] = IP
        timeRecord[ipcnt] = time
    } else {
        for i := 0; i < IPCNT-1; i++ {
            ipRecord[i] = ipRecord[i+1]
            timeRecord[i] = timeRecord[i+1]
        }
        ipRecord[IPCNT-1] = IP
        timeRecord[IPCNT-1] = time
    }

    err = this.DB.C("User").Update(bson.M{"uid": uid}, bson.M{"$set": bson.M{"iprecord": ipRecord, "timerecord": timeRecord}})
    if err == mgo.ErrNotFound {
        return NotFoundErr
    } else if err != nil {
        return OpErr
    }

    return nil

}

//这个函数貌似没干什么事啊==
func (this *UserModel) Logout() {
    logger.Debug("Server UserModel Logout")

    // var one User
    // err := this.LoadJson(r.Body, &one)
    // if err != nil {
    // 	http.Error(w, "load error", 400)
    // 	return
    // }

    // w.WriteHeader(200)
}

//设定用户密码，需提供uid和pwd
func (this *UserModel) Password(uid, pwd string) error {
    logger.Debug("Server UserModel Password")

    pwd, err := this.EncryptPassword(pwd)
    if err != nil {
        return EncryptErr
    }

    alt := make(map[string]interface{})
    alt["pwd"] = pwd

    err = this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    err = this.DB.C("User").Update(bson.M{"uid": uid}, bson.M{"$set": alt})
    if err == mgo.ErrNotFound {
        return NotFoundErr
    } else if err != nil {
        return OpErr
    }

    return nil
}

// 设定指定uid用户的权限
func (this *UserModel) Privilege(uid string, privilege int) error {
    logger.Debug("Server UserModel Privilege")

    alt := make(map[string]interface{})
    alt["privilege"] = privilege

    err := this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    err = this.DB.C("User").Update(bson.M{"uid": uid}, bson.M{"$set": alt})
    if err == mgo.ErrNotFound {
        return NotFoundErr
    } else if err != nil {
        return OpErr
    }

    return nil
}

// 获得指定uid用户的所有信息
func (this *UserModel) Detail(uid string) (*User, error) {
    logger.Debug("Server UserModel Detail")

    err := this.OpenDB()
    if err != nil {
        return nil, DBErr
    }
    defer this.CloseDB()

    var one User
    err = this.DB.C("User").Find(bson.M{"uid": uid}).Select(uDetailSelector).One(&one)
    if err == mgo.ErrNotFound {
        return nil, NotFoundErr
    } else if err != nil {
        return nil, OpErr
    }

    one.Pwd = "" //防止密码泄露
    return &one, nil
}

// 删除指定uid用户
func (this *UserModel) Delete(uid string) error {
    logger.Debug("Server UserModel Delete")

    err := this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    err = this.DB.C("User").Remove(bson.M{"uid": uid})
    if err == mgo.ErrNotFound {
        return NotFoundErr
    } else if err != nil {
        return OpErr
    }

    return nil
}

// 插入一个新的user
func (this *UserModel) Insert(one User) error {
    logger.Debug("Server UserModel Insert")

    if _, err := this.Detail(one.Uid); err != NotFoundErr { //测试uid是否已经存在
        logger.Debug(err)
        return ExistErr
    }

    var err error
    one.Pwd, err = this.EncryptPassword(one.Pwd)
    if err != nil {
        return EncryptErr
    }

    err = this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    one.Solve = 0
    one.Submit = 0
    one.Status = config.StatusAvailable
    one.Create = this.GetTime()
    one.ShareCode = true

    err = this.DB.C("User").Insert(&one)
    if err != nil {
        return OpErr
    }

    // b, err := json.Marshal(map[string]interface{}{
    // 	"uid":       one.Uid,
    // 	"privilege": one.Privilege,
    // 	"status":    one.Status,
    // })

    return nil
}

// 更新用户信息（nick，mail，scholl，motto）
func (this *UserModel) Update(uid string, ori User) error {
    logger.Debug("Server UserModel Update")

    alt := make(map[string]interface{})
    alt["nick"] = ori.Nick
    alt["mail"] = ori.Mail
    alt["school"] = ori.School
    alt["motto"] = ori.Motto
    alt["share_code"] = ori.ShareCode

    err := this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    err = this.DB.C("User").Update(bson.M{"uid": uid}, bson.M{"$set": alt})
    if err == mgo.ErrNotFound {
        return NotFoundErr
    } else if err != nil {
        return OpErr
    }

    return nil
}

// 更新用户状态
func (this *UserModel) Status(uid string) error {
    logger.Debug("Server UserModel Status")

    err := this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    err = this.DB.C("User").Update(bson.M{"uid": uid}, bson.M{"$inc": bson.M{"status": 1}})
    if err == mgo.ErrNotFound {
        return NotFoundErr
    } else if err != nil {
        return OpErr
    }

    return nil
}

// 用户做题记录，uid:<uid>,action:<solve/submit>
func (this *UserModel) Record(uid string, solve int, submit int) error {
    logger.Debug("Server UserModel Submit")

    err := this.OpenDB()
    if err != nil {
        return DBErr
    }
    defer this.CloseDB()

    err = this.DB.C("User").Update(bson.M{"uid": uid}, bson.M{"$set": bson.M{"solve": solve, "submit": submit}})
    if err == mgo.ErrNotFound {
        return NotFoundErr
    } else if err != nil {
        return OpErr
    }
    return nil
}

// 列出用户 offset:<offset>,limit:<limit>,uid:<uid>,nick:<nick>
func (this *UserModel) List(args map[string]string) ([]*User, error) {
    logger.Debug("Server UserModel List")

    query, err := this.CheckQuery(args)
    if err != nil {
        return nil, ArgsErr
    }

    err = this.OpenDB()
    if err != nil {
        return nil, DBErr
    }
    defer this.CloseDB()

    q := this.DB.C("User").Find(query).Select(uListSelector).Sort("-solve", "submit")

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

    var list []*User
    err = q.All(&list)
    if err != nil {
        return nil, QueryErr
    }

    return list, nil
}

func (this *UserModel) CheckQuery(args map[string]string) (query bson.M, err error) {
    query = make(bson.M)

    query["module"] = 0
    if v, ok := args["uid"]; ok {
        query["uid"] = v
    }
    if v, ok := args["nick"]; ok {
        query["nick"] = bson.M{"$regex": bson.RegEx{v, "i"}}
    }
    return
}
