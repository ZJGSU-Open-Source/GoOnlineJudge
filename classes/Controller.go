package classes

import (
	"log"
)

type Controller struct {
	Data map[string]interface{}
}

func (this *Controller) Init() {
	this.Data = make(map[string]interface{})
}

func (this *Controller) CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}
