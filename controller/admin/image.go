package admin

import (
	"GoOnlineJudge/class"

	"restweb"

	"encoding/json"
	"io"
	"os"
)

//ImageController handles sth. with images
type AdminImage struct {
	class.Controller
} //@Controller

type image struct {
	Error int    `json:"error"`
	Url   string `json:"url"`
}

//Upload support kindeditor upload images,the W must return json eg. like {"err":0,"url":"http:...."}
//@URL:/admin/images/ @method: POST
func (ic AdminImage) Post() {
	restweb.Logger.Debug("AdminUpload Image")

	r := ic.R
	r.ParseMultipartForm(32 << 20)
	fhs := r.MultipartForm.File["imgFile"]

	var path string
	var errflag int

	for _, fheader := range fhs {
		filename := fheader.Filename
		restweb.Logger.Debug(filename)
		file, err := fheader.Open()
		if err != nil {
			restweb.Logger.Debug(err)
			errflag++
			break
		}
		defer file.Close()
		//保存文件
		path = "static/img/" + filename
		f, err := os.Create(path)
		if err != nil {
			restweb.Logger.Debug(err)
			errflag++
			break
		}
		defer f.Close()
		io.Copy(f, file)
	}
	im := &image{Error: errflag, Url: "/" + path}
	b, _ := json.Marshal(im)
	ic.W.Write(b)
}
