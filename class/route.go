package class

import (
	"net/http"
	"reflect"
	"strings"
)

type Router interface {
	Route(w http.ResponseWriter, r *http.Request)
}

func CallMethod(c interface{}, m string, rv []reflect.Value) {
	rc := reflect.ValueOf(c)
	rm := rc.MethodByName(m)
	rm.Call(rv)
}

func GetReflectValue(w http.ResponseWriter, r *http.Request) (rv []reflect.Value) {
	rw := reflect.ValueOf(w)
	rr := reflect.ValueOf(r)
	rv = []reflect.Value{rw, rr}
	return
}

var RouterMap = map[string]Router{}

//添加路由
func AddRouter(pattern string, router Router) {
	RouterMap[pattern] = router
}

var FileMap = map[string]http.Handler{}

//添加静态文件路由
func AddFile(pattern string, fileHandler http.Handler) {
	FileMap[pattern] = fileHandler
}

type Server struct {
}

//路由，先处理静态文件，后处理控件，按照最大匹配原则匹配路由
func (this *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path + "/"

	filemaxlenth := 0
	var realFileHandler http.Handler
	for pattern, fileHandler := range FileMap {
		if len(pattern) > filemaxlenth && strings.HasPrefix(path, pattern) {
			filemaxlenth = len(pattern)
			realFileHandler = fileHandler
		}
	}

	maxlenth := 0
	var realRouter Router
	for pattern, router := range RouterMap {
		if len(pattern) > maxlenth && strings.HasPrefix(path, pattern) {
			maxlenth = len(pattern)
			realRouter = router
		}
	}

	if filemaxlenth > maxlenth {
		realFileHandler.ServeHTTP(w, r)
	} else if maxlenth > 0 {
		realRouter.Route(w, r)
	} else {
		http.Error(w, "no such page", 404)
	}
}

// 运行服务器
func Run() error {
	return http.ListenAndServe(":8080", &Server{})
}
