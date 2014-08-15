package main

import (
	"GoOnlineJudge/controller"
	"GoOnlineJudge/controller/admin"
	"GoOnlineJudge/controller/contest"
	"log"
	"net/http"
	"reflect"
	"strings"
)

// normal Page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		c := &controller.NewsController{}
		m := "List"
		rv := getReflectValue(w, r)
		callMethod(c, m, rv)
	}
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	args := ParseURL(r.URL.String())
	if args["news"] != "" {
		c := &controller.NewsController{}
		m := strings.Title(args["news"])
		rv := getReflectValue(w, r)
		callMethod(c, m, rv)
	}
}

func problemHandler(w http.ResponseWriter, r *http.Request) {
	args := ParseURL(r.URL.String())
	log.Println(args)
	if args["problem"] != "" {
		log.Println("problem")
		c := &controller.ProblemController{}
		m := strings.Title(args["problem"])
		rv := getReflectValue(w, r)
		callMethod(c, m, rv)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.String())
	args := ParseURL(r.URL.String())
	if args["status"] != "" {
		c := &controller.StatusController{}
		m := strings.Title(args["status"])
		rv := getReflectValue(w, r)
		callMethod(c, m, rv)
	}
}

func ranklistHandler(w http.ResponseWriter, r *http.Request) {
	c := &controller.RanklistController{}
	m := "Index"
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

func contestlistHandler(w http.ResponseWriter, r *http.Request) {
	c := &controller.ContestController{}
	m := "List"
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	args := ParseURL(r.URL.String())
	if args["user"] != "" {
		c := &controller.UserController{}
		m := strings.Title(args["user"])
		rv := getReflectValue(w, r)
		callMethod(c, m, rv)
	}
}

//FAQ
func FAQHandler(w http.ResponseWriter, r *http.Request) {
	c := &controller.FAQController{}
	m := "FAQ"
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

//Register User Page,need some privilege.

// Contest
func contestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("conetest")
	c := &contest.ContestUserContorller{}
	m := "Register"
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

// Admin
func adminHandler(w http.ResponseWriter, r *http.Request) {
	c := &admin.AdminUserController{}
	m := "Register"
	rv := getReflectValue(w, r)
	callMethod(c, m, rv)
}

// Common
func callMethod(c interface{}, m string, rv []reflect.Value) {
	rc := reflect.ValueOf(c)
	rm := rc.MethodByName(m)
	rm.Call(rv)
}

func getReflectValue(w http.ResponseWriter, r *http.Request) (rv []reflect.Value) {
	rw := reflect.ValueOf(w)
	rr := reflect.ValueOf(r)
	rv = []reflect.Value{rw, rr}
	return
}

func ParseURL(url string) (args map[string]string) {
	args = make(map[string]string)
	path := strings.Trim(url, "/")
	list := strings.Split(path, "/")

	log.Println(url)
	log.Println(path)
	log.Println(list)
	for _, pair := range list {
		k_v := strings.Split(pair, "?")
		log.Println(k_v)
		if len(k_v) <= 1 {
			args[k_v[0]] = ""
		} else {
			args[k_v[0]] = k_v[1]
		}
	}
	return
}
