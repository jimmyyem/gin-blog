package models

import (
	"fmt"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

type Model struct {
	ID    int `gorm:"primary_key" json:"id"`
	Ctime int `json:"ctime"`
	Mtime int `json:"mtime"`
}

const (
	STATE_ONLINE  = 0
	STATE_OFFLINE = 1
	STATE_INVALID = -1
)

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	if setting.RunMode == "debug" {
		db.LogMode(true)
		db.SetLogger(logging.SqlLogger)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// 公用查询用的map
func CommonMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["state"] = STATE_ONLINE

	return maps
}
