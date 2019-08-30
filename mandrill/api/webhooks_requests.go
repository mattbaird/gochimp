package api

type WebhooksInfoRequest struct {
	Key string `json:"key"`
	ID  int    `json:"id"`
}

type WebhooksListRequest struct {
	Key string `json:"key"`
}

type WebhooksAddRequest struct {
	Key         string   `json:"key"`
	URL         string   `json:"url"`
	Description string   `json:"description,omitempty"`
	Events      []string `json:"events,omitempty"`
}

type WebhooksUpdateRequest struct {
	Key         string   `json:"key"`
	ID          int      `json:"id"`
	URL         string   `json:"url"`
	Description string   `json:"description,omitempty"`
	Events      []string `json:"event"`
}

type WebhooksDeleteRequest struct {
	Key string `json:"key"`
	ID  int    `json:"id"`
}
