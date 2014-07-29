package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type help struct {
}

type HelpController struct {
	class.Controller
}

func (this *HelpController) Help(w http.ResponseWriter, r *http.Request) {
	log.Println("Help Page")
	this.Init(w, r)

	response, err := http.Post(config.PostHost+"/help", "application/json", nil)
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}
	defer response.Body.Close()

}
