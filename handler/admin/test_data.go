package admin

import (
	"ojapi/config"
	"ojapi/middleware"

	"github.com/zenazn/goji/web"

	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

// List 列出对应题目的test data
//@URL: /api/admin/testdata/:pid @method: GET
func ListTestdata(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		pid = c.URLParams["pid"]
	)

	file := []string{}
	dir, err := os.Open(config.Datapath + pid)
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		for _, fi := range files {
			if !fi.IsDir() {
				file = append(file, fi.Name())
			}
		}
	}

}

// 上传测试数据
//@URL: /admin/testdata/:pid @method: POST
func UploadTestdata(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		pid  = c.URLParams["pid"]
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	r.ParseMultipartForm(32 << 20)
	fhs := r.MultipartForm.File["testfiles"]
	os.Mkdir(config.Datapath+pid, os.ModePerm)
	for _, fheader := range fhs {
		filename := fheader.Filename
		file, err := fheader.Open()
		if err != nil {
			return
		}
		defer file.Close()
		//保存文件
		f, err := os.Create(config.Datapath + pid + "/" + filename)
		if err != nil {
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}

	w.WriteHeader(http.StatusCreated)
}

// Download 下载测试数据
//@URL: /admin/testdata/:pid/file @method: GET
func DownloadTestdata(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		pid = c.URLParams["pid"]
	)

	filename := r.URL.Query().Get("type")
	file, err := os.Open(config.Datapath + pid + "/" + filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	finfo, _ := file.Stat()
	w.Header().Add("ContentType", "application/octet-stream")
	w.Header().Add("Content-disposition", "attachment; filename="+filename)
	w.Header().Add("Content-Length", strconv.Itoa(int(finfo.Size())))
	io.Copy(w, file)
}

// Delete 删除指定testdata
//@URL: /api/admin/testdata/:pid @method: DELETE
func Delete(c web.C, w http.ResponseWriter, r *http.Request) {
	var (
		pid  = c.URLParams["pid"]
		user = middleware.ToUser(c)
	)

	if user.Privilege != config.PrivilegeAD {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	filetype := r.URL.Query().Get("type")

	cmd := exec.Command("rm", config.Datapath+pid+"/"+filetype)
	err := cmd.Run()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
