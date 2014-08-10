package admin

import (
	"GoOnlineJudge/class"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

//ImageController handles sth. with images
type ImageController struct {
	class.Controller
}

type image struct {
	Error int    `json:"error"`
	Url   string `json:"url"`
}

//Upload support kindeditor upload images,the response must return json eg. like {"err":0,"url":"http:...."}
func (this *ImageController) Upload(w http.ResponseWriter, r *http.Request) {
	log.Println("AdminUpload Image")
	this.Init(w, r)

	r.ParseMultipartForm(32 << 20)
	fhs := r.MultipartForm.File["imgFile"]

	var path string
	var errflag int

	for _, fheader := range fhs {
		filename := fheader.Filename
		log.Println(filename)
		file, err := fheader.Open()
		if err != nil {
			log.Println(err)
			errflag++
			break
		}
		defer file.Close()
		//保存文件
		path = "static/img/" + filename
		f, err := os.Create(path)
		if err != nil {
			log.Println(err)
			errflag++
			break
		}
		defer f.Close()
		io.Copy(f, file)
	}
	im := &image{Error: errflag, Url: "/" + path}
	b, _ := json.Marshal(im)
	w.Write(b)
}
