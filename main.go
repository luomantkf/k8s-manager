package main

import (
	"fmt"
	"private-cloud-k8s/conf"
	"private-cloud-k8s/pkg/db"
	"private-cloud-k8s/pkg/server"
	"private-cloud-k8s/pkg/setting"
)

func init() {
	setting.Load()
	db.Load()
	server.Load()
}

// @title 管理k8s平台API
// @version 1.0
// @description 管理k8s平台API
func main() {
	fmt.Println(conf.Conf.ServerConfig.HTTPPort)
}
