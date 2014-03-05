package class

import (
	"GoOnlineJudge/config"
	"strconv"
)

func ShowStatus(status int) bool {
	return status%2 != 0
}

func ShowRatio(solve int, submit int) string {
	return strconv.FormatFloat(float64(solve)/float64(submit)*100, 'f', 2, 64) + "%"
}

func ShowSpecial(special int) string {
	switch special {
	case config.SpecialST:
		return "Standard"
	case config.SpecialSP:
		return "Special"
	default:
		return "None"
	}
	return "None"
}
