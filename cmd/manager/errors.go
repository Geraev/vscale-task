package manager

import (
	"errors"
)

var (
	ErrNeedRollback = errors.New("fatal api error occurred")
	ErrGroupIDNotFound  = errors.New("groupID not found")
	ErrDeleteServer  = errors.New("delete servre error for API")
)
