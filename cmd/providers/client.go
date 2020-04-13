package providers

import "vscale-task/cmd/providers/vscale"

type Client interface {
	CreateServer(*vscale.CreateServerRequest) (vscale.CreateServerResponse, error)
	DeleteServer(int64) (vscale.DeleteServerResponse, error)
}
