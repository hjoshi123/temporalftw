package models

type SignalReject struct {
	TxRejectedReason string
}

type SignalApprove struct {
	TxApproved bool
}
