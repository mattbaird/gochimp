package api

type IPsListRequest struct {
	Key string `json:"key"`
}

type IPsInfoRequest struct {
	Key string `json:"key"`
	IP  string `json:"ip"`
}

type IPsProvisionRequest struct {
	Key    string `json:"key"`
	WarmUp bool   `json:"warmup,omitempty"`
	Pool   string `json:"pool,omitempty"`
}

type IPsStartWarmUpRequest struct {
	Key string `json:"key"`
	IP  string `json:"ip"`
}

type IPsCancelWarmUpRequest struct {
	Key string `json:"key"`
	IP  string `json:"ip"`
}

type IPsSetPoolRequest struct {
	Key        string `json:"key"`
	IP         string `json:"ip"`
	Pool       string `json:"pool"`
	CreatePool bool   `json:"create_pool,omitempty"`
}

type IPsDeleteRequest struct {
	Key string `json:"key"`
	IP  string `json:"ip"`
}

type IPsListPoolsRequest struct {
	Key string `json:"key"`
}

type IPsPoolInfoRequest struct {
	Key  string `json:"key"`
	Pool string `json:"pool"`
}

type IPsCreatePoolRequest struct {
	Key  string `json:"key"`
	Pool string `json:"pool"`
}

type IPsDeletePoolRequest struct {
	Key  string `json:"key"`
	Pool string `json:"pool"`
}

type IPsCheckCustomDNSRequest struct {
	Key    string `json:"key"`
	IP     string `json:"ip"`
	Domain string `json:"domain"`
}

type IPsSetCustomDNSRequest struct {
	Key    string `json:"key"`
	IP     string `json:"ip"`
	Domain string `json:"domain"`
}
