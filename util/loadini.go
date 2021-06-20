package util

import (
	"fmt"
	"gee/util/log"

	"github.com/go-ini/ini"
)

var (
	AppMode    string
	AppPort    string
	AppName    string
	AppVersion string

	DbType       string
	DbHost       string
	DbPort       string
	DbUser       string
	DbPassWord   string
	DbName       string
	JwtSecretKey string
)

func init() {
	fileIni, err := ini.Load(".ini")
	// if err != nil {
	// 	fmt.Println("读取配置文件错误", err)
	// }
	log.Err("读取配置文件错误", err)
	fmt.Println("读取配置文件...")
	AppMode = fileIni.Section("app").Key("app_mode").MustString("debug")
	AppPort = fileIni.Section("app").Key("app_port").MustString("8080")
	AppName = fileIni.Section("app").Key("app_name").MustString("gee")
	AppVersion = fileIni.Section("app").Key("app_version").MustString("1.0.0")

	DbType = fileIni.Section("db").Key("db_type").MustString("mysql")
	DbHost = fileIni.Section("db").Key("db_host").MustString("localhost")
	DbPort = fileIni.Section("db").Key("db_port").MustString("3306")
	DbUser = fileIni.Section("db").Key("db_user").MustString("user")
	DbPassWord = fileIni.Section("db").Key("db_password").MustString("password")
	DbName = fileIni.Section("db").Key("db_name").MustString("dbname")
	JwtSecretKey = fileIni.Section("jwt").Key("JwtSecretKey").MustString("123456")
}
