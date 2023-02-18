package config

import (
	"cache/pkg/log"
	"path"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	IP    string   `yaml:"ip"`
	Port  string   `yaml:"port"`
	Peers []string `yaml:"peer"`
}

func InitConfig(confPath string) error {
	viper.SetConfigType("yaml")
	log.Debugln(confPath)
	log.Debugln(path.Base(confPath))
	log.Debugln(path.Dir(confPath))
	viper.SetConfigName(path.Base(confPath))
	viper.AddConfigPath(path.Dir(confPath))
	if err := viper.ReadInConfig(); err != nil {
		log.Errorln(err)
		return err
	}
	viper.Unmarshal(&Conf)

	return nil
}

func HotLoad(confPath string) {
	go func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			InitConfig(confPath)
		})
	}()
}
