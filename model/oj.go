package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ojapi/model/class"
)

type OJStatus struct {
	Name   string `json:"name"bson:"name"`
	Status int    `json:"status"bson:"status"`
}

type OJModel struct {
	class.Model
}

var statusListSelector = bson.M{"_id": 0,
	"name":   1,
	"status": 1,
}

func (o *OJModel) Update(status *OJStatus) error {

	if status == nil {
		return ArgsErr
	}
	err := o.OpenDB()
	if err != nil {
		return DBErr
	}
	defer o.CloseDB()

	alt := make(map[string]interface{})
	alt["status"] = status.Status

	err = o.DB.C("OJStatus").Update(bson.M{"name": status.Name}, bson.M{"$set": alt})
	if err == mgo.ErrNotFound {
		err = o.DB.C("OJStatus").Insert(status)
		if err != nil {
			return OpErr
		}
	} else if err != nil {
		return OpErr
	}

	return nil
}

func (o *OJModel) List() ([]*OJStatus, error) {
	err := o.OpenDB()
	if err != nil {
		return nil, DBErr
	}
	defer o.CloseDB()

	q := o.DB.C("OJStatus").Find(nil).Select(statusListSelector).Sort("name")

	var list []*OJStatus
	err = q.All(&list)
	if err != nil {
		return nil, OpErr
	}
	return list, nil
}
