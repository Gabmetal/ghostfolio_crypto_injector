package blockchain

type Transaction struct {
	// Define aquí los campos comunes que necesitas de una transacción
}

type BlockchainClient interface {
	GetTransactions(address string) ([]Transaction, error)
	GetTokenTransactions(address string) ([]Transaction, error)
}
