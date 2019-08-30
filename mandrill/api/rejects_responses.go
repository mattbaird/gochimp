package api

type RejectsAddResponse struct {
	Email string `json:"email"`
	Added bool   `json:"added"`
}

type RejectsListResponse []struct {
	Email       string         `json:"email"`
	Reason      string         `json:"reason"`
	Detail      string         `json:"detail"`
	CreatedAt   Time           `json:"created_at"`
	LastEventAt Time           `json:"last_event_at"`
	ExpiresAt   Time           `json:"expires_at"`
	Expired     bool           `json:"expired"`
	SubAccount  string         `json:"subaccount"`
	Sender      SenderResponse `json:"sender"`
}

type RejectsDeleteResponse struct {
	Email      string `json:"email"`
	Deleted    bool   `json:"deleted"`
	SubAccount string `json:"subaccount"`
}
