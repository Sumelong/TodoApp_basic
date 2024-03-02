package handlers

import "errors"

var (
	ErrRequestDecodeFailed = errors.New("failed to decode request body")
)
