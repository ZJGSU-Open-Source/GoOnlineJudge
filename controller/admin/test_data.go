package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
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
		fp, err := os.Open(config.Datapath + args["pid"] + "/test.in")
		defer fp.Close()
		if os.IsNotExist(err) == false {
			file["testin"] = "test.in"
		}

		fp, err = os.Open(config.Datapath + args["pid"] + "/test.out")
		defer fp.Close()
		if os.IsNotExist(err) == false {
			file["testout"] = "test.out"
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

func (this *TestdataController) Downloadin(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Download files")
	this.Init(w, r)

	if r.Method == "POST" {
		args := this.ParseURL(r.URL.Path[6:])

		file, err := os.Open(config.Datapath + args["pid"] + "/" + "test.in")
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()

		filename := os.Getenv("HOME") + "/" + args["pid"]
		os.Mkdir(filename, os.ModePerm)
		f, err := os.OpenFile(filename+"/test.in", os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Println(err)
			return
		}
		io.Copy(f, file)
		http.Redirect(w, r, "/admin/testdata/list/pid/"+args["pid"], http.StatusFound)
	}
}

func (this *TestdataController) Downloadout(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Download files")
	this.Init(w, r)

	if r.Method == "POST" {
		args := this.ParseURL(r.URL.Path[6:])

		file, err := os.Open(config.Datapath + args["pid"] + "/" + "test.out")
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()

		filename := os.Getenv("HOME") + "/" + args["pid"]
		log.Println(filename)

		os.Mkdir(filename, os.ModePerm)
		f, err := os.OpenFile(filename+"/test.out", os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Println(err)
			return
		}
		io.Copy(f, file)
		http.Redirect(w, r, "/admin/testdata/list/pid/"+args["pid"], http.StatusFound)
	}
}

func (this *TestdataController) Deletein(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin TestData Delete")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path[6:])
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	cmd := exec.Command("rm", config.Datapath+strconv.Itoa(pid)+"/test.in")
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

func (this *TestdataController) Deleteout(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin TestData Delete")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path[6:])
	pid, err := strconv.Atoi(args["pid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}
	cmd := exec.Command("rm", config.Datapath+strconv.Itoa(pid)+"/test.out")
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
