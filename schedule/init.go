package schedule

import (
	"time"
)

type RemoteOJInterface interface {
	Init()
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

}
