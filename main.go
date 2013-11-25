package main

import (
	"net/http"
)

func main() {
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.Handle("/img/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/home", homeHandler)

	http.HandleFunc("/problem/list", problemListHandler)
	http.HandleFunc("/problem/detail", problemDetailHandler)

	http.HandleFunc("/user/login", userLoginHandler)

	http.HandleFunc("/", notFoundHandler)

	http.ListenAndServe(":8080", nil)
}
