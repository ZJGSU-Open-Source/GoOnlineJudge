package main

import (
	"net/http"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/news/", newsHandler)
	http.HandleFunc("/problem/", problemHandler)
	http.HandleFunc("/status/", statusHandler)
	http.HandleFunc("/ranklist/", ranklistHandler)
	http.HandleFunc("/user/", userHandler)

	http.HandleFunc("/contest/", contestHandler)
	http.HandleFunc("/contest/problem/", contestProblemHandler)
	http.HandleFunc("/contest/status/", contestStatusHandler)
	// http.HandleFunc("/contest/ranklist/", contestRanklistHandler)

	http.HandleFunc("/admin/", adminHomeHandler)
	http.HandleFunc("/admin/news/", adminNewsHandler)
	http.HandleFunc("/admin/problem/", adminProblemHandler)

	http.ListenAndServe(":8080", nil)
}
