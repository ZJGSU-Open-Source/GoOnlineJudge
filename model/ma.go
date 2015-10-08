package model

import (
	"ojapi/model/class"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type migrate struct {
	Version int `json:"version"bson:"version"`
}

type migrateModel struct {
	class.Model
}

var versionListSelector = bson.M{"_id": 0,
	"version": 1,
}

func (m *migrateModel) Update(version int) error {
	err := m.OpenDB()
	if err != nil {
		return DBErr
	}
	defer m.CloseDB()

	alt := make(map[string]interface{})
	alt["version"] = version

	err = m.DB.C("migrate").Update(bson.M{}, bson.M{"$set": alt})
	if err == mgo.ErrNotFound {
		err = m.DB.C("migrate").Insert(&migrate{Version: version})
		if err != nil {
			return OpErr
		}
	} else if err != nil {
		return OpErr
	}

	return nil
}

func (m *migrateModel) Get() (*migrate, error) {
	err := m.OpenDB()
	if err != nil {
		return nil, DBErr
	}
	defer m.CloseDB()

	q := m.DB.C("migrate").Find(nil).Select(versionListSelector)

	var list []*migrate
	err = q.All(&list)
	if err != nil {
		return nil, OpErr
	} else if len(list) == 0 {
		return &migrate{0}, nil
	}

	return list[0], nil

}
