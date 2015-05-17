package controller

import (
    "GoOnlineJudge/class"
    // "GoOnlineJudge/config"
    "GoOnlineJudge/model"
    "restweb"

    "encoding/json"
    "net/http"
    "strconv"
)

type StatusController struct {
    class.Controller
}   //@Controller

//@URL: /api/status @method: GET
func (sc *StatusController) List() {
    restweb.Logger.Debug("Status List")
    in := struct {
        Uid      string
        Pid      string
        Language int
        Judge    int
        Module   int
        Offset   int
        Limit    int
    }{}

    if err := json.NewDecoder(sc.R.Body).Decode(&in); err != nil {
        sc.Error(err.Error(), http.StatusBadRequest)
        return
    }
    qry := make(map[string]string)
    solutionModel := &model.SolutionModel{}

    list, err := solutionModel.List(qry)
    if err != nil {
        sc.Error(err.Error(), 500)
        return
    }

    sc.Output["Solution"] = list
    sc.RenderJson()
}

//@URL: /api/status/code @method: GET
func (sc *StatusController) Code() {
    restweb.Logger.Debug("Status Code")

    sid, err := strconv.Atoi(sc.Input.Get("sid"))
    if err != nil {
        http.Error(sc.W, "args error", 400)
        return
    }

    solutionModel := model.SolutionModel{}
    one, err := solutionModel.Detail(sid)
    if err != nil {
        sc.Error(err.Error(), 400)
        return
    }
    if one.Error != "" {
        one.Code = one.Code + "\n/*\n" + one.Error + "*/\n"
    }

    sc.Output["Solution"] = one
    sc.RenderJson()
}
