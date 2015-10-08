package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ojapi/model/class"
)

type VIds struct {
	Name string `json:"name"bson:"name"`
	Id   int    `json:"id"bson:"id"`
}

type VIdsModel struct {
	class.Model
}

func (v *VIdsModel) GetLastID(c string) (id int, err error) {
	err = v.OpenDB()
	if err != nil {
		return 0, DBErr
	}
	defer v.CloseDB()

	var ids VIds
	err = v.DB.C("VIds").Find(bson.M{"name": c}).One(&ids)
	if err == mgo.ErrNotFound {
		return 0, nil
	} else if err != nil {
		return 0, OpErr
	}
	return ids.Id, nil
}

func (v *VIdsModel) SetLastID(c string, id int) error {
	err := v.OpenDB()
	if err != nil {
		return DBErr
	}
	defer v.CloseDB()

	err = v.DB.C("VIds").Update(bson.M{"name": c}, bson.M{"$set": bson.M{"id": id}})
	if err == mgo.ErrNotFound {
		ids := VIds{Name: c, Id: id}
		err = v.DB.C("VIds").Insert(&ids)
		return err
	} else if err != nil {
		return OpErr
	}

	return nil
}
