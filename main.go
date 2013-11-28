package main

import (
	"net/http"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", notFoundHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/problem", problemListHandler)

	http.HandleFunc("/user/", userAjaxHandler)

	http.ListenAndServe(":8080", nil)
}
