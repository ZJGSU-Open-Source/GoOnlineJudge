package class

import (
	"strconv"
)

func ShowStatus(status int) bool {
	return status%2 != 0
}

func ShowRatio(solve int, submit int) string {
	return strconv.FormatFloat(float64(solve)/float64(submit)*100, 'f', 2, 64) + "%"
}
