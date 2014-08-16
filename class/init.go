package class

import (
	"golog"
	"os"
)

var SessionManager *Manager
var Logger *golog.Log

func init() {
	Logger = golog.NewLog(os.Stdout, golog.Ldebug|golog.Linfo)
	SessionManager = NewManager()
	Logger.Debug("Start new session manager")
	go SessionManager.GC()
}
