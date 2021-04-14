package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/gommon/log"
	"os"
)

func initDb() {
	var err error
	Engine, err = xorm.NewEngine("mysql", dbUser+":"+dbPwd+"@tcp("+dbIp+")/"+dbName+"?charset=utf8")
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
	Engine.SetMaxIdleConns(10)
	Engine.SetMaxOpenConns(20)
	err = Engine.Ping()
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}

func isExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
