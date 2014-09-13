package class

import (
	"GoOnlineJudge/config"
	"strconv"
	"time"
)

var specialArr = []string{"Standard", "Special"}
var judgeArr = []string{"Pengding", "Running & Judging", "Compile Error", "Accepted", "Runtime Error",
	"Wrong Answer", "Time Limit Exceeded", "Memory Limit Exceeded", "Output Limit Exceeded", "Presentation Error", "None"}
var languageArr = []string{"None", "C", "C++", "Java"}
var encryptArr = []string{"None", "Public", "Private", "Password"}
var privilegeArr = []string{"None", "Primary User", "Teacher", "Admin"}

func ShowNext(num int) (next int) {
	next = num + 1
	return
}

func ShowStatus(status int) bool {
	return status == config.StatusAvailable
}

func ShowSim(sim int) bool {
	return sim != 0
}
func ShowTime(unixtime int64) string {
	return time.Unix(unixtime, 0).Local().Format("2006-01-02 15:04:05")
}

func ShowRatio(solve int, submit int) (ratio string) {
	if submit == 0 {
		ratio = "0.00%"
	} else {
		ratio = strconv.FormatFloat(float64(solve)/float64(submit)*100, 'f', 2, 64) + "%"
	}
	return
}

func ShowSpecial(num int) (special string) {
	special = specialArr[num]
	return
}

func ShowJudge(num int) (judge string) {
	judge = judgeArr[num]
	return
}

func ShowLanguage(num int) (language string) {
	language = languageArr[num]
	return
}

func ShowEncrypt(num int) (encrypt string) {
	encrypt = encryptArr[num]
	return
}

func NumAdd(a int, b int) (ret int) {
	ret = a + b
	return
}

func NumSub(a int, b int) (ret int) {
	ret = a - b
	return
}

func LargePU(privilege int) bool {
	return privilege > config.PrivilegePU
}

func ShowPrivilege(privilege int) string {
	return privilegeArr[privilege]
}

func SameID(ID1, ID2 string) bool {
	return ID1 == ID2
}
