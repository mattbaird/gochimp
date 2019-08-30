package api

type SubAccountsListRequest struct {
	Key string `json:"key"`
	Q   string `json:"q,omitempty"`
}
type SubAccountsInfoRequest struct {
	Key string `json:"key"`
	ID  string `json:"id"`
}

type SubAccountsAddRequest struct {
	Key         string `json:"key"`
	ID          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Notes       string `json:"notes,omitempty"`
	CustomQuota int    `json:"custom_quota,omitempty"`
}

type SubAccountsUpdateRequest struct {
	Key         string `json:"key"`
	ID          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Notes       string `json:"notes,omitempty"`
	CustomQuota int    `json:"custom_quota,omitempty"`
}

type SubAccountsDeleteRequest struct {
	Key string `json:"key"`
	ID  string `json:"id"`
}

type SubAccountsPauseRequest struct {
	Key string `json:"key"`
	ID  string `json:"id"`
}

type SubAccountsResumeRequest struct {
	Key string `json:"key"`
	ID  string `json:"id"`
}
