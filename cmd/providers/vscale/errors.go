package vscale

import (
	"errors"
	"net/http"
)

var (
	ErrTooManyRequests = errors.New(http.StatusText(http.StatusTooManyRequests))
	ErrGatewayTimeout  = errors.New(http.StatusText(http.StatusGatewayTimeout))
)
