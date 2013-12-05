package classes

import (
	"log"
	"net/http"
)

type Controller struct {
	Data map[string]interface{}
}

func (this *Controller) Init(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	this.Data = make(map[string]interface{})
}

func (this *Controller) CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}
