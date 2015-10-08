package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// 问题控件
type ProblemController struct {
	class.Controller
} //@Controller

// 列出特定数量的问题?pid=<pid>&titile=<titile>&source=<source>&page=<page>
//@URL:/problems @method:GET
func (pc *ProblemController) List() {
	restweb.Logger.Debug(pc.R.RemoteAddr + " visit Problem List")

	url := "/problems?"
	// Search
	if v, ok := pc.Input["pid"]; ok { //按pid查找
		url += "pid=" + v[0] + "&"
		pc.Output["SearchPid"] = true
		pc.Output["SearchValue"] = v[0]
	} else if v, ok := pc.Input["title"]; ok { //按问题标题查找
		url += "title=" + v[0] + "&"
		pc.Output["SearchTitle"] = true
		pc.Output["SearchValue"] = v[0]
		for _, ep := range "+.?$|*^ " {
			v[0] = strings.Replace(v[0], string(ep), "\\"+string(ep), -1)
		}
	} else if v, ok := pc.Input["source"]; ok { //按问题来源查找
		url += "source=" + v[0] + "&"
		pc.Output["SearchSource"] = true
		pc.Output["SearchValue"] = v[0]
		for _, ep := range "+.?$|*^ " {
			v[0] = strings.Replace(v[0], string(ep), "\\"+string(ep), -1)
		}
	}

	pc.Output["URL"] = url

	page := 1
	var err error
	if v, ok := pc.Input["page"]; ok { //指定页码
		page, err = strconv.Atoi(v[0])
		if err != nil {
			pc.Error("args error", 400)
			return
		}
	}

	offset := (page - 1) * config.ProblemPerPage //偏移位置
	limit := config.ProblemPerPage               //每页问题数量

	restweb.Logger.Debug(offset, limit)

	var problemList []*model.Problem
	req, _ := apiClient.NewRequest("GET", fmt.Sprintf("%soffset=%d&limit=%d", url, offset, limit), pc.AccessToken, nil)
	_, err = apiClient.Do(req, &problemList)
	if err != nil {
		restweb.Logger.Debug(err)
		pc.Error(err.Error(), 500)
		return
	}

	count := len(problemList)
	restweb.Logger.Debug(count)

	pageData := pc.GetPage(page, count < config.ProblemPerPage)
	for k, v := range pageData {
		pc.Output[k] = v
	}

	// var ret *model.User
	// req, _ = apiClient.NewRequest("GET", "/profile", accessToken, nil)
	// apiClient.Do(req, &ret)

	// solutionModel := &model.SolutionModel{}
	// achieve, _ := solutionModel.Achieve(pc.Uid, config.ModuleP, config.ModuleP)
	// for _, p := range problemList {
	//     p.Flag = config.FlagNA
	//     for _, i := range achieve {
	//         if p.Pid == i {
	//             p.Flag = config.FLagAC
	//             break
	//         }
	//     }
	//     if p.Flag == config.FlagNA {
	//         args := make(map[string]string)
	//         args["pid"] = strconv.Itoa(p.Pid)
	//         args["module"] = strconv.Itoa(config.ModuleP)
	//         args["uid"] = pc.Uid
	//         l, _ := solutionModel.List(args)
	//         if len(l) > 0 {
	//             p.Flag = config.FLagER
	//         }
	//     }
	// }

	pc.Output["Problem"] = problemList
	pc.Output["Privilege"] = pc.Privilege
	pc.Output["Time"] = restweb.GetTime()
	pc.Output["Title"] = "Problem List"
	pc.Output["IsProblem"] = true
	pc.RenderTemplate("view/layout.tpl", "view/problem_list.tpl")
}

//列出某问题的详细信息
//@URL: /problems/(\d+) @method: GET
func (pc *ProblemController) Detail(pid string) {
	restweb.Logger.Debug("Problem Detail")

	var one *model.Problem
	req, _ := apiClient.NewRequest("GET", fmt.Sprintf("/problems/%s", pid), pc.AccessToken, nil)
	_, err := apiClient.Do(req, &one)
	if err != nil {
		pc.Err400("Problem "+pid, "No such problem")
		return
	}

	pc.Output["Detail"] = one
	pc.Output["Compiler_id"] = pc.GetSession("Compiler_id")
	pc.Output["Privilege"] = pc.Privilege
	pc.Output["Title"] = "Problem — " + pid
	pc.RenderTemplate("view/layout.tpl", "view/problem_detail.tpl")
}

//提交某一问题的solution
//@URL: /problems/(\d+) @method: POST
func (pc *ProblemController) Submit(pid string) {
	restweb.Logger.Debug("Problem Submit")

	out := struct {
		Code       string `json:"code"`
		CompilerID int    `json:"compiler_id"`
	}{}

	out.Code = pc.Input.Get("code")
	compilerID := pc.Input.Get("compiler_id")
	out.CompilerID, _ = strconv.Atoi(compilerID)
	pc.SetSession("Compiler_id", compilerID) //or set cookie?

	hint := make(map[string]string)

	var problem *model.Problem

	req, _ := apiClient.NewRequest("GET", fmt.Sprintf("/problems/%s", pid), pc.AccessToken, nil)
	_, err := apiClient.Do(req, &problem)

	errflag := true
	switch {
	case err != nil || problem == nil:
		hint["info"] = "No such problem."
	case len(out.Code) <= 0:
		hint["info"] = "Your source code is too short."
	default:
		errflag = false
	}

	if errflag {
		b, _ := json.Marshal(&hint)
		pc.W.WriteHeader(400)
		pc.W.Write(b)
		return
	}

	go func() {
		body, _ := pc.JsonReader(out)
		req, _ := apiClient.NewRequest("POST", fmt.Sprintf("/problems/%s/solutions", pid), pc.AccessToken, body)
		apiClient.Do(req, nil)
	}()

	pc.W.WriteHeader(201)

}
