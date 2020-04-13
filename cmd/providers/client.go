package providers

type Client interface {
	CreateServer(*CreateServerRequest) (CreateServerResponse, error)
	DeleteServer(int64) (DeleteServerResponse, error)
}
