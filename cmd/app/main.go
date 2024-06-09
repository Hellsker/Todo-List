package main

import (
	"github.com/Hellsker/Todo-List/config"
	"github.com/Hellsker/Todo-List/internal/app"
	"log"
)

var (
	confPath = "./config/local.yaml"
)

func main() {
	cfg, err := config.Load(confPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	app.Run(cfg)
}
