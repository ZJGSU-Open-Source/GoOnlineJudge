package main

import (
	"net/http"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/problem/", problemHandler)

	http.HandleFunc("/userAjax/", userAjaxHandler)

	http.ListenAndServe(":8080", nil)
}
