package main

import (
	"cache/pkg/log"
	"cache/server/api"
	"cache/server/config"
	"net/http"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Errorln(err)
	}
	config.HotLoad()
	log.Infof("load config:%+v", config.Conf)

	httpServer := api.NewHttpServer()
	err := http.ListenAndServe(config.Conf.Port, httpServer)
	if err != nil {
		log.Errorln(err)
	}
}
