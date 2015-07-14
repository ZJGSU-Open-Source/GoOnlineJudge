package class

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DBLasting = false

type ids struct {
	Name string `json:"name"bson:"name"`
	Id   int    `json:"id"bson:"id"`
}

type Model struct {
	Session *mgo.Session
	DB      *mgo.Database
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

var (
	username string
	password string
	host     string
	port     string
	instance string
)

func Config() {

	username = os.Getenv("MONGODB_USERNAME")
	password = os.Getenv("MONGODB_PASSWORD")
	host = os.Getenv("MONGODB_PORT_27017_TCP_ADDR")

	if len(host) == 0 {
		host = "localhost"
	}

	port = os.Getenv("MONGODB_PORT_27017_TCP_PORT")
	if len(port) == 0 {
		port = "27017"
	}

	instance = os.Getenv("MONGODB_INSTANCE_NAME")
	if len(instance) == 0 {
		instance = "oj"
	}
}

func (m *Model) OpenDB() error {
	conn := ""
	if len(username) > 0 {
		conn += username

		if len(password) > 0 {
			conn += ":" + password
		}

		conn += "@"
	}

	conn += fmt.Sprintf("%s:%s/%s", host, port, instance)

	var err error

	m.Session, err = mgo.Dial(conn)
	if err != nil {
		fmt.Println(err)
		return err
	}

	m.DB = m.Session.DB(instance)
	return nil
}
