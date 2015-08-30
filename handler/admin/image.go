package admin

import (
	"github.com/zenazn/goji/web"

	"encoding/json"
	"io"
	"net/http"
	"os"
)

type image struct {
	Error int    `json:"error"`
	Url   string `json:"url"`
}

//Upload support kindeditor upload images,the W must return json eg. like {"err":0,"url":"http:...."}
//@URL:/admin/images/ @method: POST
func PostImage(c web.C, w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)
	fhs := r.MultipartForm.File["imgFile"]

	var path string
	var errflag int

	for _, fheader := range fhs {
		filename := fheader.Filename
		file, err := fheader.Open()
		if err != nil {
			errflag++
			break
		}
		defer file.Close()
		//保存文件
		path = "static/img/" + filename
		f, err := os.Create(path)
		if err != nil {
			errflag++
			break
		}
		defer f.Close()
		io.Copy(f, file)
	}

	im := &image{Error: errflag, Url: "/" + path}
	json.NewEncoder(w).Encode(im)
}
