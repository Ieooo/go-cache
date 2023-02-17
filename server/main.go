package main

import (
	"cache/pkg/log"
	"cache/server/config"
	"net/http"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Errorln(err)
	}
	http.ListenAndServe(config.Conf.Port, nil)
}
