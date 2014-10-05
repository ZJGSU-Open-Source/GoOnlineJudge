package model

import (
	"errors"
	"golog"
	"os"
)

var (
	//数据库打开错误
	DBErr = errors.New("DB Open Error")

	//函数参数错误
	ArgsErr = errors.New("Args Error")

	//没有查询到指定数据
	NotFoundErr = errors.New("Not Found")

	//更新错误
	OpErr = errors.New("Operat Error")

	//ID生成错误
	IDErr = errors.New("Get ID Error")

	//查询或查询参数错误
	QueryErr = errors.New("Query Error")

	//密码加密错误
	EncryptErr = errors.New("Encrypt Error")

	ExistErr = errors.New("Id has existed")
)

var logger *golog.Log

func init() {
	logger = golog.NewLog(os.Stdout, golog.Ldebug|golog.Linfo)
}
