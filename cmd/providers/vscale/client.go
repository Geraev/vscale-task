package vscale

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	VSCALE_API_URL = "https://api.vscale.io/v1/"
	VSCALE_API_SCALETS = "scalets"
)

type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

func NewClient(token string) *Client {
	var client = &Client{
		token:      token,
		httpClient: http.DefaultClient,
		baseURL:    VSCALE_API_URL,
	}

	return client
}

func (c *Client) CreateServer(servConf *ServerConfiguration) (ctid int, err error) {
	var req *http.Request

	req, err = c.newRequest(http.MethodPost, VSCALE_API_SCALETS, servConf)
	if err != nil {
		return 0, fmt.Errorf("failed to create request with error: %v", err)
	}

	return 0, nil
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	var buf = new(bytes.Buffer)

	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, fmt.Errorf("failed to encode request body with error: %v", err)
		}
	}

	req, err := http.NewRequest(method, VSCALE_API_URL + path, buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create new http request with error: %v", err)
	}

	req.Header.Add("X-Token", c.token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req, nil
}