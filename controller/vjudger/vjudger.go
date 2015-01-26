package vjudger

type User struct {
	Uid       string
	Vid       int
	OJ        string
	Result    int
	CE        string
	Code      string
	Time      int64
	Mem       int64
	Lang      int
	ErrorCode int
}

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
