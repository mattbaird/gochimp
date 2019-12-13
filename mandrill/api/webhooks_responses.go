package api

type WebhooksInfoResponse struct {
	ID          int      `json:"id"`
	URL         string   `json:"url"`
	Description string   `json:"description"`
	AuthKey     string   `json:"auth_key"`
	Events      []string `json:"events"`
	CreatedAt   Time     `json:"created_at"`
	LastSentAt  Time     `json:"last_sent_at"`
	BatchesSent int32    `json:"batches_sent"`
	EventsSent  int32    `json:"events_sent"`
	LastError   string   `json:"last_error"`
}

type WebhooksListResponse []WebhooksInfoResponse

type WebhooksAddResponse WebhooksInfoResponse

type WebhooksUpdateResponse WebhooksInfoResponse

type WebhooksDeleteResponse WebhooksInfoResponse
