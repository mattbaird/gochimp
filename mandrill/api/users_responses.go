package api

type UsersInfoResponse struct {
	Username    string                  `json:"username"`
	CreatedAt   Time                    `json:"created_at"`
	PublicID    string                  `json:"public_id"`
	Reputation  int32                   `json:"reputation"`
	HourlyQuota int32                   `json:"hourly_quota"`
	Backlog     int32                   `json:"backlog"`
	Stats       map[string]StatResponse `json:"stats"`
}

type UsersPing2Response struct {
	Ping string `json:"PING"`
}

type UsersSenderResponse struct {
	Address   string `json:"address"`
	CreatedAt Time   `json:"created_at"`
	StatResponse
}
type UsersSendersResponse []UsersSenderResponse
