package models

import (
	models "github.com/hjoshi123/temporal-loan-app/pkg/models/postgres"
)

type FakeDetectionOutput struct {
	IsFake    bool
	FakeError error
}

type TransactionWorkflowInput struct {
	Amount            float64
	FromAccountNumber int64
	ToAccountNumber   int64
	FromBankID        int
	ToBankID          int
	TxType            string
	Memo              string
	TxUUID            string
	TxStatus          string
}

type TransactionWorkflowOutput struct {
	TransactionID     string
	TransactionStatus string
	FromAccountID     int64
	ToAccountID       int64
	Errors            []string
}

type TransactionRejectWorkflowInput struct {
	*TransactionWorkflowInput
	TransactionID string
	Reason        string
}

type TransactionApproveWorkflowInput struct {
	*TransactionWorkflowInput
	TransactionID string
}

type TransactionApproveWorkflowOutput struct {
	TransactionID string
	TxStatus      string
}

type SaveTransactionActivityOutput struct {
	AbortTx       bool
	AbortTxReason error
	TxUUID        string
	TxStatus      string
	Tx            *models.Transaction
}

type ApplyTransactionActivityInput struct {
	Tx                *models.Transaction
	TypeOfTransaction string
}

type ApplyTransactionActivityOutput struct {
	TypeOfTransaction string
	TransactionID     string
	ApplyStatus       string
}
