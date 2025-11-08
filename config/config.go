package config

import (
	"go-transaction-service/common/util"

	"github.com/sirupsen/logrus"
)

var Config AppConfig

type AppConfig struct {
	Port    int    `json:"port"`
	AppName string `json:"appName"`
	AppEnv  string `json:"appEnv"`
}

func Init() {
	err := util.BindFromJSON(&Config, "config.json", ".")
	if err != nil {
		logrus.Infof("failed to bind config: %v", err)
	}
}
