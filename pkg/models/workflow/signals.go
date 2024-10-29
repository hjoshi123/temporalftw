package models

type SignalReject struct {
	TxRejectedReason string
	TxUUID           string
}

type SignalApprove struct {
	TxApproved bool
	TxUUID     string
}
