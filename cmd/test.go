package main

import (
	"bytes"
	"errors"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	//"time"
)

type ControllerInfo struct {
	PkgPath string
	PkgName string
	Name    string
}
type RouterInfo struct {
	ControllerName string
	URL            template.HTML
	Action         string
	Method         string
}

var ContrInfos []ControllerInfo
var RouterInfos []RouterInfo

func main() {
	defer func() {
		//time.Sleep(1 * time.Second)
		log.Println("here")
		// os.Remove("main.go")
		// os.Remove("config/router.conf")
	}()
	filepath.Walk("../controller", walkFn)
	generateMain()
	generateRouter()
	os.Chdir("../")
	run()

}

func walkFn(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		return nil
	}

	fset := &token.FileSet{}
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments|parser.AllErrors)
	if err != nil {
		log.Println(err)
		return err
	}
	var pkg *ast.Package
	for _, v := range pkgs {
		pkg = v
	}

	goPath := build.Default.GOPATH

	wd, _ := os.Getwd()
	wd = wd[:len(wd)-4]
	walkAstFiles(fset, wd[len(goPath+"/src/"):]+path[2:], pkg)

	return nil
}

func walkAstFiles(fset *token.FileSet, path string, pkg *ast.Package) {
	ControllerName := ""
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {

			if funcdecl, ok := decl.(*ast.FuncDecl); ok && funcdecl.Doc != nil {
				for _, cmt := range funcdecl.Doc.List {
					if strings.HasPrefix(cmt.Text, "//@") {
						url, method, err := parseDecorator(cmt.Text)
						if err != nil {
							log.Fatal("[error] ", fset.Position(cmt.Pos()), err)
						}
						url = strings.TrimRight(url, "/") + "/"
						RouterInfos = append(RouterInfos,
							RouterInfo{ControllerName: ControllerName,
								URL:    template.HTML(url),
								Action: funcdecl.Name.Name,
								Method: method})
					}
				}
			}

			if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.TYPE {
				spec := gen.Specs[0]
				if ts, ok := spec.(*ast.TypeSpec); ok && ts.Comment.Text() == "@Controller\n" {
					ControllerName = ts.Name.Name
					ContrInfos = append(ContrInfos, ControllerInfo{PkgPath: path, PkgName: pkg.Name, Name: ControllerName})
				}
			}
		}
	}
}

func parseDecorator(decorator string) (url, method string, err error) {
	des := strings.Split(decorator, "@")
	if len(des) < 3 {
		err = errors.New("Decorators miss")
		return
	}
	urlPair := strings.Split(des[1], ":")
	if len(urlPair) < 2 {
		err = errors.New("Decorators value miss")
		return
	}
	url = strings.Trim(urlPair[1], " ")

	methodPair := strings.Split(des[2], ":")
	if len(methodPair) < 2 {
		err = errors.New("Decorators value miss")
		return
	}
	method = strings.Trim(methodPair[1], " ")
	return
}

func generateMain() {

	tpl := `package main

	import (
	"restweb"
	"log"
	{{with .ContrInfos}}
	{{range .}}"{{.PkgPath}}"
	{{end}}
	{{end}}
	)

	func main(){
	{{with .ContrInfos}}
	{{range .}}restweb.RegisterController(&{{.PkgName}}.{{.Name}}{})
	{{end}}
	{{end}}
	restweb.AddFile("/static/", ".")
	log.Fatal(restweb.Run())
	}
	`
	t, err := template.New("foo").Parse(tpl)
	if err != nil {
		log.Println(err)
		return
	}
	bf := bytes.NewBufferString("")
	data := make(map[string]interface{})
	data["ContrInfos"] = ContrInfos
	err = t.Execute(bf, data)
	if err != nil {
		log.Println(err)
	}
	b, err := format.Source([]byte(bf.String()))
	if err != nil {
		log.Println(err)
	}
	f, err := os.Create("../main.go")
	if err != nil {
		log.Println(err)
		return
	}
	f.Write(b)
}

func generateRouter() {
	tpl := `
{{with .RouterInfos}}
{{range .}}{{.Method}} 	^{{.URL}}$	 {{.ControllerName}}.{{.Action}}
{{end}}
{{end}}
`
	t, err := template.New("foo").Parse(tpl)
	if err != nil {
		log.Println(err)
		return
	}
	bf, err := os.Create("../config/router.conf")
	if err != nil {
		log.Println(err)
	}
	data := make(map[string]interface{})
	data["RouterInfos"] = RouterInfos
	err = t.Execute(bf, data)
	if err != nil {
		log.Println(err)
	}
}

func run() {
	cmd := exec.Command("go", "build")
	cmd.Run()
	cmd = exec.Command("./GoOnlineJudge")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
