package main

import (
	"GoOnlineJudge/controller"
	"GoOnlineJudge/controller/admin"
	"GoOnlineJudge/controller/contest"
	//"GoOnlineJudge/model"
	//"encoding/json"
	//"log"
	"net/http"
	"reflect"
	//"strconv"
	"strings"
)

// normal Page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		c := &controller.NewsController{}
		c.List(w, r)
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
	if args["problem"] != "" {
		c := &controller.ProblemController{}
		m := strings.Title(args["problem"])
		rv := getReflectValue(w, r)
		callMethod(c, m, rv)
	}
}

/*
func solutionHandler(w http.ResponseWriter, r *http.Request) {
	args := ParseURL(r.URL.String())

	log.Println(r.URL.String())
	log.Println(args["solution"])

	solutionModel := model.SolutionModel{}
	sid, err := strconv.Atoi(args["sid"])

	switch args["solution"] {
	case "detail":
		solution, err := solutionModel.Detail(sid)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		b, err := json.Marshal(&solution)
		if err != nil {
			http.Error(w, "json error", 500)
			return
		}

		w.WriteHeader(200)
		w.Write(b)

	case "delete":
		solution, err := solutionModel.Delete(sid)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		b, err := json.Marshal(&solution)
		if err != nil {
			http.Error(w, "json error", 500)
			return
		}

		w.WriteHeader(200)
		w.Write(b)

	}
}
*/

func statusHandler(w http.ResponseWriter, r *http.Request) {
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
	c.Index(w, r)
}

func contestlistHandler(w http.ResponseWriter, r *http.Request) {
	c := &controller.ContestController{}
	c.List(w, r)
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
	c.FAQ(w, r)
}

//Register User Page,need some privilege.

// Contest
func contestHandler(w http.ResponseWriter, r *http.Request) {
	c := &contest.ContestUserContorller{}
	c.Register(w, r)
}

// Admin
func adminHandler(w http.ResponseWriter, r *http.Request) {
	c := &admin.AdminUserController{}
	c.Register(w, r)
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

	for _, pair := range list {
		k_v := strings.Split(pair, "?")
		if len(k_v) <= 1 {
			args[k_v[0]] = ""
		} else {
			args[k_v[0]] = k_v[1]
		}
	}
	return
}
