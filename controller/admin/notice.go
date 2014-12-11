package admin

import (
	"GoOnlineJudge/class"
	"html/template"
	"net/http"
	"os"
	"restweb"
)

type AdminNotice struct {
	class.Controller
}

func (n *AdminNotice) Index() {
	restweb.Logger.Debug("Admin notice index")
	n.Output["Msg"] = string(n.Output["Msg"].(template.HTML))
	n.Output["IsNotice"] = true
	n.RenderTemplate("view/admin/layout.tpl", "view/admin/notice.tpl")
}
func (n *AdminNotice) Edit() {
	restweb.Logger.Debug("Admin notice edit")

	msg := n.Input.Get("msg")
	file, err := os.Create("view/admin/msg.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.Write([]byte(msg))
	n.Redirect("/admin/notice", http.StatusFound)
}
