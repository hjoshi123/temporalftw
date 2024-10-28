package constants

type TransactionSignals string

func (t TransactionSignals) String() string {
	return string(t)
}

const (
	TransactionSignalApprove TransactionSignals = "APPROVE_TX"
	TransactionSignalReject  TransactionSignals = "REJECT_TX"
)
