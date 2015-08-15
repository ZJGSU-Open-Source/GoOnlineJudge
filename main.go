package main

import (
    "log"
    "net/http"

    "github.com/zenazn/goji/web"

    // _ "GoOnlineJudge/schedule"

    "GoOnlineJudge/handler"
    "GoOnlineJudge/middleware"
)

func main() {

    http.Handle("/api/", router())
    log.Println("Start server on port :8080...")
    panic(http.ListenAndServe(":8080", nil))
}

func router() *web.Mux {
    mux := web.New()
    // mux.Use(SetUser)

    mux.Get("/api/problems", handler.ListProblems)
    mux.Get("/api/problems/:pid", handler.GetProblem)
    mux.Post("/api/problems/:pid/solutions", handler.PostSolution)

    mux.Get("/api/ranklist", handler.Ranklist)

    mux.Get("/api/status", handler.StatusList)
    mux.Get("/api/status/:sid/code", handler.GetCode)

    mux.Get("/api/news", handler.ListNews)
    mux.Get("/api/news/:nid", handler.GetNews)

    mux.Get("/api/contests", handler.ContestList)

    mux.Get("/api/profile", handler.GetProfile)
    mux.Post("/api/users", handler.PostUser)
    mux.Get("/api/users/:user", handler.GetUser)

    mux.Post("/api/sess", handler.PostSess)
    mux.Delete("/api/sess", handler.DeleteSess)
    mux.Use(middleware.SetUser)

    return mux
}
