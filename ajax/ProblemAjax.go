package ajax

import (
	"GoOnlineJudge/config"
	"io/ioutil"
	"log"
	"net/http"
)

type ProblemAjax struct {
}

func (this *ProblemAjax) Insert(w http.ResponseWriter, r *http.Request) {
	log.Println("Problem Insert")
	w.Header().Set("content-type", "application/json")

	r.ParseForm()

	response, _ := http.PostForm(config.Host+"/problem/insert", r.PostForm)
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		w.Write(body)
	} else {
		log.Println("Problem Insert Response Error")
	}
}
