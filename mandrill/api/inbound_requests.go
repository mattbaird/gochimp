package api

// InboundDomainRequest represents requests to inbound/*-domain.json
type InboundDomainRequest struct {
	Key    string `json:"key"`
	Domain string `json:"domain"`
}

// InboundAddRouteRequest represents a request to inbound/add-route.json
type InboundAddRouteRequest struct {
	Key     string `json:"key"`
	Domain  string `json:"domain"`
	Pattern string `json:"pattern"`
	URL     string `json:"url"`
}

// InboundUpdateRouteRequest represents a request to inbound/update-route.json
type InboundUpdateRouteRequest struct {
	Key     string `json:"key"`
	ID      string `json:"id"`
	Pattern string `json:"pattern"`
	URL     string `json:"url"`
}

// InboundDeleteRouteRequest represents a request to inbound/delete-route.json
type InboundDeleteRouteRequest struct {
	Key string `json:"key"`
	ID  string `json:"id"`
}

type InboundSendRawRequest struct {
	Key           string   `json:"key"`
	RawMessage    string   `json:"raw_message"`
	To            []string `json:"to,omitempty"`
	MailFrom      string   `json:"mail_from,omitempty"`
	Helo          string   `json:"helo,omitempty"`
	ClientAddress string   `json:"client_address,omitempty"`
}
