package vscale

type ServerConfiguration struct {
	MakeFrom string  `json:"make_from,omitempty"`
	Rplan    string  `json:"rplan,omitempty"`
	DoStart  bool    `json:"do_start,omitempty"`
	Name     string  `json:"name,omitempty"`
	Keys     []int64 `json:"keys,omitempty"`
	Location string  `json:"location,omitempty"`
}

func DefaultServerConfiguration() *ServerConfiguration {
	return &ServerConfiguration{
		MakeFrom: "ubuntu_18.04_64_001_master",
		Rplan:    "small",
		DoStart:  false,
		Name:     "DefaultServerName",
		Keys:     nil,
		Location: "spb0",
	}
}
