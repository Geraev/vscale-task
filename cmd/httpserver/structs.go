package httpserver

import "vscale-task/cmd/providers"

type (
	APICreateRequest struct {
		providers.CreateServerRequest
	}

	APICreateRespone struct {
		Status  string `json:"status"`
		GroupID int64  `json:"group_id"`
	}

	APIGetStatus struct {
		APICreateRespone
	}
)
