package helpers

type TransactionsInput struct {
	FromAccountNumber int64   `json:"from_account_number"`
	FromBankID        int     `json:"from_bank_id"`
	ToAccountNumber   int64   `json:"to_account_number"`
	ToBankID          int     `json:"to_bank_id"`
	Memo              string  `json:"memo"`
	TxType            string  `json:"tx_type"`
	Amount            float64 `json:"amount"`
}

type TransactionsOutput struct {
	Msg    string `json:"msg"`
	TxUUID string `json:"tx_uuid"`
}

type ApproveTransactionInput struct {
	TxUUID string `json:"tx_uuid"`
}

type RejectTransactionInput struct {
	TxUUID         string `json:"tx_uuid"`
	RejectedReason string `json:"rejected_reason"`
}
