package middleware

import (
	"net/http"
	"strconv"

	"ojapi/model"

	"github.com/zenazn/goji/web"
)

func SetContest(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var (
			Cid = c.URLParams["cid"]
		)

		if len(Cid) == 0 {
			h.ServeHTTP(w, r)
			return
		}

		cid, err := strconv.Atoi(Cid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		contestModel := &model.ContestModel{}
		contest, err := contestModel.Detail(cid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if contest != nil {
			ContestToC(c, contest)
		}

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
