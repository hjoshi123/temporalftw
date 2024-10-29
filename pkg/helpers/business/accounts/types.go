package helpers

type CreateAccountInput struct {
	BankID      int     `json:"bank_id,omitempty"`
	AccountType string  `json:"account_type,omitempty"`
	Balance     float64 `json:"balance,omitempty"`
}

type CreateAccountOutput struct {
	BankID        int   `json:"bank_id,omitempty"`
	AccountNumber int64 `json:"account_number,omitempty"`
}
