package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"

	"restweb"

	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

type AdminTestdata struct {
	class.Controller
}

// List 列出对应题目的test data，method：GET
func (tc *AdminTestdata) List(pid string) {
	restweb.Logger.Debug("Admin testdata list")

	file := []string{}
	dir, err := os.Open(config.Datapath + pid)
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		restweb.Logger.Debug(err)
		tc.Error("Problem Id error", 500)
		return
	} else {
		for _, fi := range files {
			if !fi.IsDir() {
				file = append(file, fi.Name())
			}
		}
	}

	tc.Output["Files"] = file
	tc.Output["Pid"] = pid
	tc.Output["Title"] = "Problem" + pid + " - Test data"
	tc.Output["IsProblem"] = true

	tc.RenderTemplate("view/admin/layout.tpl", "view/admin/test_data.tpl")
}

// 上传测试数据,URL /admin/testdata/upload?pid=<pid>，method：POST
func (tc *AdminTestdata) Upload(pid string) {
	restweb.Logger.Debug("Admin Upload files")

	if tc.Privilege != config.PrivilegeAD {
		tc.Err400("Warning", "Error Privilege to Update testdate")
		return
	}
	r := tc.Requset
	r.ParseMultipartForm(32 << 20)
	fhs := r.MultipartForm.File["testfiles"]
	os.Mkdir(config.Datapath+pid, os.ModePerm)
	for _, fheader := range fhs {
		filename := fheader.Filename
		file, err := fheader.Open()
		if err != nil {
			restweb.Logger.Debug(err)
			return
		}
		defer file.Close()
		//保存文件
		f, err := os.Create(config.Datapath + pid + "/" + filename)
		if err != nil {
			restweb.Logger.Debug(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
	tc.Redirect("/admin/testdata/"+pid, http.StatusFound)
}

// Download 下载测试数据,URL:/admin/testdata/download?type=<type>，method:POST
func (tc *AdminTestdata) Download(pid string) {
	restweb.Logger.Debug("Admin Download files")

	filename := tc.Input.Get("type")
	file, err := os.Open(config.Datapath + pid + "/" + filename)
	if err != nil {
		restweb.Logger.Debug(err)
		return
	}
	defer file.Close()
	finfo, _ := file.Stat()
	tc.Response.Header().Add("ContentType", "application/octet-stream")
	tc.Response.Header().Add("Content-disposition", "attachment; filename="+filename)
	tc.Response.Header().Add("Content-Length", strconv.Itoa(int(finfo.Size())))
	io.Copy(tc.Response, file)
}

// Delete 删除指定testdata，URL:/admin/testdata/delete?type=<type>
func (tc *AdminTestdata) Delete(pid string) {
	restweb.Logger.Debug("Admin TestData Delete")

	if tc.Privilege != config.PrivilegeAD {
		tc.Err400("Warning", "Error Privilege to Delete testdate")
		return
	}

	filetype := tc.Input.Get("type")

	cmd := exec.Command("rm", config.Datapath+pid+"/"+filetype)
	err := cmd.Run()
	if err != nil {
		restweb.Logger.Debug(err)
	}
	tc.Response.WriteHeader(200)
}
