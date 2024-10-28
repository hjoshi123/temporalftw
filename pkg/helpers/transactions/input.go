package helpers

type TransactionsInput struct {
	FromAccountNumber int64  `json:"from_account_number"`
	FromBankID        int    `json:"from_bank_id"`
	ToAccountNumber   int64  `json:"to_account_number"`
	ToBankID          int    `json:"to_bank_id"`
	Memo              string `json:"memo"`
}
