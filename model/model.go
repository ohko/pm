package model

import (
	"os"
	"pm/util"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // ...
	"github.com/ohko/logger"
)

var (
	ll  *logger.Logger
	db  *gorm.DB
	loc *time.Location
)

// Init ...
func Init(lll *logger.Logger, dbPath string) error {
	ll = lll

	var err error
	loc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}

	if err := initDB(dbPath); err != nil {
		return err
	}

	return nil
}

// initDB 初始化数据库
func initDB(dbPath string) error {
	var err error

	if db, err = gorm.Open("sqlite3", dbPath); err != nil {
		return err
	}
	if os.Getenv("DEBUG") != "" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.SetNowFuncOverride(func() time.Time {
		return time.Now().In(loc)
	})

	if err := db.AutoMigrate(&Member{}, &User{}, &Project{}, &Task{}).Error; err != nil {
		return err
	}
	var m Member
	if err := db.First(&m).Error; err != nil {
		if err := db.Save(&Member{User: "admin", Pass: string(util.Hash([]byte("admin")))}).Error; err != nil {
			return err
		}
	}
	return nil
}

// Close ...
func Close() {
	if db != nil {
		db.Close()
		db = nil
	}
}
