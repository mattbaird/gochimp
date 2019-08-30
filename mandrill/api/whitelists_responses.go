package api

type WhiteListsAddResponse struct {
	Email string `json:"email"`
	Added bool   `json:"added"`
}

type WhiteListsListResponse []struct {
	Email     string `json:"email"`
	Detail    string `json:"detail"`
	CreatedAt Time   `json:"created_at"`
}

type WhiteListsDeleteResponse struct {
	Email   string `json:"email"`
	Deleted bool   `json:"deleted"`
}
