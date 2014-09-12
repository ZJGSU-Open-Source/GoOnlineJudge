package main

import (
	"GoOnlineJudge/controller"
	"GoOnlineJudge/controller/admin"
	"GoOnlineJudge/controller/contest"
	"net/http"
)

// normal Page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		c := &controller.HomeController{}
		c.Route(w, r)
	}
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	c := &controller.NewsController{}
	c.Route(w, r)
}

func problemHandler(w http.ResponseWriter, r *http.Request) {
	c := &controller.ProblemController{}
	c.Route(w, r)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	c := &controller.StatusController{}
	c.Route(w, r)
}

func ranklistHandler(w http.ResponseWriter, r *http.Request) {
	c := &controller.RanklistController{}
	c.Route(w, r)
}

func contestlistHandler(w http.ResponseWriter, r *http.Request) {
	c := &controller.ContestController{}
	c.List(w, r)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	c := &controller.UserController{}
	c.Route(w, r)
}

//FAQ
func FAQHandler(w http.ResponseWriter, r *http.Request) {
	c := &controller.FAQController{}
	c.FAQ(w, r)
}

//Register User Page,need some privilege.

// Contest
func contestHandler(w http.ResponseWriter, r *http.Request) {
	c := &contest.ContestUserContorller{}
	c.Register(w, r)
}

// Admin
func adminHandler(w http.ResponseWriter, r *http.Request) {
	c := &admin.AdminUserController{}
	c.Register(w, r)
}
