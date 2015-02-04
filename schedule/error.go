package schedule

import (
	"errors"
)

var (
	ErrConnectFailed = errors.New("connection failed")
	ErrNoSuchProblem = errors.New("no such problem")
	ErrMatchFailed   = errors.New("match failed")
)
