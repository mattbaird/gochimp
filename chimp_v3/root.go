package gochimp

const (
	root_path = "/"
)

type AccountContact struct {
	Company string `json:"company"`
	Addr1   string `json:"addr1"`
	Addr2   string `json:"addr2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

type IndustryStats struct {
	OpenRate   float64 `json:"open_rate"`
	BounceRate float64 `json:"bounce_rate"`
	ClickRate  float64 `json:"click_rate"`
}

// RootResponse - https://developer.mailchimp.com/documentation/mailchimp/reference/root/#read-get_root
type RootResponse struct {
	AccountID        string         `json:"account_id"`
	AccountName      string         `json:"account_name"`
	Email            string         `json:"email"`
	Role             string         `json:"role"`
	Contact          AccountContact `json:"contact"`
	ProEnabled       bool           `json:"pro_enabled"`
	LastLogin        string         `json:"last_login"`
	TotalSubscribers int            `json:"total_subscribers"`
	IndustyStats     IndustryStats  `json:"industry_stats"`
	Links            []Link         `json:"_links"`
}

// GetRoot queries the root of the API for stats
func (api ChimpAPI) GetRoot(params *BasicQueryParams) (*RootResponse, error) {
	response := new(RootResponse)
	err := api.Request("GET", root_path, params, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
