package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"fmt"
	"strconv"
)

type StatusController struct {
	class.Controller
} //@Controller

//@URL: /status @method: GET
func (sc *StatusController) List() {
	restweb.Logger.Debug("Status List")

	searchUrl := "/status?"

	// Search
	if v, ok := sc.Input["uid"]; ok {
		searchUrl += "uid=" + v[0] + "&"
		sc.Output["SearchUid"] = v[0]
	}
	if v, ok := sc.Input["pid"]; ok {
		searchUrl += "pid=" + v[0] + "&"
		sc.Output["SearchPid"] = v[0]
	}
	if v, ok := sc.Input["judge"]; ok {
		searchUrl += "judge=" + v[0] + "&"
		sc.Output["SearchJudge"+v[0]] = v[0]
	}
	if v, ok := sc.Input["language"]; ok {
		searchUrl += "language=" + v[0] + "&"
		sc.Output["SearchLanguage"+v[0]] = v[0]
	}
	sc.Output["URL"] = searchUrl

	// Page
	page := 1
	var err error

	if v, ok := sc.Input["page"]; ok {
		page, err = strconv.Atoi(v[0])
		if err != nil {
			sc.Error("args error", 400)
			return
		}
	}

	offset := (page - 1) * config.SolutionPerPage
	limit := config.SolutionPerPage

	var solutions []*model.Solution
	req, _ := apiClient.NewRequest("GET", fmt.Sprintf("%soffset=%d&limit=%d", searchUrl, offset, limit), sc.AccessToken, nil)
	_, err = apiClient.Do(req, &solutions)
	if err != nil {
		sc.Error(err.Error(), 500)
		return
	}

	count := len(solutions)

	pageData := sc.GetPage(page, count < limit)
	for k, v := range pageData {
		sc.Output[k] = v
	}

	sc.Output["Solution"] = solutions
	sc.Output["Title"] = "Status List"
	sc.Output["IsStatus"] = true
	sc.Output["Privilege"] = sc.Privilege
	sc.Output["Uid"] = sc.Uid

	sc.RenderTemplate("view/layout.tpl", "view/status_list.tpl")
}

//@URL: /status/code @method: GET
func (sc *StatusController) Code() {
	restweb.Logger.Debug("Status Code")

	sid := sc.Input.Get("sid")

	var one *model.Solution
	req, _ := apiClient.NewRequest("GET", fmt.Sprintf("/status/%s", sid), sc.AccessToken, nil)
	_, err := apiClient.Do(req, &one)
	if err != nil {
		sc.Err400("Warning", "You can't see it!")
		return
	}

	if one.Error != "" {
		one.Code = one.Code + "\n/*\n" + one.Error + "*/\n"
	}

	sc.Output["Solution"] = one
	sc.Output["Title"] = "View Code"
	sc.Output["IsCode"] = true
	sc.RenderTemplate("view/layout.tpl", "view/status_code.tpl")
}
