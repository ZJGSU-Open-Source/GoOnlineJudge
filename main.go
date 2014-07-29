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
	//http.HandleFunc("/help/", helpHandler)

	http.HandleFunc("/contestlist/", contestlistHandler)
	http.HandleFunc("/contest/", contestHandler)

	http.HandleFunc("/admin/", adminHandler)
	http.ListenAndServe(":8080", nil)
}
