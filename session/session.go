package session

import (
	"log"
	"net/http"

	"GoOnlineJudge/model"
)

var SessionManager = NewManager()

func GetUser(r *http.Request) *model.User {
	sid := r.FormValue("session_id")

	if sid != "" {
		uid := GetSession(sid, "Uid")
		userModel := &model.UserModel{}
		user, err := userModel.Detail(uid)
		if err != nil {
			return nil
		}
		return user
	}
	return nil
}

func SetSession(sid string, key string, value string) {
	session := SessionManager.GetSession(sid)
	session.Set(key, value)
}

func GetSession(sid string, key string) string {
	sess := SessionManager.GetSession(sid)
	return sess.Get(key)
}

func DeleteSession(r *http.Request) {
	sid := r.FormValue("session_id")
	SessionManager.DeleteSession(sid)
}

func StartSession() *Session {
	return SessionManager.StartSession()
}
