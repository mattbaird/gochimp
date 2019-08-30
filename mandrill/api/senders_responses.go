package api

import (
	"time"
)

// SenderResponse represents the elements in a senders/list.json api call
type SenderResponse struct {
	Address      string `json:"address"`
	CreatedAt    Time   `json:"created_at"`
	Sent         int32  `json:"sent"`
	HardBounces  int32  `json:"hard_bounces"`
	SoftBounces  int32  `json:"soft_bounces"`
	Rejects      int32  `json:"rejects"`
	Complaints   int32  `json:"complaints"`
	Unsubs       int32  `json:"unsubs"`
	Opens        int32  `json:"opens"`
	Clicks       int32  `json:"clicks"`
	UniqueOpens  int32  `json:"unique_opens"`
	UniqueClicks int32  `json:"unique_clicks"`
}

// SendersInfoResponse represents the response from a senders/info.json call to the api
type SendersInfoResponse struct {
	Address     string        `json:"address"`
	CreatedAt   time.Time     `json:"created_at"`
	Sent        int32         `json:"sent"`
	HardBounces int32         `json:"hard_bounces"`
	SoftBounces int32         `json:"soft_bounces"`
	Rejects     int32         `json:"rejects"`
	Complaints  int32         `json:"complaints"`
	Unsubs      int32         `json:"unsubs"`
	Opens       int32         `json:"opens"`
	Clicks      int32         `json:"clicks"`
	Stats       StatsResponse `json:"stats"`
}

// SendersListResponse represents the response from a senders/list.json api call
type SendersListResponse []SenderResponse

type SendersDomainResponse struct {
	Domain       string `json:"domain"`
	CreatedAt    Time   `json:"created_at"`
	LastTestedAt Time   `json:"last_tested_at"`
	SPF          struct {
		Valid      bool   `json:"valid"`
		ValidAfter Time   `json:"valid_after"`
		Error      string `json:"error"`
	} `json:"spf"`
	DKIM struct {
		Valid      bool   `json:"valid"`
		ValidAfter Time   `json:"valid_after"`
		Error      string `json:"error"`
	} `json:"dkim"`
	VerifiedAt   Time `json:"verified_at"`
	ValidSigning bool `json:"valid_signing"`
}

type SendersDomainsResponse []SendersDomainResponse

type SendersAddDomainResponse SendersDomainResponse

type SendersCheckDomainResponse SendersDomainResponse

type SendersVerifyDomainResponse struct {
	Status string `json:"status"`
	Domain string `json:"domain"`
	Email  string `json:"email"`
}

type SendersTimeSeriesResponse []struct {
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
