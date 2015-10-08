package model

import (
	"ojapi/config"
)

func init() {
	setup()
}

type Migrator func()

func setup() {
	var Migrations = []Migrator{
		Migrate_2015_07_09,
	}

	mgt := &migrateModel{}
	mi, err := mgt.Get()
	if err != nil {
		return
	}

	if len(Migrations) < mi.Version {
		return
	}

	for _, m := range Migrations[mi.Version:] {
		m()
	}

	mgt.Update(len(Migrations))

}

func Migrate_2015_07_09() {
	userModel := &UserModel{}
	user := &User{}
	user.Uid = "admin"
	user.Pwd = "admin"
	user.Nick = "admin"
	user.Privilege = config.PrivilegeAD

	userModel.Insert(*user)
}
