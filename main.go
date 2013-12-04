package main

import (
	"net/http"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/problem", problemListHandler)
	http.HandleFunc("/problem/", problemDetailHandler)
	http.HandleFunc("/close", closeHandler)

	http.HandleFunc("/userAjax/", userAjaxHandler)

	http.ListenAndServe(":8080", nil)
}
