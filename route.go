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
	if r.URL.Path == "/" {
		c := &controllers.HomeController{}
		m := r.Method
		rv := getReflectValue(w, r)
		callMethod(c, m, rv)
	}
}

func problemListHandler() (w http.ResponseWriter, r *http.Request) {
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

func closeHandler(w http.ResponseWriter, r *http.Request) {
	c := &controllers.CloseController{}
	m := r.Method
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

// Ajax

func userAjaxHandler(w http.ResponseWriter, r *http.Request) {
	p := strings.Trim(r.URL.Path, "/")
	s := strings.Split(p, "/")
	if l := len(s); l >= 2 {
		c := &ajax.UserAjax{}
		m := strings.Title(s[1])
		rv := getReflectValue(w, r)
		callMethod(c, m, rv)
	}
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
