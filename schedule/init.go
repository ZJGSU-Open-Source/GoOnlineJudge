package schedule

import (
    "GoOnlineJudge/config"
    "GoOnlineJudge/model"
    "time"
)

type RemoteOJInterface interface {
    Init()
    Host() string
    Ping() error
    GetProblems() error
}

var ROJs = []RemoteOJInterface{&HDUJudger{}, &PKUJudger{}}

func init() {
    go func() {
        for _, oj := range ROJs {
            oj.Init()
        }
        for {
            errCnt := 0
            for _, oj := range ROJs {
            again:
                if err := oj.GetProblems(); err != nil {
                    errCnt++
                    if errCnt > 5 {
                        continue
                    }
                    time.Sleep(10 * time.Second)
                    goto again
                }
            }
            time.Sleep(7 * 24 * time.Hour) //update per week
        }
    }()
    go func() {
        ojModel := &model.OJModel{}
        status := &model.OJStatus{}
        for {
            for _, oj := range ROJs {
                err := oj.Ping()
                status.Name = oj.Host()
                if err != nil {
                    status.Status = config.StatusUnavailable
                    ojModel.Update(status)
                } else {
                    status.Status = config.StatusOk
                    ojModel.Update(status)
                }
            }
            time.Sleep(1 * time.Minute)
        }
    }()
}
