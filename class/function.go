package class

import (
	"strconv"
)

var specialArr = []string{"Standard", "Special"}
var judgeArr = []string{"None", "Pengding", "Running & Judging", "Accept", "Compile Error", "Runtime Error", "Wrong Answer", "Time Limit Exceeded", "Memory Limit Exceeded", "Output Limit Exceeded"}
var languageArr = []string{"None", "C", "C++", "Java"}
var encryptArr = []string{"None", "Public", "Private", "Password"}

func ShowNext(num int) (next int) {
	next = num + 1
	return
}

func ShowStatus(num int) (status bool) {
	status = num%2 != 0
	return
}

func ShowExpire(str string, time string) (expire bool) {
	expire = time > str
	return
}

func ShowRatio(solve int, submit int) (ratio string) {
	ratio = strconv.FormatFloat(float64(solve)/float64(submit)*100, 'f', 2, 64) + "%"
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

func NumEqual(a int, b int) (ret bool) {
	return a == b
}

func NumAdd(a int, b int) (ret int) {
	ret = a + b
	return
}

func NumSub(a int, b int) (ret int) {
	ret = a - b
	return
}
