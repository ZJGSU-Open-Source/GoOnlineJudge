package handler

import (

    // "GoOnlineJudge/config"
    "GoOnlineJudge/model"
    "github.com/zenazn/goji/web"
    "restweb"

    "encoding/json"
    "net/http"
    "strconv"
)

//@URL: /api/status @method: GET
func StatusList(c web.C, w http.ResponseWriter, r *http.Request) {

    r.ParseForm()

    qry := make(map[string]string)

    if v := r.FormValue("pid"); len(v) > 0 {
        qry["pid"] = v
    }
    if v := r.FormValue("uid"); len(v) > 0 {
        qry["uid"] = v
    }
    if v := r.FormValue("language"); len(v) > 0 {
        qry["language"] = v
    }
    if v := r.FormValue("offset"); len(v) > 0 {
        qry["offset"] = v
    }
    if v := r.FormValue("limit"); len(v) > 0 {
        qry["limit"] = v
    }
    if v := r.FormValue("judge"); len(v) > 0 {
        qry["judge"] = v
    }
    if v := r.FormValue("module"); len(v) > 0 {
        qry["module"] = v
    }

    solutionModel := &model.SolutionModel{}

    list, err := solutionModel.List(qry)
    if err != nil {
        w.WriteHeader(500)
        return
    }

    json.NewEncoder(w).Encode(list)
}

//@URL: /api/status/:sid/code @method: GET
func GetCode(c web.C, w http.ResponseWriter, r *http.Request) {
    restweb.Logger.Debug("Status Code")
    var (
        Sid = c.URLParams["sid"]
    )

    sid, err := strconv.Atoi(Sid)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    solutionModel := model.SolutionModel{}
    one, err := solutionModel.Detail(sid)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    if one.Error != "" {
        one.Code = one.Code + "\n/*\n" + one.Error + "*/\n"
    }

    json.NewEncoder(w).Encode(one)
}
