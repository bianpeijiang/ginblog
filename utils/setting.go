package utils

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadData(file *ini.File) {
	DbSec := file.Section("database")
	Db = DbSec.Key("Db").MustString("mysql")
	DbHost = DbSec.Key("DbHost").MustString("localhost")
	DbPort = DbSec.Key("DbPort").MustString("3306")
	DbUser = DbSec.Key("DbUser").MustString("root")
	DbPassWord = DbSec.Key("DbPassWord").MustString("bianpeijiang123")
	DbName = DbSec.Key("DbName").MustString("ginblog")

}
