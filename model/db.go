package model

import (
	"fmt"
	"ginblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	//db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   //string类型字段的默认长度
		DisableDatetimePrecision:  true,  //禁用datatime精度， MySQL5.6之前的数据库不支持
		DontSupportRenameIndex:    true,  //重命名索引是采用删除并新建的方式，MySQL5.7之前的数据库和MariaDB不支持重命名索引
		DontSupportRenameColumn:   true,  //用`change`重命名列，MySQL8之前的数据库和MariaDB不支持重命名列
		SkipInitializeWithVersion: false, //根据当前MySQL版本自动配置
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}

	//获取通用数据帝牙卢卡对象sql.DB, 然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("获取sql.DB失败：", err)
	}
	err = db.AutoMigrate()
	if err != nil {
		fmt.Println("自动迁移数据库错误：", err)
	}
	db.SingularTable(true)
	// 设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	//设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	//设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Second * 10)
}
