package class

import (
	"html/template"
)

var funcMap map[string]interface{}

func AddFuncMap(key string, f interface{}) {
	funcMap[key] = f
}

func ParseFiles(tplfiles ...string) (*template.Template, error) {
	t := template.New("layout.tpl").Funcs(template.FuncMap(funcMap))
	t, err := t.ParseFiles(tplfiles...)
	return t, err
}

func initFuncMap() {
	funcMap = make(map[string]interface{})
	funcMap["NumAdd"] = NumAdd
	funcMap["NumSub"] = NumSub
	funcMap["ShowRatio"] = ShowRatio
	funcMap["ShowSpecial"] = ShowSpecial
	funcMap["ShowJudge"] = ShowJudge
	funcMap["ShowLanguage"] = ShowLanguage
	funcMap["ShowEncrypt"] = ShowEncrypt
	funcMap["ShowPrivilege"] = ShowPrivilege
	funcMap["LargePU"] = LargePU
	funcMap["ShowTime"] = ShowTime
	funcMap["ShowStatus"] = ShowStatus
	funcMap["ShowSim"] = ShowSim
}
