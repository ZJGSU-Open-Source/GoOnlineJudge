package middleware

import (
	"net/http"
	"strconv"

	"ojapi/model"

	"github.com/zenazn/goji/web"
)

func SetNews(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var (
			Nid = c.URLParams["nid"]
		)

		if len(Nid) == 0 {
			h.ServeHTTP(w, r)
			return
		}

		nid, err := strconv.Atoi(Nid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newsModel := &model.NewsModel{}
		news, err := newsModel.Detail(nid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if news != nil {
			NewsToC(c, news)
		}

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
