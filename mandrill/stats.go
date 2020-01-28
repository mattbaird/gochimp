package mandrill

import (
	"github.com/lusis/gochimp/mandrill/api"
	"time"
)

// Stats are data about various objects in the api
type Stats struct {
	Sent         int32
	HardBounces  int32
	SoftBounces  int32
	Rejects      int32
	Complaints   int32
	Unsubs       int32
	Opens        int32
	UniqueOpens  int32
	Clicks       int32
	UniqueClicks int32
	Reputation   int32
}

// TimeSeries is more detailed Stat data with additional context related to the time
type TimeSeries struct {
	Time         time.Time
	Sent         int32
	HardBounces  int32
	SoftBounces  int32
	Rejects      int32
	Complaints   int32
	Unsubs       int32
	Opens        int32
	UniqueOpens  int32
	Clicks       int32
	UniqueClicks int32
	Reputation   int32
}

func statsResponseToStats(resp api.StatResponse) Stats {
	return Stats{
		Sent:         resp.Sent,
		HardBounces:  resp.HardBounces,
		SoftBounces:  resp.SoftBounces,
		Rejects:      resp.Rejects,
		Complaints:   resp.Complaints,
		Unsubs:       resp.Unsubs,
		Opens:        resp.Opens,
		UniqueOpens:  resp.UniqueOpens,
		Clicks:       resp.Clicks,
		UniqueClicks: resp.UniqueClicks,
		Reputation:   resp.Reputation,
	}
}
