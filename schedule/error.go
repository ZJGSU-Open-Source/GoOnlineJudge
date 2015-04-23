package schedule

import (
    "errors"
)

var (
    ErrConnectFailed = errors.New("connection failed")
    ErrNoSuchProblem = errors.New("no such problem")
    ErrMatchFailed   = errors.New("match failed")
    ErrResponse      = errors.New("can't get response")
)

const (
    StatusReverse   = 0 //不可用
    StatusIncon     = 1 //正在比赛中
    StatusAvailable = 2 //可用
    StatusPending   = 3 //等待
    StatusRunning   = 4 //进行中
    StatusEnding    = 5 //结束
)
