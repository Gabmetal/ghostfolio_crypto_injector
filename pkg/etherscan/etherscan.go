package etherscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	EtherscanAPIURL = "https://api.etherscan.io/api"
	EndBlock        = "99999999"
)

type EtherscanClient struct {
	APIKey string
}

type EtherscanResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []Transaction `json:"result"`
}

type Transaction struct {
	// Define aquí los campos relevantes de una transacción que necesitas
}

type GhostfolioImportData struct {
	Meta struct {
		Date    string `json:"date"`
		Version string `json:"version"`
	} `json:"meta"`
	Accounts   []Account  `json:"accounts"`
	Activities []Activity `json:"activities"`
}

type Account struct {
	Balance    float64 `json:"balance"`
	Comment    *string `json:"comment"`
	Currency   string  `json:"currency"`
	ID         string  `json:"id"`
	IsExcluded bool    `json:"isExcluded"`
	Name       string  `json:"name"`
	PlatformID *string `json:"platformId"`
}

type Activity struct {
	AccountID  string  `json:"accountId"`
	Comment    *string `json:"comment"`
	Fee        float64 `json:"fee"`
	Quantity   float64 `json:"quantity"`
	Type       string  `json:"type"`
	UnitPrice  float64 `json:"unitPrice"`
	Currency   string  `json:"currency"`
	DataSource string  `json:"dataSource"`
	Date       string  `json:"date"`
	Symbol     string  `json:"symbol"`
}

func NewEtherscanClient(apiKey string) *EtherscanClient {
	return &EtherscanClient{APIKey: apiKey}
}

func (client *EtherscanClient) fetchFromEtherscan(params url.Values) (*EtherscanResponse, error) {
	params.Add("apikey", client.APIKey)
	resp, err := http.Get(fmt.Sprintf("%s?%s", EtherscanAPIURL, params.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result EtherscanResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.Status != "1" {
		return nil, fmt.Errorf("etherscan error: %s", result.Message)
	}

	return &result, nil
}

// Implementa GetTransactions y GetTokenTransactions similar a lo mostrado anteriormente

func TransformToGhostfolioFormat(transactions []Transaction, tokenBalances []TokenBalance) *GhostfolioImportData {
	importData := &GhostfolioImportData{}
	// Configura la metadata
	importData.Meta.Date = time.Now().Format(time.RFC3339)
	importData.Meta.Version = "v2.28.0"

	// Transforma las transacciones y balances en cuentas y actividades de Ghostfolio
	// Aquí necesitarás iterar sobre las transacciones y balances, creando instancias de Account y Activity según corresponda

	return importData
}
