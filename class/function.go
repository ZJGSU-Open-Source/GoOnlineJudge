package class

import (
	"GoOnlineJudge/config"
	"fmt"
	"strconv"
	"time"
)

var specialArr = []string{"Standard", "Special"}
var judgeArr = []string{"Pengding", "Running & Judging", "Compile Error", "Accepted", "Runtime Error",
	"Wrong Answer", "Time Limit Exceeded", "Memory Limit Exceeded", "Output Limit Exceeded", "Presentation Error", "System Error"}
var languageArr = []string{"None", "C", "C++", "Java"}
var encryptArr = []string{"None", "Public", "Private", "Password"}
var privilegeArr = []string{"None", "Primary User", "Teacher", "Admin"}

// ShowNext 返回num加1的值
func ShowNext(num int) (next int) {
	next = num + 1
	return
}

// ShowStatus 根据status确定状态是否可达的
func ShowStatus(status int) bool {
	return status == config.StatusAvailable
}

// ShowSim 是否显示相似度
func ShowSim(sim int) bool {
	return sim != 0
}

// ShowTime 将unixtime转换为当地时间
func ShowTime(unixtime int64) string {
	return time.Unix(unixtime, 0).Local().Format("2006-01-02 15:04:05")
}

// ShowRatio 显示solve/submit的比率
func ShowRatio(solve int, submit int) (ratio string) {
	if submit == 0 {
		ratio = "0.00%"
	} else {
		ratio = strconv.FormatFloat(float64(solve)/float64(submit)*100, 'f', 2, 64) + "%"
	}
	return
}

// ShowSpecial显示Judge程序是标准或是特判
func ShowSpecial(num int) (special string) {
	special = specialArr[num]
	return
}

// ShowJudge显示判题结果
func ShowJudge(num int) (judge string) {
	judge = judgeArr[num]
	return
}

// ShowLanguage 显示代码语言类型
func ShowLanguage(num int) (language string) {
	language = languageArr[num]
	return
}

// ShowEncrypt显示比赛的加密方式，公开，私有或者密码
func ShowEncrypt(num int) (encrypt string) {
	encrypt = encryptArr[num]
	return
}

// NumAdd 将两数相加
func NumAdd(a int, b int) (ret int) {
	ret = a + b
	return
}

// NumSub 两数相减a-b
func NumSub(a int, b int) (ret int) {
	ret = a - b
	return
}

// LargePU 判断privilege是否大于普通用户
func LargePU(privilege int) bool {
	return privilege > config.PrivilegePU
}

// ShowPrivilege 显示用户权限
func ShowPrivilege(privilege int) string {
	return privilegeArr[privilege]
}

// 判断两个ID是否相等
func SameID(ID1, ID2 string) bool {
	return ID1 == ID2
}

func ShowGapTime(gaptime int64) string {
	sec := gaptime % 60
	hour := gaptime / 3600
	minute := (gaptime - hour*3600) / 60
	return fmt.Sprintf("%d:%02d:%02d", hour, minute, sec)
}
