package class

import (
	"net/http"
)

type Controller struct {
	Data map[string]interface{}
}

func (this *Controller) Init(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	this.Data = make(map[string]interface{})
}
