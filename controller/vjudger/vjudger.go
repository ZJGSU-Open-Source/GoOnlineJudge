package vjudger

import (
	"time"
)

type User struct {
	Uid string
	Sid int
	Vid int

	OJ     string
	Result int
	CE     string
	Code   string
	Time   int
	Mem    int
	Lang   int
	Length int

	ErrorCode  int
	SubmitTime time.Time
}

const MAX_WaitTime = 120

type Vjudger interface {
	Init(*User) error
	Login(*User) error
	Submit(*User) error
	GetStatus(*User) error
	Run(*User) error
	Match(string) bool
}

func (u *User) NewSolution() {

}

func (u *User) UpdateSolution() {

}

var VJs = []Vjudger{&HDUJudger{}}

func Judge(u *User) {
	for _, vj := range VJs {
		if vj.Match(u.OJ) {
			vj.Run(u)
			break
		}
	}
}
