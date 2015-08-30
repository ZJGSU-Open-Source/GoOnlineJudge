package main

import (
	"log"
	"net/http"

	"github.com/zenazn/goji/web"

	// _ "GoOnlineJudge/schedule"

	"GoOnlineJudge/handler"
	"GoOnlineJudge/handler/admin"
	"GoOnlineJudge/middleware"
)

func main() {

	http.Handle("/api/", router())
	log.Println("Start server on port :8080...")
	panic(http.ListenAndServe(":8080", nil))
}

func router() *web.Mux {
	mux := web.New()

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

	mux.Handle("/api/admin/*", rootRouter())
	mux.Use(middleware.SetUser)

	return mux
}

func rootRouter() *web.Mux {
	root := web.New()
	root.Post("/api/admin/images", admin.PostImage)

	contest := web.New()
	contest.Post("/api/admin/contest", admin.PostContest)
	contest.Patch("/api/admin/contests/:cid/status", admin.StatusContest)
	contest.Delete("/api/admin/contests/:cid", admin.DeleteContest)
	contest.Put("/api/admin/contests/:cid", admin.PutContest)
	contest.Use(middleware.SetContest)
	root.Handle("/api/admin/contest", contest)
	root.Handle("/api/admin/contest/:cid", contest)
	root.Handle("/api/admin/contest/:cid/*", contest)

	news := web.New()
	news.Use(middleware.SetNews)
	news.Post("/api/admin/news", admin.PostNews)
	news.Patch("/api/admin/news/:nid/status", admin.NewsStatus)
	news.Delete("/api/admin/news/:nid", admin.DeleteNews)
	news.Put("/api/admin/news/:nid", admin.PutNews)
	root.Handle("/api/admin/news", news)
	root.Handle("/api/admin/news/:nid", news)
	root.Handle("/api/admin/news/:nid/*", news)

	problem := web.New()
	problem.Post("/api/admin/problems", admin.PostProblem)
	problem.Patch("/api/admin/problems/:pid/status", admin.StatusProblem)
	problem.Delete("/api/admin/problems/:pid", admin.DeleteProblem)
	problem.Put("/api/admin/problems/:pid", admin.PutProblem)
	problem.Use(middleware.SetProblem)
	root.Handle("/api/admin/problems", problem)
	root.Handle("/api/admin/problems/:pid", problem)
	root.Handle("/api/admin/problems/:pid/*", problem)

	root.Use(middleware.RequireUserAdmin)

	return root
}
