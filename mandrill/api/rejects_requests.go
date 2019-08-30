package api

type RejectsAddRequest struct {
	Key        string `json:"key"`
	Email      string `json:"email"`
	Comment    string `json:"comment,omitempty"`
	SubAccount string `json:"subaccount,omitempty"`
}

type RejectsListRequest struct {
	Key            string `json:"key"`
	Email          string `json:"email,omitempty"`
	IncludeExpired bool   `json:"include_expired,omitempty"`
	SubAccount     string `json:"subaccount,omitempty"`
}

type RejectsDeleteRequest struct {
	Key        string `json:"key"`
	Email      string `json:"email"`
	SubAccount string `json:"subaccount,omitempty"`
}
