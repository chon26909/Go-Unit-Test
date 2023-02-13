package services

import "errors"

var (
	ErrZeroAmount = errors.New("purchase amout coould not be zero")
	ErrRepository = errors.New("repository error")
)
