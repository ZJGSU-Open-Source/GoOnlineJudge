package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type TestdataController struct {
	class.Controller
}

func (this *TestdataController) List(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin testdata list")
	this.Init(w, r)

	if r.Method == "GET" {
		args := this.ParseURL(r.URL.Path[6:])

		file := make(map[string]string)
		fp, err := os.Open(config.Datapath + args["pid"] + "/sample.in")
		defer fp.Close()
		if os.IsNotExist(err) == false {
			file["testin"] = "sample.in"
			file["testin_path"] = config.Datapath + args["pid"] + "/sample.in"
		}

		fp, err = os.Open(config.Datapath + args["pid"] + "/sample.out")
		defer fp.Close()
		if os.IsNotExist(err) == false {
			file["testout"] = "sample.out"
			file["testout_path"] = config.Datapath + args["pid"] + "/sample.out"
		}

		if len(file) > 0 {
			this.Data["hasFile"] = true
		}
		this.Data["Files"] = file
		this.Data["Pid"] = args["pid"]

		t := template.New("layout.tpl")
		t, err = t.ParseFiles("view/admin/layout.tpl", "view/admin/test_data.tpl")
		if err != nil {
			log.Println(err)
			http.Error(w, "tpl error", 500)
			return
		}

		this.Data["Title"] = "Problem" + args["pid"] + " - Test data"
		this.Data["IsProblem"] = true

		err = t.Execute(w, this.Data)
		if err != nil {
			http.Error(w, "tpl error", 500)
			return
		}
	}
}

func (this *TestdataController) Upload(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Upload files")
	this.Init(w, r)

	if r.Method == "POST" {
		args := this.ParseURL(r.URL.Path[6:])

		r.ParseMultipartForm(32 << 20)
		fhs := r.MultipartForm.File["testfiles"]
		os.Mkdir(config.Datapath+args["pid"], os.ModePerm)
		for _, fheader := range fhs {
			filename := fheader.Filename
			file, err := fheader.Open()
			if err != nil {
				log.Println(err)
				return
			}
			defer file.Close()
			//保存文件
			f, err := os.Create(config.Datapath + args["pid"] + "/" + filename)
			if err != nil {
				log.Println(err)
				return
			}
			defer f.Close()
			io.Copy(f, file)
		}
		http.Redirect(w, r, "/admin/testdata/list/pid/"+args["pid"], http.StatusFound)
	}
}

func (this *TestdataController) Delete(w http.ResponseWriter, r *http.Request) {

}
