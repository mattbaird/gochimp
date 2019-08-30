package api

type IPInfoResponse struct {
	IP        string `json:"ip"`
	CreatedAt Time   `json:"created_at"`
	Pool      string `json:"pool"`
	Domain    string `json:"domain"`
	CustomDNS struct {
		Enabled bool   `json:"enabled"`
		Valid   bool   `json:"valid"`
		Error   string `json:"error"`
	} `json:"custom_dns"`
	WarmUp struct {
		WarmingUp bool `json:"warming_up"`
		StartAt   Time `json:"start_at"`
		EndAt     Time `json:"end_at"`
	}
}

type IPsProvisionResponse struct {
	RequestedAt Time `json:"requested_at"`
}

type IPsListResponse []IPInfoResponse

type IPsStartWarmUpResponse IPInfoResponse

type IPsCancelWarmUpResponse IPInfoResponse

type IPsSetPoolResponse IPInfoResponse

type IPsDeleteResponse struct {
	IP      string `json:"ip"`
	Deleted bool   `json:"deleted"`
}

type IPsPoolInfoResponse struct {
	Name      string           `json:"name"`
	CreatedAt Time             `json:"created_at"`
	IPs       []IPInfoResponse `json:"ips"`
}

type IPsListPoolsResponse []IPsPoolInfoResponse

type IPsCreatePoolResponse IPsPoolInfoResponse

type IPsDeletePoolResponse struct {
	Pool    string `json:"pool"`
	Deleted bool   `json:"deleted"`
}

type IPsCheckCustomDNSResponse struct {
	Valid bool   `json:"valid"`
	Error string `json:"error"`
}

type IPsSetCustomDNSResponse IPInfoResponse
