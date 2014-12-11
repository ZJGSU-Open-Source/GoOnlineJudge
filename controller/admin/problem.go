package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"

	"restweb"

	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type AdminProblem struct {
	class.Controller
}

// func (pc *AdminProblem) Detail() {
// 	restweb.Logger.Debug("Admin Problem Detail")

// 	pid, err := strconv.Atoi(r.URL.Query().Get("pid"))
// 	if err != nil {
// 		http.Error(w, "args error", 400)
// 		return
// 	}

// 	problemModel := model.ProblemModel{}
// 	one, err := problemModel.Detail(pid)
// 	if err != nil {
// 		pc.Error(err.Error(), 400)
// 		return
// 	}
// 	pc.Output["Detail"] = one
// 	pc.Output["Title"] = "Admin - Problem Detail"
// 	pc.Output["IsProblem"] = true
// 	pc.Output["IsList"] = false

// 	pc.RenderTemplate("view/admin/layout.tpl", "view/problem_detail.tpl")
// }

func (pc *AdminProblem) List() {
	restweb.Logger.Debug("Admin Problem List")

	problemModel := model.ProblemModel{}
	qry := make(map[string]string)

	qry["page"] = "1"
	if v, ok := pc.Input["page"]; ok { //指定页码
		qry["page"] = v[0]
	}
	count, err := problemModel.Count(qry)
	if err != nil {
		pc.Error(err.Error(), 500)
		return
	}

	restweb.Logger.Debug(count)
	var pageCount = (count-1)/config.ProblemPerPage + 1
	page, err := strconv.Atoi(qry["page"])
	if err != nil {
		pc.Error("args error", 400)
		return
	}
	if page > pageCount {
		pc.Error("args error", 400)
		return
	}

	qry["offset"] = strconv.Itoa((page - 1) * config.ProblemPerPage) //偏移位置
	qry["limit"] = strconv.Itoa(config.ProblemPerPage)               //每页问题数量
	pageData := pc.GetPage(page, pageCount)
	for k, v := range pageData {
		pc.Output[k] = v
	}

	proList, err := problemModel.List(qry)
	if err != nil {
		pc.Error(err.Error(), 400)
		return
	}

	pc.Output["Problem"] = proList
	pc.Output["Title"] = "Admin - Problem List"
	pc.Output["IsProblem"] = true
	pc.Output["IsList"] = true

	pc.RenderTemplate("view/admin/layout.tpl", "view/admin/problem_list.tpl")
}

func (pc *AdminProblem) Add() {
	restweb.Logger.Debug("Admin Problem Add")

	pc.Output["Title"] = "Admin - Problem Add"
	pc.Output["IsProblem"] = true
	pc.Output["IsAdd"] = true
	pc.Output["IsEdit"] = true

	pc.RenderTemplate("view/admin/layout.tpl", "view/admin/problem_add.tpl")
}

func (pc *AdminProblem) Insert() {
	restweb.Logger.Debug("Admin Problem Insert")

	one := pc.problem()

	problemModel := model.ProblemModel{}
	pid, err := problemModel.Insert(one)
	if err != nil {
		pc.Error(err.Error(), 500)
		return
	}

	createfile(config.Datapath+strconv.Itoa(pid), "sample.in", one.In)
	createfile(config.Datapath+strconv.Itoa(pid), "sample.out", one.Out)

	pc.Redirect("/problems/"+strconv.Itoa(pid), http.StatusFound)
}

func createfile(path, filename string, context string) {

	err := os.Mkdir(path, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		restweb.Logger.Debug("create dir error")
		return
	}

	file, err := os.Create(path + "/" + filename)
	if err != nil {
		restweb.Logger.Debug(err)
		return
	}
	defer file.Close()

	var cr rune = 13
	crStr := string(cr)
	context = strings.Replace(context, "\r\n", "\n", -1)
	context = strings.Replace(context, crStr, "\n", -1)
	file.WriteString(context)
}

func (pc *AdminProblem) Status(Pid string) {
	restweb.Logger.Debug("Admin Problem Status")

	pid, err := strconv.Atoi(Pid)
	if err != nil {
		pc.Error("args error", 400)
		return
	}

	problemModel := model.ProblemModel{}
	one, err := problemModel.Detail(pid)
	if err != nil {
		pc.Error(err.Error(), 400)
		return
	}
	pc.Output["Detail"] = one
	var status int
	switch one.Status {
	case config.StatusAvailable:
		status = config.StatusReverse
	case config.StatusReverse:
		status = config.StatusAvailable
	}
	err = problemModel.Status(pid, status)
	if err != nil {
		pc.Error(err.Error(), 500)
		return
	}

	pc.Redirect("/admin/problems", http.StatusFound)
}

func (pc *AdminProblem) Delete(Pid string) {
	restweb.Logger.Debug("Admin Problem Delete")

	if pc.Privilege != config.PrivilegeAD {
		pc.Err400("Warning", "Error Privilege to Delete problem")
		return
	}

	pid, err := strconv.Atoi(Pid)
	if err != nil {
		pc.Error("args error", 400)
		return
	}

	problemModel := model.ProblemModel{}
	problemModel.Delete(pid)

	os.RemoveAll(config.Datapath + Pid) //delete test data
	pc.W.WriteHeader(200)
}

func (pc *AdminProblem) Edit(Pid string) {
	restweb.Logger.Debug("Admin Problem Edit")

	if pc.Privilege != config.PrivilegeAD {
		pc.Err400("Warning", "Error Privilege to Edit problem")
		return
	}

	pid, err := strconv.Atoi(Pid)
	if err != nil {
		pc.Error("args error", 400)
		return
	}

	problemModel := model.ProblemModel{}
	one, err := problemModel.Detail(pid)
	if err != nil {
		pc.Error(err.Error(), 500)
		return
	}

	pc.Output["Detail"] = one
	pc.Output["Title"] = "Admin - Problem Edit"
	pc.Output["IsProblem"] = true
	pc.Output["IsList"] = false
	pc.Output["IsEdit"] = true

	pc.RenderTemplate("view/admin/layout.tpl", "view/admin/problem_edit.tpl")
}

func (pc *AdminProblem) Update(Pid string) {
	restweb.Logger.Debug("Admin Problem Update")

	if pc.Privilege != config.PrivilegeAD {
		pc.Err400("Warning", "Error Privilege to Update problem")
		return
	}

	pid, err := strconv.Atoi(Pid)
	if err != nil {
		pc.Error("args error", 400)
		return
	}
	one := pc.problem()

	problemModel := model.ProblemModel{}
	err = problemModel.Update(pid, one)
	if err != nil {
		pc.Error(err.Error(), 500)
		return
	}
	createfile(config.Datapath+strconv.Itoa(pid), "sample.in", one.In)
	createfile(config.Datapath+strconv.Itoa(pid), "sample.out", one.Out)

	pc.Redirect("/problems/"+strconv.Itoa(pid), http.StatusFound)
}

func (pc *AdminProblem) problem() (one model.Problem) {
	one.Title = pc.Input.Get("title")
	time, err := strconv.Atoi(pc.Input.Get("time"))
	if err != nil {
		pc.Error("The value 'Time' is neither too short nor too large", 400)
		return
	}
	one.Time = time
	memory, err := strconv.Atoi(pc.Input.Get("memory"))
	if err != nil {
		pc.Error("The value 'Memory' is neither too short nor too large", 400)
		return
	}
	one.Memory = memory
	if _, ok := pc.Input["special"]; !ok {
		one.Special = 0
	} else {
		one.Special = 1
	}

	in := pc.Input.Get("in")
	out := pc.Input.Get("out")
	one.Description = template.HTML(pc.Input.Get("description"))
	one.Input = template.HTML(pc.Input.Get("input"))
	one.Output = template.HTML(pc.Input.Get("output"))
	one.In = in
	one.Out = out
	one.Source = pc.Input.Get("source")
	one.Hint = pc.Input.Get("hint")

	return one
}

func (pc *AdminProblem) ImportPage() {
	pc.Output["Title"] = "Problem Import"
	pc.Output["IsProblem"] = true
	pc.Output["IsImport"] = true
	pc.RenderTemplate("view/admin/layout.tpl", "view/admin/problem_import.tpl")
}

func (pc *AdminProblem) Import() {
	pc.R.ParseMultipartForm(32 << 20)
	fhs := pc.R.MultipartForm.File["fps.xml"]
	file, err := fhs[0].Open()
	if err != nil {
		restweb.Logger.Debug(err)
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		restweb.Logger.Debug(err)
		return
	}
	contentStr := string(content)

	problem := model.Problem{}
	protype := reflect.TypeOf(problem)
	proValue := reflect.ValueOf(&problem).Elem()
	restweb.Logger.Debug(protype.NumField())
	for i, lenth := 0, protype.NumField(); i < lenth; i++ {
		tag := protype.Field(i).Tag.Get("xml")
		restweb.Logger.Debug(i, tag)
		if tag == "" {
			continue
		}
		matchStr := "<" + tag + `><!\[CDATA\[(?ms:(.*?))\]\]></` + tag + ">"
		tagRx := regexp.MustCompile(matchStr)
		tagString := tagRx.FindAllStringSubmatch(contentStr, -1)
		restweb.Logger.Debug(tag)
		if len(tagString) > 0 {
			switch tag {
			case "time_limit", "memory_limit":
				limit, err := strconv.Atoi(tagString[0][1])
				if err != nil {
					restweb.Logger.Debug(err)
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
	proModel := model.ProblemModel{}
	pid, err := proModel.Insert(problem)
	if err != nil {
		restweb.Logger.Debug(err)
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

	pc.Redirect("/admin/problems", http.StatusFound)
}
