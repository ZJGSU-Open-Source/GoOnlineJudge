package main

import (
	"GoOnlineJudge/ajax"
	"GoOnlineJudge/controllers"
	"GoOnlineJudge/controllers/admin"
	"net/http"
	"reflect"
	"strings"
)

// Page

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		c := &controllers.HomeController{}
		m := r.Method
		rv := getReflectValue(w, r)
		callMethod(c, m, rv)
	}
}

func problemHandler(w http.ResponseWriter, r *http.Request) {
	p := strings.Trim(r.URL.Path, "/")
	s := strings.Split(p, "/")
	if l := len(s); l >= 2 {
		c := &controllers.ProblemController{}
		m := strings.Title(s[1])
		rv := getReflectValue(w, r)
		callMethod(c, m, rv)
	}
}

func closeHandler(w http.ResponseWriter, r *http.Request) {
	c := &controllers.CloseController{}
	m := r.Method
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

// Admin

func adminMenuHandler(w http.ResponseWriter, r *http.Request) {
	c := &admin.MenuController{}
	m := r.Method
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

func adminItemHandler(w http.ResponseWriter, r *http.Request) {
	p := strings.Trim(r.URL.Path, "/")
	s := strings.Split(p, "/")
	if l := len(s); l >= 2 {
		var c interface{}
		m := r.Method
		rv := getReflectValue(w, r)
		switch s[1] {
		case "notice":
			c = &admin.NoticeController{}
		case "news":
			c = &admin.NewsController{}
			m = strings.Title(s[2])
		case "problem":
			c = &admin.ProblemController{}
			m = strings.Title(s[2])
		}
		callMethod(c, m, rv)
	}
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

func newsAjaxHandler(w http.ResponseWriter, r *http.Request) {
	p := strings.Trim(r.URL.Path, "/")
	s := strings.Split(p, "/")
	if l := len(s); l >= 2 {
		c := &ajax.NewsAjax{}
		m := strings.Title(s[1])
		rv := getReflectValue(w, r)
		callMethod(c, m, rv)
	}
}

func problemAjaxHandler(w http.ResponseWriter, r *http.Request) {
	p := strings.Trim(r.URL.Path, "/")
	s := strings.Split(p, "/")
	if l := len(s); l >= 2 {
		c := &ajax.ProblemAjax{}
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
