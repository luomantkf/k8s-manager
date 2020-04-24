package setting

import (
	"github.com/go-ini/ini"
	"log"
	"private-cloud-k8s/conf"
	"time"
)

var ConfFile *ini.File

func init() {
	var err error
	ConfFile, err = ini.Load("app.ini")
	if err != nil {
		log.Fatalf("读取配置文件错误：%v", err)
	}

}

//从配置文件中读取项目配置
func Load() {
	conf.Conf = conf.Config{}
	loadBase()
	loadDatabase()
	loadServer()
}

func loadBase() {
	conf.Conf.RunMode = ConfFile.Section("").Key("RUN_MODE").MustString("debug")
}

func loadDatabase() {
	databaseConf := conf.DatabaseConfig{}
	databaseConf.Host = ConfFile.Section("database").Key("HOST").MustString("localhost")
	databaseConf.Name = ConfFile.Section("database").Key("NAME").String()
	databaseConf.User = ConfFile.Section("database").Key("USER").MustString("root")
	databaseConf.Password = ConfFile.Section("database").Key("PASSWORD").MustString("")
	databaseConf.TablePrefix = ConfFile.Section("database").Key("TABLE_PREFIX").String()
	databaseConf.Type = ConfFile.Section("database").Key("TYPE").MustString("mysql")

	conf.Conf.DatabaseConfig = databaseConf
}

func loadServer() {
	serverConf := conf.ServerConfig{}
	serverConf.HTTPPort = ConfFile.Section("server").Key("HTTP_PORT").MustInt(8080)
	serverConf.ReadTimeout = time.Duration(ConfFile.Section("server").Key("READ_TIMEOUT").MustInt(60)) * time.Second
	serverConf.WriteTimeout = time.Duration(ConfFile.Section("server").Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	serverConf.ServerContextPath = ConfFile.Section("server").Key("SERVER_CONTEXT_PATH").MustString("")

	conf.Conf.ServerConfig = serverConf
}
