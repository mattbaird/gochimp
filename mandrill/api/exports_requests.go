package api

type ExportsInfoRequest struct {
	Key string `json:"key"`
	ID  string `json:"id"`
}

type ExportsListRequest struct {
	Key string `json:"key"`
}

type ExportsRejectsRequest struct {
	Key         string `json:"key"`
	NotifyEmail string `json:"notify_email,omitempty"`
}

type ExportsWhiteListRequest struct {
	Key         string `json:"key"`
	NotifyEmail string `json:"notify_email,omitempty"`
}

type ExportsActivityRequest struct {
	Key         string   `json:"key"`
	NotifyEmail string   `json:"notify_email,omitempty"`
	DateFrom    string   `json:"date_from,omitempty"`
	DateTo      string   `json:"date_to,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Senders     []string `json:"senders,omitempty"`
	States      []string `json:"states,omitempty"`
	APIKeys     []string `json:"api_keys,omitempty"`
}
