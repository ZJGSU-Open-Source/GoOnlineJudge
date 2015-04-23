package controller

import (
    "GoOnlineJudge/class"
    "GoOnlineJudge/model"
    "restweb"
)

type HomeController struct {
    class.Controller
}   //@Controller

//@URL: / @method: GET
func (hc *HomeController) Index() {
    restweb.Logger.Debug("Home")

    newsModel := model.NewsModel{}
    newsList, err := newsModel.List(-1, -1)
    if err != nil {
        hc.Error(err.Error(), 500)
        return
    }
    hc.Output["News"] = newsList
    hc.Output["Title"] = "Welcome to ZJGSU Online Judge"
    hc.Output["IsNews"] = true

    ojModel := &model.OJModel{}
    list, err := ojModel.List()
    if err == nil {
        for _, l := range list {
            restweb.Logger.Debug(*l)
        }
        hc.Output["OJStatus"] = list
    }

    hc.RenderTemplate("view/layout.tpl", "view/news_list.tpl")
}
