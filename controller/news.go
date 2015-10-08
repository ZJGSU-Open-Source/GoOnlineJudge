package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"fmt"
)

//新闻控件

type NewsController struct {
	class.Controller
} //@Controller

//列出所有新闻
//@URL: /news @method: GET
func (nc *NewsController) List() {
	restweb.Logger.Debug("News List")

	var newsList []*model.News
	req, _ := apiClient.NewRequest("GET", "/news", nc.AccessToken, nil)
	_, err := apiClient.Do(req, &newsList)
	if err != nil {
		restweb.Logger.Debug(err)
		nc.Error(err.Error(), 500)
		return
	}

	nc.Output["News"] = newsList

	nc.Output["Title"] = "Welcome to ZJGSU Online Judge"
	nc.Output["IsNews"] = true
	nc.RenderTemplate("view/layout.tpl", "view/news_list.tpl")
}

//@URL: /news/(\d+) @method: GET
func (nc *NewsController) Detail(nid string) {

	var one *model.News
	req, _ := apiClient.NewRequest("GET", fmt.Sprintf("/news/%s", nid), nc.AccessToken, nil)
	_, err := apiClient.Do(req, &one)
	if err != nil {
		restweb.Logger.Debug(err)
		nc.Error(err.Error(), 500)
		return
	}

	nc.Output["Detail"] = one

	if one.Status == config.StatusReverse {
		nc.Err400("No such news", "No such news")
		return
	}

	nc.Output["Title"] = "News Detail"
	nc.RenderTemplate("view/layout.tpl", "view/news_detail.tpl")
}
