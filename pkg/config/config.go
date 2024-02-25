package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiUrl      string
	AccessToken string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	return &Config{
		ApiUrl:      os.Getenv("GHOSTFOLIO_API_URL"),
		AccessToken: os.Getenv("GHOSTFOLIO_ACCESS_TOKEN"),
	}
}
