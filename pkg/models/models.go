package models

type AuthResponse struct {
	AuthToken string `json:"authToken"`
}

type ImportData struct {
	Accounts   []interface{} `json:"accounts"`
	Activities []interface{} `json:"activities"`
}

type EtherscanResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Result  []EtherscanTransaction `json:"result"`
}

type EtherscanTransaction struct {
	// Define los campos según la respuesta de Etherscan
}

type Transaction struct {
	// Define los campos según el formato de importación de Ghostfolio
}
