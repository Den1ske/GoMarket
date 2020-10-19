package apiserver
import (
	//"github.com/Den1ske/GoMarket/src/internal/app/store"
)

type Config struct {
	APIServerPort string `toml:"server_port"`
	LogLevel string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	//Store *store.Config
}

func NewConfig() *Config {
	return &Config {
		APIServerPort: ":80",
		LogLevel: "debug",
		//Store: store.NewConfig(),
	}
}