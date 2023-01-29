package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件讀取失敗，請檢查文件路徑:", err)
		return
	}
	LoadServer(file)
	LoadDatabase(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("sdui5fsj3hf9fi")
}

func LoadDatabase(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassword = file.Section("database").Key("DbPassword").MustString("admin8857810")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
