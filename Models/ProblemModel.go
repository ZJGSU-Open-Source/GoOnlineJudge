package models

import (
	"GoOnlineJudge/config"
	"labix.org/v2/mgo"
)

type Result struct {
	Pid   int
	Title string
}

type ProblemModel struct {
}

func (this *ProblemModel) List() []Result {
	session, err := mgo.Dial(config.Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB(config.DB).C("problem")

	r := []Result{}
	c.Find(nil).All(&r)
	return r
}
