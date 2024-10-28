package constants

import "errors"

var (
	ErrFakeOrInvalidAccount = errors.New("fake or invalid account")
	ErrFakeOrInvalidBank    = errors.New("fake or invalid bank")
	ErrInvalidAmount        = errors.New("invalid amount")
	ErrSavingTransaction    = errors.New("error saving transaction")
)
