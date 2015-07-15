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
}   //@Controller

// List 列出对应题目的test data
//@URL: /admin/testdata/(\d+) @method: GET
func (tc *AdminTestdata) List(pid string) {
    restweb.Logger.Debug("Admin testdata list")

    file := []string{}
    dir, err := os.Open(config.Datapath + pid)
    defer dir.Close()

    files, err := dir.Readdir(-1)
    if err != nil {
        restweb.Logger.Debug(config.Datapath + pid)
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

// 上传测试数据
//@URL: /admin/testdata/(\d+) @method: POST
func (tc *AdminTestdata) Upload(pid string) {
    restweb.Logger.Debug("Admin Upload files")

    if tc.Privilege != config.PrivilegeAD {
        tc.Err400("Warning", "Error Privilege to Update testdate")
        return
    }
    r := tc.R
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

// Download 下载测试数据
//@URL: /admin/testdata/(\d+)/file @method: GET
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
    tc.W.Header().Add("ContentType", "application/octet-stream")
    tc.W.Header().Add("Content-disposition", "attachment; filename="+filename)
    tc.W.Header().Add("Content-Length", strconv.Itoa(int(finfo.Size())))
    io.Copy(tc.W, file)
}

// Delete 删除指定testdata
//@URL: /admin/testdata/(\d+) @method: DELETE
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
    tc.W.WriteHeader(200)
}
