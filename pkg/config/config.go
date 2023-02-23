package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go-demo-project/domain/model"
	"os"
	"strings"
)

var config *model.Config

func Init() {
	v := viper.New()
	v.SetConfigName(os.Args[1])
	v.SetConfigType("yml")
	v.AutomaticEnv()
	v.AddConfigPath("./pkg/config_env")
	v.AddConfigPath("../pkg/config_env")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		var msgErr any
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			msgErr = "config file not found"
			panic(msgErr)
		} else {
			msgErr = fmt.Errorf("fatal error config file: %s", err)
			panic(msgErr)
		}
	}

	v.Unmarshal(&config)
}

func All() model.Config {
	return *config
}

func Server() model.ServerType {
	return *config.Server
}

func Service() model.ServiceType {
	return *config.Service
}
