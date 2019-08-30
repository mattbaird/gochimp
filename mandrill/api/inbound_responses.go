package api

// InboundDomain represents an inbound domain entry in calls to various inbound domain api
type InboundDomainResponse struct {
	Domain    string `json:"domain"`
	ValidMx   bool   `json:"valid_mx"`
	CreatedAt Time   `json:"created_at"`
}

// InboundRouteResponse represents a mailbox route entry in calls to various inbound route api
type InboundRouteResponse struct {
	ID      string `json:"id"`
	Pattern string `json:"pattern"`
	URL     string `json:"url"`
}

// InboundRawResponse represents an element in an inbound/send-raw.json api call
type InboundRawResponse struct {
	Email   string `json:"email"`
	Pattern string `json:"pattern"`
	URL     string `json:"url"`
}

// InboundListDomainsResponse represents an inbound/domains.json call to the api
type InboundListDomainsResponse []InboundDomainResponse

// InboundAddDomainResponse represents an inbound/add-domain.json call to the api
type InboundAddDomainResponse InboundDomainResponse

// InboundCheckDomainResponse represents an inbound/check-domain.json call to the api
type InboundCheckDomainResponse InboundDomainResponse

// InboundDeleteDomainResponse represents an inbound/delete-domain.json call to the api
type InboundDeleteDomainResponse InboundDomainResponse

// InboundListRoutesResponse represents a collection of Route
type InboundListRoutesResponse []InboundRouteResponse

// InboundAddRouteResponse represents an inbound/add-route.json call to the api
type InboundAddRouteResponse InboundRouteResponse

// InboundUpdateRouteResponse represents an inbound/update-route.json call to the api
type InboundUpdateRouteResponse InboundRouteResponse

// InboundDeleteRouteResponse represents an inbound/delete-route.json call to the api
type InboundDeleteRouteResponse InboundRouteResponse

// InboundSendRawResponse represents an inbound/send-raw.json call to the api
type InboundSendRawResponse []InboundRawResponse
