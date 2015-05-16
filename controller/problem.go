package controller

import (
    "GoOnlineJudge/class"
    "GoOnlineJudge/config"
    "GoOnlineJudge/model"

    "encoding/json"
    "net/http"
    "restweb"
    "strconv"
    "strings"
)

// 问题控件
type ProblemController struct {
    class.Controller
}   //@Controller

// 列出特定数量的问题?pid=<pid>&titile=<titile>&source=<source>&page=<page>
//@URL:/problems @method:GET
func (pc *ProblemController) List() {
    restweb.Logger.Debug(pc.R.RemoteAddr + "visit Problem List")

    qry := make(map[string]string)
    url := "/problems?"

    // Search
    if v, ok := pc.Input["pid"]; ok { //按pid查找
        qry["pid"] = v[0]
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
        qry["title"] = v[0]
    } else if v, ok := pc.Input["source"]; ok { //按问题来源查找
        url += "source=" + v[0] + "&"
        pc.Output["SearchSource"] = true
        pc.Output["SearchValue"] = v[0]
        for _, ep := range "+.?$|*^ " {
            v[0] = strings.Replace(v[0], string(ep), "\\"+string(ep), -1)
        }
        qry["source"] = v[0]
    }
    pc.Output["URL"] = url

    // Page
    qry["page"] = "1"
    if v, ok := pc.Input["page"]; ok { //指定页码
        qry["page"] = v[0]
    }

    if pc.Privilege <= config.PrivilegePU {
        qry["status"] = "2" //strconv.Itoa(config.StatusAvailable)
    }

    problemModel := model.ProblemModel{}
    count, err := problemModel.Count(qry)
    if err != nil {
        pc.Error(err.Error(), 500)
        return
    }

    restweb.Logger.Debug(count)
    var pageCount = (count-1)/config.ProblemPerPage + 1
    page, err := strconv.Atoi(qry["page"])
    if err != nil {
        pc.Error("args error", 400)
        return
    }
    if page > pageCount {
        pc.Error("args error", 400)
        return
    }

    qry["offset"] = strconv.Itoa((page - 1) * config.ProblemPerPage) //偏移位置
    qry["limit"] = strconv.Itoa(config.ProblemPerPage)               //每页问题数量
    pageData := pc.GetPage(page, pageCount)
    for k, v := range pageData {
        pc.Output[k] = v
    }

    problemList, err := problemModel.List(qry)
    if err != nil {
        pc.Error("post error", 500)
        return
    }
    restweb.Logger.Debug(len(problemList))

    pc.Output["Problem"] = problemList
    pc.Output["Privilege"] = pc.Privilege
    pc.Output["Time"] = restweb.GetTime()
    pc.Output["Title"] = "Problem List"
    pc.Output["IsProblem"] = true
    pc.RenderTemplate("view/layout.tpl", "view/problem_list.tpl")
}

//列出某问题的详细信息
//@URL: /problems/(\d+) @method: GET
func (pc *ProblemController) Detail(Pid string) {
    restweb.Logger.Debug("Problem Detail")

    pid, err := strconv.Atoi(Pid)
    if err != nil {
        pc.Error("args error", 400)
        return
    }

    problemModel := model.ProblemModel{}
    one, err := problemModel.Detail(pid)
    if err != nil {
        pc.Err400("Problem "+Pid, "No such problem")
        return
    }
    pc.Output["Detail"] = one

    if pc.Privilege <= config.PrivilegePU && one.Status == config.StatusReverse { // 如果问题状态为普通用户不可见
        pc.Err400("Problem "+Pid, "No such problem")
        return
    }

    pc.Output["Compiler_id"] = pc.GetSession("Compiler_id")
    pc.Output["Privilege"] = pc.Privilege
    pc.Output["Title"] = "Problem — " + Pid
    pc.RenderTemplate("view/layout.tpl", "view/problem_detail.tpl")
}

//提交某一问题的solution
//@URL: /problems/(\d+) @method: POST
func (pc *ProblemController) Submit(Pid string) {
    restweb.Logger.Debug("Problem Submit")

    pid, err := strconv.Atoi(Pid)
    if err != nil {
        pc.Error("args error", 400)
        return
    }

    var one model.Solution
    one.Pid = pid
    one.Uid = pc.Uid
    one.Module = config.ModuleP
    one.Mid = config.ModuleP

    problemModel := model.ProblemModel{}
    pro, err := problemModel.Detail(pid)
    if err != nil {
        pc.Error(err.Error(), 500)
        return
    }
    code := pc.Input.Get("code")

    one.Code = code
    one.Length = pc.GetCodeLen(len(pc.Input.Get("code")))
    one.Language, _ = strconv.Atoi(pc.Input.Get("compiler_id"))
    pc.SetSession("Compiler_id", pc.Input.Get("compiler_id")) //or set cookie?
    userModel := model.UserModel{}
    user, _ := userModel.Detail(pc.Uid)
    one.Share = user.ShareCode

    hint := make(map[string]string)
    errflag := true
    switch {
    case pro.Pid == 0:
        hint["info"] = "No such problem."
    case code == "":
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

    one.Status = config.StatusAvailable
    one.Judge = config.JudgePD

    solutionModel := model.SolutionModel{}
    sid, err := solutionModel.Insert(one)
    if err != nil {
        pc.Error(err.Error(), 500)
        return
    }

    pc.W.WriteHeader(201)
    go func() { //编译运行solution
        one := make(map[string]interface{})
        one["Sid"] = sid
        one["Pid"] = pro.RPid
        one["OJ"] = pro.ROJ
        one["Rejudge"] = false
        reader, _ := pc.JsonReader(&one)
        restweb.Logger.Debug(reader)
        _, err := http.Post(config.JudgeHost, "application/json", reader)
        if err != nil {
            restweb.Logger.Debug("sid[", sid, "] submit post error")
        }
    }()
}
