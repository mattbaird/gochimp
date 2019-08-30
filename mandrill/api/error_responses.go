package api

// ErrorResponse represents the format of errors from calling the api
type ErrorResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}
