package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"private-cloud-k8s/conf"
	"private-cloud-k8s/pkg/logging"
)

var DB *gorm.DB

//配置数据库连接
func Load() {
	dbConf := conf.Conf.DatabaseConfig

	DB, err := gorm.Open(dbConf.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConf.User,
		dbConf.Password,
		dbConf.Host,
		dbConf.Name))

	if err != nil {
		logging.Error("连接数据库[{}]异常！", dbConf.Name)
		return
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func Close() {
	defer DB.Close()
}
