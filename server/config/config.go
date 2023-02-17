package config

import (
	"cache/pkg/log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	Port string `yaml:"port"`
}

func InitConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("conf.yaml")
	viper.AddConfigPath("./server/etc/")
	if err := viper.ReadInConfig(); err != nil {
		log.Errorln(err)
		return err
	}
	viper.Unmarshal(&Conf)

	return nil
}

func HotLoad() {
	go func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			InitConfig()
		})
	}()
}
