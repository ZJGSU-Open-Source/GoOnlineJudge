package contest

import (
	"GoOnlineJudge/class"
	// 	"GoOnlineJudge/config"
	"net/http"
	// 	"strconv"
)

type ranklist struct {
	Uid      string
	costtime int
}
type RanklistController struct {
	class.Controller
}

func Index(w http.ResponseWriter, r *http.Request) {
	return
}
