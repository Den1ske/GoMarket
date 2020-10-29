package config

type ConfigStruct struct {
	DBHost  string `envconfig:"DB_HOST"`
	DBUser  string `envconfig:"DB_USER"`
	DBPass string `envconfig:"DB_PASS"`
	DBName  string `envconfig:"DB_NAME"`
}
