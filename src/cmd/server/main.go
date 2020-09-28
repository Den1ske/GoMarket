package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/Den1ske/GoMarket/src/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "Config file path")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}