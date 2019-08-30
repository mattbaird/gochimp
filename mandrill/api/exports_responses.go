package api

type ExportsInfoResponse struct {
	ID         string `json:"id"`
	CreatedAt  Time   `json:"created_at"`
	Type       string `json:"type"`
	FinishedAt Time   `json:"finished_at"`
	State      string `json:"state"`
	ResultURL  string `json:"result_url"`
}

type ExportsListResponse []ExportsInfoResponse

type ExportsRejectsResponse struct {
	ID         string `json:"id"`
	CreatedAt  Time   `json:"created_at"`
	Type       string `json:"type"`
	FinishedAt Time   `json:"finished_at"`
	State      string `json:"state"`
	ResultURL  string `json:"result_url"`
}

type ExportsWhiteListResponse struct {
	ID         string `json:"id"`
	CreatedAt  Time   `json:"created_at"`
	Type       string `json:"type"`
	FinishedAt Time   `json:"finished_at"`
	State      string `json:"state"`
	ResultURL  string `json:"result_url"`
}

type ExportsActivityResponse struct {
	ID         string `json:"id"`
	CreatedAt  Time   `json:"created_at"`
	Type       string `json:"type"`
	FinishedAt Time   `json:"finished_at"`
	State      string `json:"state"`
	ResultURL  string `json:"result_url"`
}
