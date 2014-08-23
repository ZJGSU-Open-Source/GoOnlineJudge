package model

import (
	"errors"
)

var DBErr = errors.New("DB Open Error")
var ArgsErr = errors.New("Args Error")
var NotFoundErr = errors.New("Not Found")
var OpErr = errors.New("Operat Error")
var IDErr = errors.New("Get ID Error")
var QueryErr = errors.New("Query Error")
var EncryptErr = errors.New("Encrypt Error")
