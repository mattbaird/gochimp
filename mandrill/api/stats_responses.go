package api

// StatResponse represents a Stat collection entry from the Mandrill API
type StatResponse struct {
	Sent         int32 `json:"sent"`
	HardBounces  int32 `json:"hard_bounces"`
	SoftBounces  int32 `json:"soft_bounces"`
	Rejects      int32 `json:"rejects"`
	Complaints   int32 `json:"complaints"`
	Unsubs       int32 `json:"unsubs"`
	Opens        int32 `json:"opens"`
	UniqueOpens  int32 `json:"unique_opens"`
	Clicks       int32 `json:"clicks"`
	UniqueClicks int32 `json:"unique_clicks"`
	Reputation   int32 `json:"reputation"`
}

// StatsResponse represents the stats collection in the Mandrill API
type StatsResponse map[string]StatResponse
