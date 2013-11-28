package main

import (
	"GoOnlineJudge/ajax"
	"GoOnlineJudge/controllers"
	"net/http"
	"reflect"
	"strings"
)

// Controller

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

// Ajax

func userAjaxHandler(w http.ResponseWriter, r *http.Request) {
	c := &ajax.UserAjax{}
	p := strings.Trim(r.URL.Path, "/")
	s := strings.Split(p, "/")
	m := strings.Title(s[len(s)-1])
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

// Common

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
