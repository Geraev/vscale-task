package vscale

import "errors"

var (
	ErrTooManyRequests = errors.New("too many requests")
	ErrGatewayTimeout  = errors.New("gateway timeout")
)
