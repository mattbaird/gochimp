package api

type WhiteListsAddRequest struct {
	Key     string `json:"key"`
	Email   string `json:"email"`
	Comment string `json:"comment,omitempty"`
}

type WhiteListsListRequest struct {
	Key   string `json:"key"`
	Email string `json:"email,omitempty"`
}

type WhiteListsDeleteRequest struct {
	Key   string `json:"key"`
	Email string `json:"email"`
}
