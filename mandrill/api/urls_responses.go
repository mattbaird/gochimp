package api

// URLResponse represents an individual item in a call to urls/list.json
type URLResponse struct {
	URL          string `json:"url"`
	Sent         int32  `json:"sent"`
	Clicks       int32  `json:"clicks"`
	UniqueClicks int32  `json:"unique_clicks"`
}

// URLTimeSeriesResponse represents an individual item in a call to urls/time-series.json
type URLTimeSeriesResponse struct {
	Time         Time  `json:"time"`
	Sent         int32 `json:"sent"`
	Clicks       int32 `json:"clicks"`
	UniqueClicks int32 `json:"unique_clicks"`
}

// URLTrackingDomainResponse represents an individual item in a call to urls/tracking-domains.json
type URLTrackingDomainResponse struct {
	Domain       string `json:"domain"`
	CreatedAt    Time   `json:"created_at"`
	LastTestedAt Time   `json:"last_tested_at"`
	CName        struct {
		Valid      bool   `json:"valid"`
		ValidAfter Time   `json:"valid_after"`
		Error      string `json:"error"`
	} `json:"cname"`
	ValidTracking bool `json:"valid_tracking"`
}

// URLsSearchResponse represents a call to urls/search.json
type URLsSearchResponse []URLResponse

// URLsListResponse represents a call to urls/list.json
type URLsListResponse []URLResponse

// URLsTimeSeriesResponse represents a call to urls/time-series.json
type URLsTimeSeriesResponse []URLTimeSeriesResponse

// URLsTrackingDomainsResponse represents a call to urls/tracking-domains.json
type URLsTrackingDomainsResponse []URLTrackingDomainResponse

// URLsAddTrackingDomainResponse represents a call to urls/add-tracking-domain.json
type URLsAddTrackingDomainResponse URLTrackingDomainResponse

// URLsCheckTrackingDomainResponse represents a call to urls/check-tracking-domain.json
type URLsCheckTrackingDomainResponse URLTrackingDomainResponse
