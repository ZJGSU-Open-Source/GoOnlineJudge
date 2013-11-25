package models

import (
	"GoOnlineJudge/config"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type UserModel struct {
}

func (this *UserModel) SignIn(uid, pwd string) bool {
	session, err := mgo.Dial(config.Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(config.DB).C("user")
	count, _ := c.Find(bson.M{"uid": uid, "pwd": pwd}).Count()
	if count > 0 {
		return true
	} else {
		return false
	}
}
