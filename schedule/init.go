package schedule

import (
	"log"
	"os"
)

func init() {
	hduLogfile, err := os.Create("log/hdu.log")
	if err != nil {
		log.Println(err)
		return
	}
	hdulogger = log.New(hduLogfile, "[Hdu]", log.Ldate|log.Ltime)
	hdu := &HDUJudger{}
	hdu.Init()
	go hdu.GetProblems()

	PKULogfile, err := os.Create("log/pku.log")
	if err != nil {
		log.Println(err)
		return
	}
	PKUlogger = log.New(PKULogfile, "[PKU]", log.Ldate|log.Ltime)
	PKU := &PKUJudger{}
	PKU.Init()
	go PKU.GetProblems()
}
