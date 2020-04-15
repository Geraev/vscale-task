package manager

import (
	"errors"
)

var (
	ErrFatalApiError   = errors.New("fatal api error occurred")
	ErrGroupIDNotFound = errors.New("groupID not found")
	ErrDeleteServer    = errors.New("delete servre error for API")
)
