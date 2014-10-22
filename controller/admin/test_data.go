package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"

	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type TestdataController struct {
	class.Controller
}

func (tc TestdataController) Route(w http.ResponseWriter, r *http.Request) {
	tc.Init(w, r)
	action := tc.GetAction(r.URL.Path, 2)
	defer func() {
		if e := recover(); e != nil {
			http.Error(w, "no such page", 404)
		}
	}()
	rv := class.GetReflectValue(w, r)
	class.CallMethod(&tc, strings.Title(action), rv)
}

// List 列出对应题目的test data，method：GET
func (tc *TestdataController) List(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin testdata list")

	pid := r.URL.Query().Get("pid")
	file := []string{}
	dir, err := os.Open(config.Datapath + pid)
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		class.Logger.Debug(err)
		http.Error(w, "Problem Id error", 500)
		return
	} else {
		for _, fi := range files {
			if !fi.IsDir() {
				file = append(file, fi.Name())
			}
		}
	}

	tc.Data["Files"] = file
	tc.Data["Pid"] = pid
	tc.Data["Title"] = "Problem" + pid + " - Test data"
	tc.Data["IsProblem"] = true

	tc.Execute(w, "view/admin/layout.tpl", "view/admin/test_data.tpl")
}

// 上传测试数据,URL /admin/testdata/upload?pid=<pid>，method：POST
func (tc *TestdataController) Upload(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Upload files")
	if r.Method != "POST" {
		tc.Err400(w, r, "Error", "Error Method to Update testdate")
		return
	}

	if tc.Privilege != config.PrivilegeAD {
		tc.Err400(w, r, "Warning", "Error Privilege to Update testdate")
		return
	}

	pid := r.URL.Query().Get("pid")

	r.ParseMultipartForm(32 << 20)
	fhs := r.MultipartForm.File["testfiles"]
	os.Mkdir(config.Datapath+pid, os.ModePerm)
	for _, fheader := range fhs {
		filename := fheader.Filename
		file, err := fheader.Open()
		if err != nil {
			class.Logger.Debug(err)
			return
		}
		defer file.Close()
		//保存文件
		f, err := os.Create(config.Datapath + pid + "/" + filename)
		if err != nil {
			class.Logger.Debug(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
	http.Redirect(w, r, "/admin/testdata/list?pid="+pid, http.StatusFound)
}

// Download 下载测试数据,URL:/admin/testdata/download?type=<type>，method:POST
func (tc *TestdataController) Download(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin Download files")
	tc.Init(w, r)

	args := r.URL.Query()
	filename := args.Get("type")
	file, err := os.Open(config.Datapath + args.Get("pid") + "/" + filename)
	if err != nil {
		class.Logger.Debug(err)
		return
	}
	defer file.Close()
	finfo, _ := file.Stat()
	w.Header().Add("ContentType", "application/octet-stream")
	w.Header().Add("Content-disposition", "attachment; filename="+filename)
	w.Header().Add("Content-Length", strconv.Itoa(int(finfo.Size())))
	io.Copy(w, file)
}

// Delete 删除指定testdata，URL:/admin/testdata/delete?type=<type>
func (tc *TestdataController) Delete(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("Admin TestData Delete")
	if r.Method != "POST" {
		tc.Err400(w, r, "Error", "Error Method to Delete testdate")
		return
	}

	if tc.Privilege != config.PrivilegeAD {
		tc.Err400(w, r, "Warning", "Error Privilege to Delete testdate")
		return
	}

	args := r.URL.Query()
	filetype := args.Get("type")
	pid := args.Get("pid")

	cmd := exec.Command("rm", config.Datapath+pid+"/"+filetype)
	err := cmd.Run()
	if err != nil {
		class.Logger.Debug(err)
	}
	http.Redirect(w, r, "/admin/testdata/list?pid="+pid, http.StatusFound)
}
