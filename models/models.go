package models

import (
	"fmt"
	"log"
	"time"

	"github.com/bianpeijiang/ginblog/pkg/setting"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v ", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)

	db, err = gorm.Open(dbType, args)
	if err != nil {
		log.Println(err)
	}

	//设置默认表名加前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	//使用单数表名
	db.SingularTable(true)
	//设置log模式，true为详细日志，false无日志 default将只输出错误日志
	db.LogMode(true)
	// SetMaxIdleCon设置空闲连接池中的最大连接数。
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns设置数据库的最大打开连接数。
	db.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetime设置可以重复使用连接的最长时间。
	db.DB().SetConnMaxLifetime(time.Hour)
}

func CloseDB() {
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)
}
