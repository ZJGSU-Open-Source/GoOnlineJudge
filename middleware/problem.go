package middleware

import (
	"net/http"
	"strconv"

	"ojapi/model"

	"github.com/zenazn/goji/web"
)

func SetProblem(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var (
			Pid = c.URLParams["pid"]
		)
		if len(Pid) == 0 {
			h.ServeHTTP(w, r)
			return
		}

		pid, err := strconv.Atoi(Pid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		problemModel := &model.ProblemModel{}
		problem, err := problemModel.Detail(pid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if problem != nil {
			ProblemToC(c, problem)
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
