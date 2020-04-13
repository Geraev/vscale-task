package providers

type (
	CreateServerRequest struct {
		MakeFrom string  `json:"make_from,omitempty"`
		Rplan    string  `json:"rplan,omitempty"`
		DoStart  bool    `json:"do_start,omitempty"`
		Name     string  `json:"name,omitempty"`
		Keys     []int64 `json:"keys,omitempty"`
		Location string  `json:"location,omitempty"`
	}

	CreateServerResponse struct {
		Status           string  `json:"status,omitempty"`
		PublicAddresses  Address `json:"public_address,omitempty"`
		Active           bool    `json:"active,omitempty"`
		Location         string  `json:"location,omitempty"`
		Locked           bool    `json:"locked,omitempty"`
		Hostname         string  `json:"hostname,omitempty"`
		Created          string  `json:"created,omitempty"`
		Keys             []Keys  `json:"keys,omitempty"`
		PrivateAddresses Address `json:"private_address,omitempty"`
		MadeFrom         string  `json:"made_from,omitempty"`
		Name             string  `json:"name,omitempty"`
		Rplan            string  `json:"rplan,omitempty"`
		CTID             int64   `json:"ctid,omitempty"`
	}

	DeleteServerResponse struct {
		CreateServerResponse
		Deleted           string `json:"deleted,omitempty"`
		BlockReason       string `json:"block_reason"`
		BlockReasonCustom string `json:"block_reason_custom"`
		DateBlock         string `json:"date_block"`
	}

	Address struct {
		Netmask string `json:"netmask,omitempty"`
		Gateway string `json:"gateway,omitempty"`
		Address string `json:"address,omitempty"`
	}

	Keys struct {
		Name string `json:"name,omitempty"`
		ID   int64  `json:"id,omitempty"`
	}

	ErrorResponse struct {
		Error struct {
			Code    string `json:"code"`
			Status  int    `json:"status"`
			Message string `json:"message"`
		} `json:"error"`
	}
)

func DefaultCreateServerRequest() *CreateServerRequest {
	return &CreateServerRequest{
		MakeFrom: "ubuntu_18.04_64_001_master",
		Rplan:    "small",
		DoStart:  false,
		Name:     "DefaultServerName",
		Keys:     nil,
		Location: "spb0",
	}
}
