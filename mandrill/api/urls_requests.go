package api

type URLsListRequest struct {
	Key string `json:"key"`
}

type URLsSearchRequest struct {
	Key string `json:"key"`
	Q   string `json:"q"`
}

type URLsTimeSeriesRequest struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

type URLsTrackingDomainsRequest struct {
	Key string `json:"key"`
}

type URLsAddTrackingDomainRequest struct {
	Key    string `json:"key"`
	Domain string `json:"domain"`
}

type URLsCheckTrackingDomainRequest struct {
	Key    string `json:"key"`
	Domain string `json:"domain"`
}
