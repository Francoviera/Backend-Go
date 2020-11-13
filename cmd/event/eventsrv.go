package main

import (
	"backendo-go/internal/config"
	"backendo-go/internal/service/event"
	"flag"
	"fmt"
	"os"
)

func main() {
	cfg := initConfig()
	fmt.Println(cfg.DB.Driver)
	fmt.Println(cfg.Version)

	service, _ := event.NewEventService(cfg)
	for _, m := range service.FindAll() {
		fmt.Println(m)
	}
}

func initConfig() *config.Config {
	configFile := flag.String("config", "./config/config.yaml", "this is the service config.")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}
