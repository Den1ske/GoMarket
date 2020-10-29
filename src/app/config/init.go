package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

var Config ConfigStruct

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = envconfig.Process("", &Config)
	if err != nil {
		log.Fatal(err.Error())
	}

}
