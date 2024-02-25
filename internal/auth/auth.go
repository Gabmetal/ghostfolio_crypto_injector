package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ghostfolio_crypto_injector/pkg/config"
	"ghostfolio_crypto_injector/pkg/models"
	"io"
	"net/http"
)

func GetAuthToken(cfg *config.Config) (string, error) {
	requestBody := []byte(fmt.Sprintf(`{"accessToken":"%s"}`, cfg.AccessToken))
	req, err := http.NewRequest("POST", cfg.ApiUrl+"/api/auth/token", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var authResp models.AuthResponse
	if err := json.Unmarshal(body, &authResp); err != nil {
		return "", err
	}

	return authResp.AuthToken, nil
}
