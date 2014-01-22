package ajax

import (
	"GoOnlineJudge/config"
	"io/ioutil"
	"log"
	"net/http"
)

type NewsAjax struct {
}

func (this *NewsAjax) Insert(w http.ResponseWriter, r *http.Request) {
	log.Println("News Insert")
	w.Header().Set("content-type", "application/json")

	r.ParseForm()

	response, _ := http.PostForm(config.Host+"/news/insert", r.PostForm)
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		w.Write(body)
	} else {
		log.Println("News Insert Response Error")
	}
}
