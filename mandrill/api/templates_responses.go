package api

type TemplatesInfoResponse struct {
	Slug             string   `json:"slug"`
	Name             string   `json:"name"`
	Labels           []string `json:"labels"`
	Code             string   `json:"code"`
	FromEmail        string   `json:"from_email"`
	FromName         string   `json:"from_name"`
	Text             string   `json:"text"`
	PublishName      string   `json:"publish_name"`
	PublishCode      string   `json:"publish_code"`
	PublishSubject   string   `json:"publish_subject"`
	PublishFromEmail string   `json:"publish_from_email"`
	PublishFromName  string   `json:"publish_from_name"`
	PublishText      string   `json:"publish_text"`
	PublishedAt      Time     `json:"published_at"`
	CreatedAt        Time     `json:"created_at"`
	UpdatedAt        Time     `json:"updated_at"`
}

type TemplatesListResponse []TemplatesInfoResponse

type TemplatesAddResponse TemplatesInfoResponse

type TemplatesUpdateResponse TemplatesInfoResponse

type TemplatesPublishResponse TemplatesInfoResponse

type TemplatesDeleteResponse TemplatesInfoResponse

type TemplatesTimeSeriesResponse []struct {
	Time         Time  `json:"time"`
	Sent         int32 `json:"sent"`
	HardBounces  int32 `json:"hard_bounces"`
	SoftBounces  int32 `json:"soft_bounces"`
	Rejects      int32 `json:"rejects"`
	Complaints   int32 `json:"complaints"`
	Unsubs       int32 `json:"unsubs"`
	Opens        int32 `json:"opens"`
	Clicks       int32 `json:"clicks"`
	UniqueOpens  int32 `json:"unique_opens"`
	UniqueClicks int32 `json:"unique_clicks"`
}

type TemplatesRenderResponse struct {
	HTML string `json:"html"`
}
