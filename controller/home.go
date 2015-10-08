package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/model"

	"restweb"
)

type HomeController struct {
	class.Controller
} //@Controller

//@URL: / @method: GET
func (hc *HomeController) Index() {
	restweb.Logger.Debug("Home")

	var newsList []*model.News
	req, _ := apiClient.NewRequest("GET", "/news", hc.AccessToken, nil)
	_, err := apiClient.Do(req, &newsList)
	if err != nil {
		restweb.Logger.Debug(err)
		hc.Error(err.Error(), 500)
		return
	}

	hc.Output["News"] = newsList
	hc.Output["Title"] = "Welcome to ZJGSU Online Judge"
	hc.Output["IsNews"] = true

	// ojModel := &model.OJModel{}
	// list, err := ojModel.List()
	// if err == nil {
	//     hc.Output["OJStatus"] = list
	// }

	hc.RenderTemplate("view/layout.tpl", "view/news_list.tpl")
}
