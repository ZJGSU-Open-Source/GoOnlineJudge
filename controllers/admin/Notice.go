package admin

import (
	"GoOnlineJudge/classes"
	"html/template"
	"log"
	"net/http"
)

type NoticeController struct {
	classes.Controller
}

func (this *NoticeController) GET(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Menu")
	this.Init(w, r)

	t, _ := template.ParseFiles("views/admin/menu.tpl", "views/admin/notice.tpl")

	this.Data["Title"] = "Admin Notice"
	t.Execute(w, this.Data)
}
