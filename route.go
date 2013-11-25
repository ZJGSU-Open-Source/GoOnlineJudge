package main

import (
	"GoOnlineJudge/controllers"
	"net/http"
	"reflect"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	c := &controllers.HomeController{}
	m := r.Method
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

func problemListHandler(w http.ResponseWriter, r *http.Request) {
	c := &controllers.ProblemListController{}
	m := r.Method
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

func problemDetailHandler(w http.ResponseWriter, r *http.Request) {
	c := &controllers.ProblemDetailController{}
	m := r.Method
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		c := &controllers.NotFoundController{}
		m := r.Method
		rv := getReflectValue(w, r)
		callMethod(c, m, rv)
	}
}

func userLoginHandler(w http.ResponseWriter, r *http.Request) {
	c := &controllers.UserLoginController{}
	m := r.Method
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

func callMethod(c interface{}, m string, rv []reflect.Value) {
	rc := reflect.ValueOf(c)
	rm := rc.MethodByName(m)
	rm.Call(rv)
}

func getReflectValue(w http.ResponseWriter, r *http.Request) []reflect.Value {
	rw := reflect.ValueOf(w)
	rr := reflect.ValueOf(r)
	return []reflect.Value{rw, rr}
}
