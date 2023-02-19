package main

import (
	"cache/pkg/log"
	"cache/server/api"
	"cache/server/config"
	"cache/server/core"
	"hash/crc32"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "go-cache",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"f"},
				Usage:   "-f etc/conf.yaml",
			},
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   ":9000",
			},
			&cli.StringFlag{
				Name:  "peers",
				Usage: "-endpoints=127.0.0.1:9000,127.0.0.1:9001",
				Value: "127.0.0.1:9000",
			},
		},
		Action: Main,
	}

	if err := app.Run(os.Args); err != nil {
		log.Errorln(err)
	}
}

func Main(ctx *cli.Context) error {
	confPath := ctx.String("path")
	port := ctx.String("port")
	peers := strings.Split(ctx.String("peers"), ",")
	config.Conf = config.Config{
		IP:    "127.0.0.1",
		Port:  port,
		Peers: peers,
	}

	log.SetLevel(log.InfoLevel)

	if err := config.InitConfig(confPath); err != nil {
		log.Errorln(err)
	}
	config.HotLoad(confPath)
	log.Infof("Load config:%+v", config.Conf)

	core.InitCache(config.Conf.Peers, crc32.ChecksumIEEE)
	go core.CheckHealth()
	log.Infoln("Load core")

	log.Infof("Listen port%v...", config.Conf.Port)
	httpServer := api.NewHttpServer()
	if err := http.ListenAndServe(config.Conf.Port, httpServer); err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}
