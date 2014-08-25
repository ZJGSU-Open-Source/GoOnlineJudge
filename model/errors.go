package model

import (
	"errors"
)

//数据库打开错误
var DBErr = errors.New("DB Open Error")

//函数参数错误
var ArgsErr = errors.New("Args Error")

//没有查询到指定数据
var NotFoundErr = errors.New("Not Found")

//更新错误
var OpErr = errors.New("Operat Error")

//ID生成错误
var IDErr = errors.New("Get ID Error")

//查询或查询参数错误
var QueryErr = errors.New("Query Error")

//密码加密错误
var EncryptErr = errors.New("Encrypt Error")
