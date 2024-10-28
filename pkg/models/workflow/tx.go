package models

type TransactionWorkflowInput struct {
	Amount            float64
	FromAccountNumber int64
	ToAccountNumber   int64
	FromBankID        int
	ToBankID          int
	TxType            string
}

type TransactionWorkflowOutput struct {
	TransactionID     string
	TransactionStatus string
	FromAccountID     int64
	ToAccountID       int64
	Errors            []string
}

type TransactionRejectWorkflowInput struct {
	TransactionID string
	Reason        string
}

type TransactionApproveWorkflowInput struct {
	*TransactionWorkflowInput
	TransactionID string
}

type SaveTransactionActivityOutput struct {
	AbortTx       bool
	AbortTxReason error
	TxUUID        string
	TxStatus      string
}

type ApplyTransactionActivityInput struct {
	TypeOfTransaction string
}
