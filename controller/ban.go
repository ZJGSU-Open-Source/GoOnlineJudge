package controller

var bans [9]string = [9]string{"392464930", "CCF", "ccf", "Ccf", "cCf", "CCf", "cCF", "CcF", "ccF"}

func Ban(str string) bool {
	l := len(str)
	flag := false
	for _, ban := range bans {
		idx := 0
		for _, char := range ban {
			flag = false
			for i := idx; i < l; i++ {
				if rune(str[i]) == char {
					idx = i
					flag = true
					break
				}
			}
		}
		if flag {
			return flag
		}
	}
	return flag
}
