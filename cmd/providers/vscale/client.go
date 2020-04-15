package vscale

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"vscale-task/cmd/providers"
)

const (
	VSCALE_API_URL     = "https://api.vscale.io/v1/"
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

func (c *Client) CreateServer(servReq *providers.CreateServerRequest) (servResp providers.CreateServerResponse, err error) {
	var req *http.Request
	if req, err = c.newRequest(http.MethodPost, VSCALE_API_SCALETS, servReq); err != nil {
		return servResp, fmt.Errorf("failed to create request with error: %v", err)
	}

	var (
		resp *http.Response
		body []byte
	)

	if resp, err = c.httpClient.Do(req); err != nil {
		return servResp, fmt.Errorf("failed to response with error: %v", err)
	}
	body, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusNoContent {
		switch resp.StatusCode {
		case http.StatusTooManyRequests:
			return servResp, providers.ErrTooManyRequests
		case http.StatusGatewayTimeout:
			return servResp, providers.ErrGatewayTimeout
		default:
			var errorResponse providers.ErrorResponse
			if err = json.Unmarshal(body, &errorResponse); err != nil {
				return servResp, fmt.Errorf("unmarshaling error: %v: %s", err, string(body))
			}
			return servResp, fmt.Errorf("error with status code: %d", resp.StatusCode)
		}
	}

	if err = json.Unmarshal(body, &servResp); err != nil {
		return servResp, fmt.Errorf("unmarshaling error: %v: %s", err, string(body))
	}

	return servResp, nil
}

func (c *Client) DeleteServer(ctid int64) (servResp providers.DeleteServerResponse, err error) {
	var req *http.Request
	if req, err = c.newRequest(http.MethodDelete, fmt.Sprint(VSCALE_API_SCALETS, ctid), nil); err != nil {
		return servResp, fmt.Errorf("failed to create request with error: %v", err)
	}

	var (
		resp *http.Response
		body []byte
	)

	if resp, err = c.httpClient.Do(req); err != nil {
		return servResp, fmt.Errorf("failed to response with error: %v", err)
	}
	body, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusNoContent {
		switch resp.StatusCode {
		case http.StatusTooManyRequests:
			return servResp, providers.ErrTooManyRequests
		case http.StatusGatewayTimeout:
			return servResp, providers.ErrGatewayTimeout
		default:
			var errorResponse providers.ErrorResponse
			if err = json.Unmarshal(body, &errorResponse); err != nil {
				return servResp, fmt.Errorf("unmarshaling error: %v: %s", err, string(body))
			}
			return servResp, fmt.Errorf("error with status code: %d", resp.StatusCode)
		}
	}

	if err = json.Unmarshal(body, &servResp); err != nil {
		return servResp, fmt.Errorf("unmarshaling error: %v: %s", err, string(body))
	}

	return servResp, nil
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	var buf = new(bytes.Buffer)

	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, fmt.Errorf("failed to encode request body with error: %v", err)
		}
	}

	req, err := http.NewRequest(method, VSCALE_API_URL+path, buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create new http request with error: %v", err)
	}

	req.Header.Add("X-Token", c.token)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Accept", "application/json")

	return req, nil
}
