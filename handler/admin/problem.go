package admin

import (
	"ojapi/config"
	"ojapi/middleware"
	"ojapi/model"

	"github.com/zenazn/goji/web"

	"encoding/json"
	// "html/template"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//@URL: /admin/problems/ @method: POST
func PostProblem(c web.C, w http.ResponseWriter, r *http.Request) {

	one := problem(r)

	problemModel := model.ProblemModel{}
	pid, err := problemModel.Insert(one)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createfile(config.Datapath+strconv.Itoa(pid), "sample.in", one.In)
	createfile(config.Datapath+strconv.Itoa(pid), "sample.out", one.Out)

	w.WriteHeader(http.StatusCreated)
}

func createfile(path, filename string, context string) {

	err := os.Mkdir(path, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return
	}

	file, err := os.Create(path + "/" + filename)
	if err != nil {
		return
	}
	defer file.Close()

	var cr rune = 13
	crStr := string(cr)
	context = strings.Replace(context, "\r\n", "\n", -1)
	context = strings.Replace(context, crStr, "\n", -1)
	file.WriteString(context)
}

//@URL: /api/admin/problems/:pid/status/ @method: POST
func StatusProblem(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		Pid = c.URLParams["pid"]
		// user = middleware.ToUser(c)
	)

	pid, err := strconv.Atoi(Pid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	problemModel := model.ProblemModel{}
	one, err := problemModel.Detail(pid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var status int
	switch one.Status {
	case config.StatusAvailable:
		status = config.StatusReverse
	case config.StatusReverse:
		status = config.StatusAvailable
	}
	err = problemModel.Status(pid, status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//@URL: /api/admin/problems/:pid @method: DELETE
func DeleteProblem(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		Pid  = c.URLParams["pid"]
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	pid, err := strconv.Atoi(Pid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	problemModel := model.ProblemModel{}
	problemModel.Delete(pid)

	os.RemoveAll(config.Datapath + Pid) //delete test data
	w.WriteHeader(200)
}

//@URL: /admin/problems/(\d+)/ @method: PUT
func PutProblem(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		Pid  = c.URLParams["pid"]
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	pid, err := strconv.Atoi(Pid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	one := problem(r)

	problemModel := model.ProblemModel{}
	err = problemModel.Update(pid, one)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	createfile(config.Datapath+strconv.Itoa(pid), "sample.in", one.In)
	createfile(config.Datapath+strconv.Itoa(pid), "sample.out", one.Out)

	w.WriteHeader(http.StatusOK)
}

func problem(r *http.Request) (one model.Problem) {

	json.NewDecoder(r.Body).Decode(&one)

	// one.Title = pc.Input.Get("title")
	// time, err := strconv.Atoi(pc.Input.Get("time"))
	// if err != nil {
	//     pc.Error("The value 'Time' is neither too short nor too large", 400)
	//     return
	// }
	// one.Time = time
	// memory, err := strconv.Atoi(pc.Input.Get("memory"))
	// if err != nil {
	//     pc.Error("The value 'Memory' is neither too short nor too large", 400)
	//     return
	// }
	// one.Memory = memory
	// if _, ok := pc.Input["special"]; !ok {
	//     one.Special = 0
	// } else {
	//     one.Special = 1
	// }

	// in := pc.Input.Get("in")
	// out := pc.Input.Get("out")
	// one.Description = template.HTML(pc.Input.Get("description"))
	// one.Input = template.HTML(pc.Input.Get("input"))
	// one.Output = template.HTML(pc.Input.Get("output"))
	// one.In = in
	// one.Out = out
	// one.Source = pc.Input.Get("source")
	// one.Hint = template.HTML(pc.Input.Get("hint"))
	// one.Status = config.StatusReverse
	// one.ROJ = "ZJGSU"

	return one
}

//@URL: /api/admin/problems/importor/ @method: POST
func ImportProblem(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	fhs := r.MultipartForm.File["fps.xml"]
	file, err := fhs[0].Open()
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	contentStr := string(content)

	problem := model.Problem{}
	protype := reflect.TypeOf(problem)
	proValue := reflect.ValueOf(&problem).Elem()

	for i, lenth := 0, protype.NumField(); i < lenth; i++ {
		tag := protype.Field(i).Tag.Get("xml")

		if tag == "" {
			continue
		}
		matchStr := "<" + tag + `><!\[CDATA\[(?ms:(.*?))\]\]></` + tag + ">"
		tagRx := regexp.MustCompile(matchStr)
		tagString := tagRx.FindAllStringSubmatch(contentStr, -1)

		if len(tagString) > 0 {
			switch tag {
			case "time_limit", "memory_limit":
				limit, err := strconv.Atoi(tagString[0][1])
				if err != nil {
					limit = 1
				}
				proValue.Field(i).Set(reflect.ValueOf(limit))
			case "description", "input", "output":
				proValue.Field(i).SetString(tagString[0][1])
			default:
				proValue.Field(i).SetString(tagString[0][1])
			}
		}
	}
	problem.ROJ = "ZJGSU"
	proModel := model.ProblemModel{}
	pid, err := proModel.Insert(problem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 建立测试数据文件
	createfile(config.Datapath+strconv.Itoa(pid), "sample.in", problem.In)
	createfile(config.Datapath+strconv.Itoa(pid), "sample.out", problem.Out)

	flag, flagJ := true, -1
	for _, tag := range []string{"test_input", "test_output"} {
		// restweb.Logger.Debug(tag)
		matchStr := "<" + tag + `><!\[CDATA\[(?ms:(.*?))\]\]></` + tag + ">"
		tagRx := regexp.MustCompile(matchStr)
		tagString := tagRx.FindAllStringSubmatch(contentStr, -1)
		// restweb.Logger.Debug(tagString)
		if flag {
			flag = false
			caselenth := 0
			for matchLen, j := len(tagString), 0; j < matchLen; j++ {
				if len(tagString[j][1]) > caselenth {
					caselenth = len(tagString[j][1])
					flagJ = j
				}
			}
		}
		if flagJ >= 0 && flagJ < len(tagString) {
			// restweb.Logger.Debug(tagString[flagJ][1])
			filename := strings.Replace(tag, "_", ".", 1)
			filename = strings.Replace(filename, "put", "", -1)
			createfile(config.Datapath+strconv.Itoa(pid), filename, tagString[flagJ][1])
		}
	}

	w.WriteHeader(http.StatusCreated)
}
