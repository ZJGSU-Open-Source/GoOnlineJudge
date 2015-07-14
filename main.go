package main

import (
	// "log"
	"net/http"

	"github.com/zenazn/goji/web"

	// _ "GoOnlineJudge/schedule"

	"GoOnlineJudge/handler"
)

func main() {

	http.Handle("/api/", router())
	panic(http.ListenAndServe(":8080", nil))
}

func router() *web.Mux {
	mux := web.New()
	mux.Use(SetUser)

	mux.Get("/api/problems", handler.ListProblems)
	mux.Get("/api/problems/:pid", handler.GetProblem)

	return mux
}

func SetUser(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.FromC(*c)
		var user = session.GetUser(ctx, r)
		if user != nil && user.ID != 0 {
			UserToC(c, user)
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)

}
