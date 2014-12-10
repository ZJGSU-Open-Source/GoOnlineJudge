package admin

import (
	"GoOnlineJudge/class"
	"os"
	"restweb"
)

type NoticeAdmin struct {
	class.Controller
}

func (n *NoticeAdmin) Index() {
	n.RenderTemplate("view/admin/layout.tpl", "view/admin/notice.tpl")
}
func (n *NoticeAdmin) Set() {
	restweb.Logger.Debug("Admin set notice")

	msg := n.Input.Get("msg")
	file, err := os.Create("view/admin/msg.txt")
	if err != nil {
		restweb.Logger.Debug(err)
		return
	}
	defer file.Close()
	file.Write([]byte(msg))
	n.Redirect("/admin/notice", 307)
}
