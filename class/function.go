package class

import (
	"GoOnlineJudge/config"
	"strconv"
)

func ShowStatus(num int) (status bool) {
	status = num%2 != 0
	return
}

func ShowRatio(solve int, submit int) (ratio string) {
	ratio = strconv.FormatFloat(float64(solve)/float64(submit)*100, 'f', 2, 64) + "%"
	return
}

func ShowSpecial(num int) (special string) {
	switch special {
	case config.SpecialST:
		special = "Standard"
	case config.SpecialSP:
		special = "Special"
	default:
		special = "None"
	}
	return
}
