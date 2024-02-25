package importer

import (
	"bytes"
	"encoding/json"
	"ghostfolio_crypto_injector/pkg/blockchain"
	"ghostfolio_crypto_injector/pkg/config"
	"ghostfolio_crypto_injector/pkg/models"
	"net/http"
)

func ImportData(cfg *config.Config, authToken string, client blockchain.BlockchainClient, address string) error {

	importData := models.ImportData{
		// llenar los datos de importación
	}

	importDataJson, err := json.Marshal(importData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", cfg.ApiUrl+"/api/v1/import", bytes.NewBuffer(importDataJson))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// procesar la respuesta de la importación

	return nil
}
