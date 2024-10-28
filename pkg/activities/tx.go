package activities

import (
	"context"
	"fmt"
	"github.com/ericlagergren/decimal"
	"github.com/google/uuid"
	"github.com/hjoshi123/temporal-loan-app/pkg/constants"
	datastoreIface "github.com/hjoshi123/temporal-loan-app/pkg/datastore/interface"
	datastore "github.com/hjoshi123/temporal-loan-app/pkg/datastore/postgres"
	models "github.com/hjoshi123/temporal-loan-app/pkg/models/postgres"
	wfModels "github.com/hjoshi123/temporal-loan-app/pkg/models/workflow"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
	"math"
)

type TransactionActivities struct {
	txStore datastoreIface.TransactionStore
}

func NewTransactionActivities() *TransactionActivities {
	txActivities := new(TransactionActivities)
	txActivities.txStore = datastore.NewTransactionStore()
	return txActivities
}

func (txAc *TransactionActivities) SaveTransactionActivity(ctx context.Context, txInput *wfModels.TransactionWorkflowInput) (*wfModels.SaveTransactionActivityOutput, error) {
	output := new(wfModels.SaveTransactionActivityOutput)

	if txInput.Amount < 0 {
		output.AbortTx = true
		output.AbortTxReason = constants.ErrInvalidAmount
		return output, nil
	}

	decimalScaleAmount := decimal.New(int64(txInput.Amount*math.Pow10(6)), 6)
	decimalAmount := types.NewDecimal(decimalScaleAmount)

	txUUID := null.NewString(uuid.New().String(), true)
	err := txAc.txStore.SaveTransaction(ctx, &models.Transaction{
		Amount:            decimalAmount,
		FromAccountNumber: txInput.FromAccountNumber,
		FromBankID:        txInput.FromBankID,
		ToAccountNumber:   txInput.ToAccountNumber,
		ToBankID:          txInput.ToBankID,
		TransactionID:     txUUID,
	}, txInput.TxType)
	if err != nil {
		output.AbortTx = true
		output.AbortTxReason = fmt.Errorf("failed to save transaction: %w", err)
		return output, nil
	}

	savedTx, err := txAc.txStore.GetTransactionByUUID(ctx, txUUID.String)
	if err != nil {
		output.AbortTx = true
		output.AbortTxReason = fmt.Errorf("failed to get saved transaction: %w", err)
		return output, nil
	}

	output.TxUUID = savedTx.TransactionID.String
	output.TxStatus = savedTx.R.TransactionStatus.Name

	return output, nil
}

func (txAc *TransactionActivities) ApplyTransactionActivity(ctx context.Context, applyTx *wfModels.ApplyTransactionActivityInput) {

}
