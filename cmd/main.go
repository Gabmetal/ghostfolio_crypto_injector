package main

import (
	"ghostfolio_crypto_injector/internal/auth"
	"ghostfolio_crypto_injector/internal/importer"
	"ghostfolio_crypto_injector/pkg/config"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	authToken, err := auth.GetAuthToken(cfg)
	if err != nil {
		log.Fatalf("Error obteniendo el token de autenticación: %v", err)
	}

	if err := importer.ImportData(cfg, authToken); err != nil {
		log.Fatalf("Error importando datos: %v", err)
	}
}
