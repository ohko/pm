package main

import (
	"flag"
	"log"
	"runtime"

	"pm/controller"
	"pm/model"

	"github.com/ohko/logger"
)

var (
	addr         = flag.String("s", ":8080", "server address")
	sessionPath  = flag.String("sp", "/tmp/pm_session", "session path")
	dbPath       = flag.String("db", "./db/sqlite3.db", "database path")
	resetAdmin   = flag.String("resetAdmin", "", "reset admin new password")
	ll           = logger.NewLogger()
	oauth2Server = flag.String("o2", "http://127.0.0.1:8000", "oauth2 server")
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Flags() | log.Lshortfile)

	// 初始化数据库
	if err := model.Init(ll, *dbPath); err != nil {
		ll.Log3Fatal(err)
	}

	// 重设管理员密码
	if *resetAdmin != "" {
		if err := model.ResetAdmin(*resetAdmin); err != nil {
			ll.Log4Trace(err)
		}
		return
	}

	controller.Start(*addr, *sessionPath, *oauth2Server, ll)
}
