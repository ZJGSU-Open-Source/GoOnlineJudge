package handler

import (
    "github.com/zenazn/goji/web"
    "ojapi/config"
    "ojapi/middleware"
    "ojapi/model"

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

//@URL: /api/status/:sid @method: GET
func GetStatus(c web.C, w http.ResponseWriter, r *http.Request) {

    var (
        Sid  = c.URLParams["sid"]
        user = middleware.ToUser(c)
    )

    sid, err := strconv.Atoi(Sid)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    solutionModel := model.SolutionModel{}
    one, err := solutionModel.Detail(sid)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    // no login or Primary user
    pu := (user == nil || user.Privilege <= config.PrivilegePU)
    if one.Status == config.StatusReverse && pu {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    // neither solution poster nor admin
    pu = user == nil || (user.Privilege <= config.PrivilegePU && user.Uid != one.Uid)
    if one.Share == false && (pu) {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(one)
}
