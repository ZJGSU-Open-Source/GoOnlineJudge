package class

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"time"
)

const DBHost = "127.0.0.1"
const DBName = "oj"
const DBLasting = false

type ids struct {
	Name string `json:"name"bson:"name"`
	Id   int    `json:"id"bson:"id"`
}

type Model struct {
	Session *mgo.Session
	DB      *mgo.Database
}

func (this *Model) OpenDB() (err error) {
	this.Session, err = mgo.Dial(DBHost)
	if err != nil {
		return
	}

	this.DB = this.Session.DB(DBName)
	return
}

func (this *Model) CloseDB() {
	if !DBLasting {
		this.Session.Close()
	}
}

func (this *Model) GetID(c string) (id int, err error) {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"id": 1}},
		Upsert:    true,
		ReturnNew: true,
	}

	var one ids
	_, err = this.DB.C("ids").Find(bson.M{"name": c}).Apply(change, &one)
	id = one.Id
	return
}

func (this *Model) GetTime() (ft string) {
	t := time.Now().Unix()
	ft = time.Unix(t, 0).Format("2006-01-02 15:04:05")
	return
}

func (this *Model) EncryptPassword(str string) (pwd string, err error) {
	m := md5.New()
	_, err = io.WriteString(m, str)
	p1 := fmt.Sprintf("%x", m.Sum(nil))
	s := sha1.New()
	_, err = io.WriteString(s, str)
	p2 := fmt.Sprintf("%x", s.Sum(nil))
	pwd = p1 + p2
	return
}
