package schedule

import (
	"testing"
)

func Test_Hdu(t *testing.T) {
	hdu := &HDUJudger{}
	hdu.Init()
	hdu.GetProblems()
}
